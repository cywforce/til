package metago

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"reflect"
	"strconv"

	"github.com/go-toolsmith/astcopy"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

const metagoBuildTag = "metago"
const metagoPackagePath = "github.com/vvakame/til/go/metago"

const (
	valueOfFuncNameString       = "ValueOf"      // metago.[ValueOf](obj)
	valueTypeString             = "Value"        // metago.[Value]
	fieldTypeString             = "Field"        // metago.[Field]
	valueFieldsMethodName       = "Fields"       // mv.[Fields]()
	fieldNameMethodName         = "Name"         // mf.[Name]()
	fieldValueMethodName        = "Value"        // mf.[Value]()
	fieldStructTagGetMethodName = "StructTagGet" // mf.[StructTagGet]("json")
)

type metaProcessor struct {
	cfg *Config

	currentPkg         *packages.Package
	currentFile        *ast.File
	currentTargetField *ast.Object
	currentBlockStmt   *ast.BlockStmt

	hasMetagoBuildTag     bool
	removeNodes           map[ast.Node]bool
	replaceNodes          map[ast.Node]ast.Node
	gotoCounter           int
	requiredContinueLabel []string
	requiredBreakLabel    []string

	// mv → obj への変換用
	valueMapping map[*ast.Object]ast.Expr
	// mf → obj.X への変換用 X 部分はBlockStmtに潜っていくまでわからんので obj 部分を持つ
	fieldMapping map[*ast.Object]ast.Expr

	nodeErrors NodeErrors
}

func (p *metaProcessor) Process(cfg *Config) (*Result, error) {
	p.cfg = cfg
	p.removeNodes = make(map[ast.Node]bool)
	p.replaceNodes = make(map[ast.Node]ast.Node)
	p.valueMapping = make(map[*ast.Object]ast.Expr)
	p.fieldMapping = make(map[*ast.Object]ast.Expr)
	p.requiredBreakLabel = nil
	p.requiredContinueLabel = nil

	pkgCfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedImports |
			packages.NeedDeps |
			packages.NeedTypes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
		BuildFlags: []string{"-tags", metagoBuildTag},
	}
	pkgs, err := packages.Load(pkgCfg, p.cfg.TargetPackages...)
	if err != nil {
		return nil, err
	}

	var errs []packages.Error
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		errs = append(errs, pkg.Errors...)
	})
	if len(errs) != 0 {
		return &Result{CompileErrors: errs}, errors.New("some errors occured")
	}

	result := &Result{}
	for _, pkg := range pkgs {
	file:
		for idx, file := range pkg.Syntax {
			p.hasMetagoBuildTag = false
			p.currentPkg = pkg
			p.currentFile = file
			if len(p.requiredContinueLabel) != 0 {
				panic("unknown state about requiredContinueLabel")
			}
			if len(p.requiredBreakLabel) != 0 {
				panic("unknown state about requiredBreakLabel")
			}

			fileResult := &FileResult{
				Package:  pkg,
				File:     file,
				FilePath: pkg.GoFiles[idx], // TODO これほんとうに安全？
			}
			result.Results = append(result.Results, fileResult)

			useMetagoPackage := astutil.UsesImport(file, metagoPackagePath)

			// . import してたら殺す
			for _, importSpec := range file.Imports {
				importPath, err := strconv.Unquote(importSpec.Path.Value)
				if err != nil {
					return nil, err
				}
				if metagoPackagePath == importPath && importSpec.Name != nil && importSpec.Name.Name == "." {
					p.Errorf(importSpec.Name, "don't use '.' import")
					fileResult.Errors = append(fileResult.Errors, p.nodeErrors...)
					p.nodeErrors = nil
					continue file
				}
			}

			for _, commentGroup := range file.Comments {
				// FileのCommentsは自動では歩かれないので自分でやる
				astutil.Apply(
					commentGroup,
					p.ApplyPre,
					p.ApplyPost,
				)
			}

			astutil.Apply(
				file,
				p.ApplyPre,
				p.ApplyPost,
			)

			if !p.hasMetagoBuildTag {
				if useMetagoPackage {
					p.Noticef(file, "this file has %s buildtag but doesn't use metago package. ignored", metagoBuildTag)
				}
				fileResult.Errors = append(fileResult.Errors, p.nodeErrors...)
				p.nodeErrors = nil

				continue
			}

			fileResult.Errors = append(fileResult.Errors, p.nodeErrors...)
			p.nodeErrors = nil

			var buf bytes.Buffer
			buf.Write([]byte("// Code generated by metago. DO NOT EDIT.\n\n"))
			buf.Write([]byte(fmt.Sprintf("//+build !%s\n\n", metagoBuildTag)))
			err := format.Node(&buf, pkg.Fset, file)
			if err != nil {
				return nil, err
			}

			fileResult.GeneratedCode = buf.String()
		}
	}

	var nErrs NodeErrors
	for _, fileResult := range result.Results {
		for _, nErr := range fileResult.Errors {
			if nErr.ErrorLevel <= ErrorLevelWarning {
				nErrs = append(nErrs, nErr)
			}
		}
	}
	if len(nErrs) != 0 {
		return result, nErrs
	}

	return result, nil
}

