// 返り値の型は自動的にいい感じに推論される
async function* counter() /* Generator<number, string, boolean> */ {
    console.log("Start!");
    let i = 0;
    while (true) {
        // ここの変数の型指定は必要 next の引数の型の推論に利用されるため
        const v = yield i;
        if (v) {
            break;
        }
        i++;
    }
    return "done!";
}

(async () => {
    let iter = counter();
    console.log("ready?");

    // 最初の yeild までを実行
    let curr = await iter.next();

    while (!curr.done) {
        console.log(curr.value);
        // whileループ内では curr.value は number とわかっている
        curr = await iter.next(curr.value === 3);

        // next の引数もチェックされる
        // error TS2345: Argument of type '[123]' is not assignable to parameter of type '[] | [boolean]'.
        // iter.next(123);

        // 残念ながらこれはvalid
        // [] or [boolean] を受け付けるため
        // iter.next();
    }

    // ループの外では curr.done === true なので curr.value は string とわかっている
    console.log(curr.value.toUpperCase());

    // 次のような出力になる
    // ready?
    // Start!
    // 0
    // 1
    // 2
    // 3
    // DONE!
})()
    .catch(err => console.error(err));

export { }