func (p *metaProcessor) ApplyPre(cursor *astutil.Cursor) bool {
	current := cursor.Node()
	if current == nil {
		return true
	}

	if p.removeNodes[current] {
		cursor.Delete()
		return false
	}
	if n := p.replaceNodes[current]; n != nil {
		cursor.Replace(n)
		return false
	}

	switch node := current.(type) {
	case *ast.Comment:
		if p.checkMetagoBuildTagComment(cursor, node) {
			return false
		}

	case *ast.AssignStmt:
		if p.checkMetagoValueOfAssignStmt(cursor, node) {
			return true
		}

	case *ast.Ident:
		if p.checkReplaceTargetIdent(cursor, node) {
			return false
		}

	case *ast.RangeStmt:
		if p.checkMetagoFieldRange(cursor, node) {
			return false
		}

	case *ast.CallExpr:
		if p.checkInlineTemplateCallExpr(cursor, node) {
			return false
		}
		if p.checkUseMetagoFieldValue(cursor, node) {
			return false
		}
		if p.checkUseMetagoFieldName(cursor, node) {
			return false
		}
		if p.checkUseMetagoStructTagGet(cursor, node) {
			return false
		}

	case *ast.IfStmt:
		if p.checkIfStmtInInitWithTypeAssert(cursor, node) {
			return false
		}
		if p.checkIfStmtInCondWithTypeAssert(cursor, node) {
			return false
		}

	case *ast.TypeSwitchStmt:
		if p.checkTypeSwitchStmt(cursor, node) {
			return false
		}

	case *ast.FuncDecl:
		if p.isInlineTemplateFuncDecl(cursor, node) {
			// 仮引数に metago.Value があったら、展開処理の対象ではないのでskip
			return false
		}

	case *ast.BranchStmt:
		if p.checkInRangeBranchStmt(cursor, node) {
			return false
		}

	case *ast.BlockStmt:
		p.currentBlockStmt = node
	}

	return true
}

func (p *metaProcessor) ApplyPost(cursor *astutil.Cursor) bool {
	// NOTE Postでは基本的に return false しない panic+recover されて処理がわからなくなるぞ！

	current := cursor.Node()
	if current == nil {
		return true
	}

	switch node := current.(type) {
	case *ast.AssignStmt:
		if len(node.Lhs) == 0 && len(node.Rhs) == 0 {
			// metago.ValueOf を消した結果空になる場合がある
			cursor.Delete()
			return true
		}
	}

	return true
}

func (p *metaProcessor) relateToMetagoPackage(ident *ast.Ident) bool {
	// [mv] ← metago packageですか？ とか
	// mv.[Value]() ← metago packageですか？ とか
	v := p.currentPkg.TypesInfo.Defs[ident]
	if v == nil {
		return false
	}
	t := v.Type()
	tn, ok := t.(*types.Named)
	if !ok {
		return false
	}
	typeName := tn.Obj()
	typePkg := typeName.Pkg()
	if typePkg == nil {
		return false
	}
	if typePkg.Path() != metagoPackagePath {
		return false
	}

	return true
}

func (p *metaProcessor) isMetagoValue(ident *ast.Ident) bool {
	// var ident metago.Value ← true
	// var ident FooBar ← false
	v := p.currentPkg.TypesInfo.Defs[ident]
	if v == nil {
		return false
	}
	t := v.Type()
	tn, ok := t.(*types.Named)
	if !ok {
		return false
	}
	typeName := tn.Obj()
	typePkg := typeName.Pkg()
	if typePkg == nil {
		return false
	}
	if typePkg.Path() != metagoPackagePath {
		return false
	} else if typeName.Name() != valueTypeString {
		return false
	}

	return true
}

func (p *metaProcessor) isMetagoField(ident *ast.Ident) bool {
	// var ident metago.Field ← true
	// var ident FooBar ← false

	v := p.currentPkg.TypesInfo.Defs[ident]
	if v == nil {
		return false
	}
	t := v.Type()
	tn, ok := t.(*types.Named)
	if !ok {
		return false
	}
	typeName := tn.Obj()
	typePkg := typeName.Pkg()
	if typePkg == nil {
		return false
	}
	if typePkg.Path() != metagoPackagePath {
		return false
	} else if typeName.Name() != fieldTypeString {
		return false
	}

	return true
}

func (p *metaProcessor) isMetagoValueOf(selectorExpr *ast.SelectorExpr) bool {
	// selectorExpr == metago.ValueOf
	lhsIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}
	var found bool
	for _, importSpec := range p.currentFile.Imports {
		pkgPath, err := strconv.Unquote(importSpec.Path.Value)
		if err != nil {
			panic(err)
		}
		if pkgPath != metagoPackagePath {
			continue
		}
		if importSpec.Name != nil && importSpec.Name.Name == lhsIdent.Name {
			found = true
			break
		}
		if pkg := p.currentPkg.Imports[pkgPath]; pkg != nil && pkg.Name == lhsIdent.Name {
			found = true
			break
		}
	}
	if !found {
		return false
	}

	if selectorExpr.Sel.Name != valueOfFuncNameString {
		return false
	}

	return true
}

func (p *metaProcessor) isCallMetagoFieldValue(node *ast.CallExpr) bool {
	// mf.Value() 系かどうか
	selectorExpr, ok := node.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	objIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	target := p.fieldMapping[objIdent.Obj]
	if target == nil {
		return false
	}

	if selectorExpr.Sel.Name != fieldValueMethodName {
		return false
	}

	return true
}

func (p *metaProcessor) isCallMetagoFieldName(node *ast.CallExpr) bool {
	// mf.Name() 系かどうか
	selectorExpr, ok := node.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	objIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	target := p.fieldMapping[objIdent.Obj]
	if target == nil {
		return false
	}

	if selectorExpr.Sel.Name != fieldNameMethodName {
		return false
	}

	return true
}

func (p *metaProcessor) isAssignable(expr1 ast.Expr, expr2 ast.Expr) bool {
	if p.isSameType(expr1, expr2) {
		return true
	}

	// TODO ものすごく素晴らしくする
	// time.Time が json.Marshaler に assign できることがわかると素敵なコードが書けるぞ！
	// …といっても p.currentPkg.TypesInfo.Uses とかは astcopy との相性が悪くて死だ！
	// https://golang.org/pkg/go/types/#AssignableTo
	// https://golang.org/pkg/go/types/#Info.TypeOf
	// ↑この辺ちゃう？ってtenntennさんが言ってた

	return false
}

func (p *metaProcessor) isSameType(expr1 ast.Expr, expr2 ast.Expr) bool {
	{
		ident1, ok1 := expr1.(*ast.Ident)
		ident2, ok2 := expr2.(*ast.Ident)
		if ok1 && ok2 {
			if ident1.Obj == nil && ident2.Obj == nil {
				type1 := types.Universe.Lookup(ident1.Name)
				type2 := types.Universe.Lookup(ident2.Name)
				if type1 != nil && type2 != nil {
					return type1 == type2
				} else if type1 == nil && type2 == nil {
					// TODO package の ident の場合の比較が甘い…！
					// import hoge "hoge" と import h "hoge" で hoge と h 比較した時にtrueにならない
					return ident1.Name == ident2.Name
				}
				return false
			} else if ident1.Obj == ident2.Obj {
				return true
			}
			return false
		} else if ok1 {
			return false
		} else if ok2 {
			return false
		}
	}
	{
		sel1, ok1 := expr1.(*ast.SelectorExpr)
		sel2, ok2 := expr2.(*ast.SelectorExpr)
		if ok1 && ok2 {
			if p.isSameType(sel1.X, sel2.X) && p.isSameType(sel1.Sel, sel2.Sel) {
				return true
			}
			return false

		} else if ok1 {
			return false
		} else if ok2 {
			return false
		}
	}

	panic("unreachable")
}

func (p *metaProcessor) isInlineTemplateFuncDecl(cursor *astutil.Cursor, node *ast.FuncDecl) bool {
	var found bool
outer:
	for _, params := range node.Type.Params.List {
		for _, param := range params.Names {
			if p.isMetagoValue(param) {
				found = true
				break outer
			}
		}
	}
	if !found {
		return false
	}

	if node.Recv != nil {
		p.Errorf(node, "method (function with receiver) can't become inline template")
		return true // 子方向の展開処理されるとコンパイルエラーになりがちなので
	}

	return true
}

func (p *metaProcessor) extractMetagoBaseVariable(expr ast.Expr) *ast.Ident {
	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return nil
	}
	selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil
	}

	if !p.isMetagoValueOf(selectorExpr) {
		return nil
	}

	arg1 := callExpr.Args[0]
	targetIdent, ok := arg1.(*ast.Ident)
	if !ok {
		p.Errorf(arg1, "argument must be ident")
		return nil
	}

	return targetIdent
}

func (p *metaProcessor) checkMetagoBuildTagComment(cursor *astutil.Cursor, node *ast.Comment) bool {
	if node.Text != fmt.Sprintf("//+build %s", metagoBuildTag) {
		return false
	}

	p.hasMetagoBuildTag = true
	cursor.Delete()
	return true
}

func (p *metaProcessor) checkReplaceTargetIdent(cursor *astutil.Cursor, node *ast.Ident) bool {
	// mv系の単純な置き換え
	if target := p.valueMapping[node.Obj]; target != nil {
		cursor.Replace(astcopy.Node(target))
		return true
	}
	// mf系の単純な置き換え
	if target := p.fieldMapping[node.Obj]; target != nil {
		field := p.currentTargetField
		cursor.Replace(&ast.SelectorExpr{
			X: target,
			Sel: &ast.Ident{
				Name: field.Name,
				Obj:  field,
			},
		})
		return true
	}

	return false
}

// checkMetagoValueOfAssignStmt is capture `mv := metago.ValueOf(foo)` format assignment.
// it marks up convert rule about `mv` to `foo` and remove these assignment.
func (p *metaProcessor) checkMetagoValueOfAssignStmt(cursor *astutil.Cursor, stmt *ast.AssignStmt) bool {
	// mv := metago.ValueOf(foo) 系を処理する。
	// mv と foo の紐付けを覚える。
	// 該当のassignmentをNode毎削除するようマークする。

	var found bool
	for idx, lhs := range stmt.Lhs {
		ident, ok := lhs.(*ast.Ident)
		if !ok {
			continue
		}
		if !p.isMetagoValue(ident) {
			continue
		}

		rhs := stmt.Rhs[idx]
		targetIdent := p.extractMetagoBaseVariable(rhs)
		if targetIdent == nil {
			// TODO なんらかの警告を出したほうがよさそう
			continue
		}

		found = true
		p.valueMapping[ident.Obj] = &ast.Ident{
			Name: targetIdent.Name,
			Obj:  targetIdent.Obj,
		}

		p.removeNodes[lhs] = true
		p.removeNodes[rhs] = true
	}

	return found
}

func (p *metaProcessor) checkMetagoFieldRange(cursor *astutil.Cursor, node *ast.RangeStmt) bool {
	// for _, mf := range mv.Fields() {
	// ↑的なヤツをサポートしていく

	ident, ok := node.Key.(*ast.Ident)
	if !ok {
		return false
	}
	if ident.Name != "_" {
		p.Warningf(ident, "index part assignment should be '_'")
		return false
	}

	// assignStmt → _, mf := range mv.Fields() 部分相当
	assignStmt, ok := ident.Obj.Decl.(*ast.AssignStmt)
	if !ok {
		return false
	}
	if len(assignStmt.Lhs) != 2 {
		// := の左が2未満だとindexとかしか取ってない
		p.Errorf(assignStmt, "value part assignment must required")
		return false
	}
	// _, mf ← の mf 相当の場所の名前取る
	fieldIdent, ok := assignStmt.Lhs[1].(*ast.Ident)
	if !ok {
		return false
	}

	// range 以外は対応しないめんどいから
	unaryExpr, ok := assignStmt.Rhs[0].(*ast.UnaryExpr)
	if !ok {
		return false
	}
	if unaryExpr.Op != token.RANGE {
		return false
	}
	if len(assignStmt.Rhs) != 1 {
		// 右辺が1以外は知らないパターンだ
		return false
	}

	// mv.Fields() 相当の部分を調べていく
	// とりあえずCallExprじゃなかったら知らないパターン
	callExpr, ok := unaryExpr.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	// mv.Fields の部分調べる
	// mv 部分がが変換候補じゃない場合処理対象ではない
	selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	xIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}
	target, ok := p.valueMapping[xIdent.Obj]
	if !ok {
		return false
	}
	// Fields 部分が Fields 以外だったら処理対象ではない
	if selectorExpr.Sel.Name != valueFieldsMethodName {
		return false
	}

	p.fieldMapping[fieldIdent.Obj] = target

	// 大本のfor句は全部捨てる必要がある
	cursor.Delete()

	// RangeStmtのBody部分をフィールドの数だけコピペする
	var targetObjDef *ast.Object
	switch targetNode := target.(type) {
	case *ast.Ident:
		// TODO 今出てくるパターンを決め打ちでサポートしてるだけなのでいい感じにする
		// ↓は foo *Foo 的なやつの *Foo から Foo の定義を取ろうとしている
		targetObjDef = targetNode.Obj.Decl.(*ast.Field).Type.(*ast.StarExpr).X.(*ast.Ident).Obj
	default:
		panic("unknown type")
	}
	defFields := targetObjDef.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields
	for _, field := range defFields.List {
		for _, name := range field.Names {
			bk := p.currentTargetField

			bodyStmt := astcopy.BlockStmt(node.Body)
			p.currentTargetField = name.Obj
			astutil.Apply(
				bodyStmt,
				p.ApplyPre,
				p.ApplyPost,
			)
			cursor.InsertBefore(bodyStmt)
			for _, labelName := range p.requiredContinueLabel {
				cursor.InsertBefore(&ast.LabeledStmt{
					Label: &ast.Ident{
						Name: labelName,
					},
					Stmt: &ast.EmptyStmt{},
				})
			}
			p.requiredContinueLabel = nil

			p.currentTargetField = bk
		}
	}
	for _, labelName := range p.requiredBreakLabel {
		cursor.InsertAfter(&ast.LabeledStmt{
			Label: &ast.Ident{
				Name: labelName,
			},
			Stmt: &ast.EmptyStmt{},
		})
	}
	p.requiredBreakLabel = nil

	return true
}

func (p *metaProcessor) checkIfStmtInInitWithTypeAssert(cursor *astutil.Cursor, node *ast.IfStmt) bool {
	// if v, ok := mf.Value().(time.Time); ok { ... } 系
	// condがokの参照か !ok か ok == false 以外の場合怒る
	if node.Init == nil {
		return false
	}

	assignStmt, ok := node.Init.(*ast.AssignStmt)
	if !ok {
		return false
	}

	// mf.Value().(time.Time) 的なやつかチェック
	if len(assignStmt.Rhs) != 1 {
		return false
	}
	typeAssertExpr, ok := assignStmt.Rhs[0].(*ast.TypeAssertExpr)
	if !ok {
		return false
	}
	callExpr, ok := typeAssertExpr.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	if !p.isCallMetagoFieldValue(callExpr) {
		return false
	}

	// v, ok := 的なやつかチェック
	if len(assignStmt.Lhs) != 2 {
		p.Errorf(assignStmt, "lhs assignment should be 2")
		return false
	}

	varIdent, ok := assignStmt.Lhs[0].(*ast.Ident)
	if !ok {
		p.Errorf(assignStmt.Lhs[0], "var assignment should be ident")
		return false
	}
	okIdent, ok := assignStmt.Lhs[1].(*ast.Ident)
	if !ok {
		p.Errorf(assignStmt.Lhs[1], "ok assignment should be ident")
		return false
	}

	// IfStmtでスコープに新しい変数が導入されるので置き換えルールを登録
	p.valueMapping[varIdent.Obj] = &ast.SelectorExpr{
		X: p.fieldMapping[callExpr.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Obj],
		Sel: &ast.Ident{
			Name: p.currentTargetField.Name,
			Obj:  p.currentTargetField,
		},
	}

	_, ok = node.Cond.(*ast.Ident)
	// TODO この辺もうちょっと柔軟性をもたせる 静的にboolに還元できる範囲であれば許容してあげたい
	if !ok {
		p.Errorf(node.Cond, "must be '%s'", okIdent.Name)
		return false
	}

	if p.isAssignable(typeAssertExpr.Type, p.currentTargetField.Decl.(*ast.Field).Type) {
		// Bodyが評価される & if全体を置き換え
		astutil.Apply(
			node.Body,
			p.ApplyPre,
			p.ApplyPost,
		)
		cursor.Replace(node.Body)
	} else if node.Else != nil {
		// Elseが評価される & if全体を置き換え
		astutil.Apply(
			node.Else,
			p.ApplyPre,
			p.ApplyPost,
		)
		cursor.Replace(node.Else)
	}

	return true
}

func (p *metaProcessor) checkIfStmtInCondWithTypeAssert(cursor *astutil.Cursor, node *ast.IfStmt) bool {
	// if mf.Value().(time.Time).IsZero() { ... } 系のハンドリング
	// condで、もしvalueの型とassert先がマッチしてたらBlockStmt残し、それ以外は削除

	var typeAssertExpr *ast.TypeAssertExpr
	ast.Walk(astVisitorFunc(func(node ast.Node) bool {
		found, ok := node.(*ast.TypeAssertExpr)
		if ok {
			typeAssertExpr = found
			return false
		}
		return true
	}), node.Cond)
	if typeAssertExpr == nil {
		return false
	}

	callExpr, ok := typeAssertExpr.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	if !p.isCallMetagoFieldValue(callExpr) {
		return false
	}

	targetField := p.currentTargetField
	if targetField == nil {
		p.Errorf(callExpr, "invalid context. not in metago.Field range statement")
		return false
	}

	if !p.isSameType(typeAssertExpr.Type, targetField.Decl.(*ast.Field).Type) {
		cursor.Delete()
		return true // 子をApplyされたくない
	}

	p.replaceNodes[typeAssertExpr] = &ast.SelectorExpr{
		X: p.fieldMapping[callExpr.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Obj],
		Sel: &ast.Ident{
			Name: p.currentTargetField.Name,
			Obj:  p.currentTargetField,
		},
	}

	// 呼び出し元でreturn falseするので必要なモノを自分でApplyしてやる必要がある
	astutil.Apply(
		node.Init,
		p.ApplyPre,
		p.ApplyPost,
	)
	astutil.Apply(
		node.Cond,
		p.ApplyPre,
		p.ApplyPost,
	)
	astutil.Apply(
		node.Body,
		p.ApplyPre,
		p.ApplyPost,
	)
	astutil.Apply(
		node.Else,
		p.ApplyPre,
		p.ApplyPost,
	)

	return true
}

func (p *metaProcessor) checkTypeSwitchStmt(cursor *astutil.Cursor, node *ast.TypeSwitchStmt) bool {
	assignStmt, ok := node.Assign.(*ast.AssignStmt)
	if !ok {
		return false
	}

	// mf.Value().(type) 的なやつかチェック
	if len(assignStmt.Rhs) != 1 {
		return false
	}
	typeAssertExpr, ok := assignStmt.Rhs[0].(*ast.TypeAssertExpr)
	if !ok {
		return false
	}
	if typeAssertExpr.Type != nil {
		p.Errorf(typeAssertExpr, "unknown type assert expr. must be use foo.(type)")
		return false
	}
	callExpr, ok := typeAssertExpr.X.(*ast.CallExpr)
	if !ok {
		return false
	}
	if !p.isCallMetagoFieldValue(callExpr) {
		return false
	}

	// 新しい変数が導入される場合そのマッピングルールを記録
	if len(assignStmt.Lhs) != 1 {
		p.Errorf(assignStmt, "var assignment must required")
		return false
	}
	varIdent, ok := assignStmt.Lhs[0].(*ast.Ident)
	if !ok {
		p.Errorf(assignStmt.Lhs[0], "var assignment should be ident")
		return false
	}
	p.valueMapping[varIdent.Obj] = &ast.SelectorExpr{
		X: p.fieldMapping[callExpr.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Obj],
		Sel: &ast.Ident{
			Name: p.currentTargetField.Name,
			Obj:  p.currentTargetField,
		},
	}

	var targetBody []ast.Stmt // nil と len == 0 を区別する
	for _, stmt := range node.Body.List {
		switch stmt := stmt.(type) {
		case *ast.CaseClause:
			if targetBody == nil && len(stmt.List) == 0 {
				// 長さ 0 は default
				targetBody = stmt.Body
			} else {
				for _, typeExpr := range stmt.List {
					if p.isAssignable(typeExpr, p.currentTargetField.Decl.(*ast.Field).Type) {
						targetBody = stmt.Body
					}
				}
			}

		default:
			panic("unreachable")
		}
	}

	newBlock := &ast.BlockStmt{
		List: targetBody,
	}
	astutil.Apply(newBlock, p.ApplyPre, p.ApplyPost)
	cursor.Replace(newBlock)

	return true
}

func (p *metaProcessor) checkInRangeBranchStmt(cursor *astutil.Cursor, node *ast.BranchStmt) bool {
	switch node.Tok {
	case token.CONTINUE:
		labelName := fmt.Sprintf("metagoGoto%d", p.gotoCounter)
		p.gotoCounter++
		p.requiredContinueLabel = append(p.requiredContinueLabel, labelName)
		cursor.Replace(&ast.BranchStmt{
			Tok: token.GOTO,
			Label: &ast.Ident{
				Name: labelName,
			},
		})
	case token.BREAK:
		labelName := fmt.Sprintf("metagoGoto%d", p.gotoCounter)
		p.gotoCounter++
		p.requiredBreakLabel = append(p.requiredBreakLabel, labelName)
		cursor.Replace(&ast.BranchStmt{
			Tok: token.GOTO,
			Label: &ast.Ident{
				Name: labelName,
			},
		})
	default:
		return false
	}

	return true
}

func (p *metaProcessor) checkInlineTemplateCallExpr(cursor *astutil.Cursor, node *ast.CallExpr) bool {
	// func fooBarTemplate(mv metago.Value, basic, b string) bool 的なやつの変換
	// 第一引数が metago.Value だったら対象

	// foo(mv) 形式のみ対応 メソッド類は対応が大変
	funcName, ok := node.Fun.(*ast.Ident)
	if !ok {
		return false
	}

	if funcName.Obj == nil {
		// panic("foo") とかが該当
		return false
	}

	funcDecl, ok := funcName.Obj.Decl.(*ast.FuncDecl)
	if !ok {
		return false
	}

	// 引数無しは対象外
	if len(funcDecl.Type.Params.List) == 0 {
		return false
	}
	// 引数の最初が metago.Value じゃないものは対象外
	metaValueArg := funcDecl.Type.Params.List[0].Names[0]
	if !p.isMetagoValue(metaValueArg) {
		return false
	}

	// 実引数側の数が0ってことはここまで来たらないだろうけど一応
	if len(node.Args) == 0 {
		return false
	}
	arg, ok := node.Args[0].(*ast.Ident)
	if !ok {
		return false
	}

	funcDecl = astcopy.FuncDecl(funcDecl)

	// 実引数側の mv がマッピングされる先を 仮引数側の mv にも継承させる
	// *ast.Ident#Obj はコピーされないので metaValueArg 取り直さなくても大丈夫
	p.valueMapping[metaValueArg.Obj] = p.valueMapping[arg.Obj]

	// 引数が metago.Value だけならinline展開するやつ
	// goroutineの境界変わったりするとめんどいので即時実行関数で包む
	newCallExpr := &ast.CallExpr{
		Fun: &ast.FuncLit{
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: funcDecl.Type.Params.List[1:], // 先頭は metago.Valueなので
				},
				Results: &ast.FieldList{
					List: funcDecl.Type.Results.List,
				},
			},
			Body: funcDecl.Body,
		},
		Args: node.Args[1:], // 先頭は metago.Valueなので
	}

	// 操作したNodeには入っていってくれないので自分で歩く必要がある
	astutil.Apply(
		newCallExpr,
		p.ApplyPre,
		p.ApplyPost,
	)
	cursor.Replace(newCallExpr)

	return true
}

func (p *metaProcessor) checkUseMetagoFieldValue(cursor *astutil.Cursor, node *ast.CallExpr) bool {
	// mf.Value() 系を obj.Foo 的なのに置き換える
	selectorExpr, ok := node.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	objIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	target := p.fieldMapping[objIdent.Obj]
	if target == nil {
		return false
	}

	if selectorExpr.Sel.Name != fieldValueMethodName {
		return false
	}

	cursor.Replace(&ast.SelectorExpr{
		X: target,
		Sel: &ast.Ident{
			Name: p.currentTargetField.Name,
			Obj:  p.currentTargetField,
		},
	})

	return false
}

func (p *metaProcessor) checkUseMetagoFieldName(cursor *astutil.Cursor, node *ast.CallExpr) bool {
	// mf.Name() 系を "Foo" 的なのに置き換える
	if !p.isCallMetagoFieldName(node) {
		return false
	}

	cursor.Replace(&ast.BasicLit{
		Kind:  token.STRING,
		Value: strconv.Quote(p.currentTargetField.Name),
	})

	return false
}

func (p *metaProcessor) checkUseMetagoStructTagGet(cursor *astutil.Cursor, node *ast.CallExpr) bool {
	// mf.Name() 系を "Foo" 的なのに置き換える
	selectorExpr, ok := node.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	objIdent, ok := selectorExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	target := p.fieldMapping[objIdent.Obj]
	if target == nil {
		return false
	}

	if selectorExpr.Sel.Name != fieldStructTagGetMethodName {
		return false
	}

	if len(node.Args) != 1 {
		p.Errorf(node, "string literal argument must required")
		return false
	}

	basicLit, ok := node.Args[0].(*ast.BasicLit)
	if !ok || basicLit.Kind != token.STRING {
		p.Errorf(node.Args[0], "string literal argument must required")
		return false
	}

	tagName, err := strconv.Unquote(basicLit.Value)
	if err != nil {
		p.Errorf(node.Args[0], "unexpected string literal format. %s: %s", basicLit.Value, err.Error())
		return false
	}

	targetField := p.currentTargetField.Decl.(*ast.Field)
	if targetField.Tag == nil {
		cursor.Replace(&ast.BasicLit{
			Kind:  token.STRING,
			Value: `""`,
		})
		return false
	}
	structTagValue, err := strconv.Unquote(targetField.Tag.Value)
	if err != nil {
		p.Errorf(targetField, "unexpected string literal format. %s: %s", targetField.Tag.Value, err.Error())
		return false
	}

	tagValue := reflect.StructTag(structTagValue).Get(tagName)

	cursor.Replace(&ast.BasicLit{
		Kind:  token.STRING,
		Value: strconv.Quote(tagValue),
	})

	return false
}

func (p *metaProcessor) Debugf(node ast.Node, format string, a ...interface{}) {
	p.nodeErrors = append(p.nodeErrors, &NodeError{
		ErrorLevel: ErrorLevelDebug,
		Fset:       p.currentPkg.Fset,
		Node:       node,
		Message:    fmt.Sprintf(format, a...),
	})
}

func (p *metaProcessor) Noticef(node ast.Node, format string, a ...interface{}) {
	p.nodeErrors = append(p.nodeErrors, &NodeError{
		ErrorLevel: ErrorLevelNotice,
		Fset:       p.currentPkg.Fset,
		Node:       node,
		Message:    fmt.Sprintf(format, a...),
	})
}

func (p *metaProcessor) Warningf(node ast.Node, format string, a ...interface{}) {
	p.nodeErrors = append(p.nodeErrors, &NodeError{
		ErrorLevel: ErrorLevelWarning,
		Fset:       p.currentPkg.Fset,
		Node:       node,
		Message:    fmt.Sprintf(format, a...),
	})
}

func (p *metaProcessor) Errorf(node ast.Node, format string, a ...interface{}) {
	p.nodeErrors = append(p.nodeErrors, &NodeError{
		ErrorLevel: ErrorLevelError,
		Fset:       p.currentPkg.Fset,
		Node:       node,
		Message:    fmt.Sprintf(format, a...),
	})
}
