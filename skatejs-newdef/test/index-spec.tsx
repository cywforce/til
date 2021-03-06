{
    customElements.define('x-hello', class extends skate.Component {
        name: string;
        static get props() {
            return {
                name: { attribute: true }
            };
        }
        renderCallback() {
            return skate.h('div', `Hello, ${this.name}`);
        }
    });
}
{ // https://github.com/skatejs/skatejs#counter
    customElements.define('x-counter', class extends skate.Component {
        static get props() {
            return {
                // By declaring the property an attribute, we can now pass an initial value
                // for the count as part of the HTML.
                count: skate.prop.number({ attribute: true })
            };
        }

        count: number;
        intervalID?: number;

        connectedCallback() {
            // Ensure we call the parent.
            super.connectedCallback();

            // We use a symbol so we don't pollute the element's namespace.
            this.intervalID = setInterval(() => ++this.count, 1000);
        }
        disconnectedCallback() {
            // Ensure we callback the parent.
            super.disconnectedCallback();

            // If we didn't clean up after ourselves, we'd continue to render
            // unnecessarily.
            if (this.intervalID) {
                clearInterval(this.intervalID);
            }
        }
        renderCallback() {
            // By separating the strings (and not using template literals or string
            // concatenation) it ensures the strings are diffed indepenedently. If
            // you select "Count" with your mouse, it will not deselect whenr endered.
            return skate.h('div', 'Count ', this.count);
        }
    });
}
{ // https://github.com/skatejs/skatejs#constructor---supersedes-static-created
    customElements.define('my-component', class extends skate.Component {
        constructor() {
            super();
        }
    });
}
{ // https://github.com/skatejs/skatejs#disconnectedcallback---supersedes-static-detached
    customElements.define('my-component', class extends skate.Component {
        disconnectedCallback() {
            super.disconnectedCallback();
        }
    });
}
{ // https://github.com/skatejs/skatejs#attributechangedcallback---supersedes-static-attributechanged
    customElements.define('my-component', class extends skate.Component {
        attributeChangedCallback(name: string, oldValue: any, newValue: any) {
            super.attributeChangedCallback(name, oldValue, newValue);
        }
    });
}
{ // https://github.com/skatejs/skatejs#static-observedattributes
    customElements.define('my-component', class extends skate.Component {
        static get observedAttributes() {
            // return super.observedAttributes.concat('my-attribute');
            return skate.Component.observedAttributes.concat('my-attribute');
        }
    });
}
{ // https://github.com/skatejs/skatejs#static-props
    customElements.define('my-component', class extends skate.Component {
        static get props() {
            return {};
        }
    });
}
{ // https://github.com/skatejs/skatejs#attribute
    customElements.define('my-component', class extends skate.Component {
        static get props() {
            return {
                myProp: { attribute: true }
            };
        }
    });
}
{ // https://github.com/skatejs/skatejs#coerce
    customElements.define('my-component', class extends skate.Component {
        static get props() {
            return {
                myProp: {
                    coerce(value: any) {
                        return value;
                    }
                }
            };
        }
    });
}
{ // https://github.com/skatejs/skatejs#default
    customElements.define('my-component', class extends skate.Component {
        static get props() {
            return {
                myProp: {
                    default: 'default value'
                }
            };
        }
    });

    customElements.define('my-component', class B extends skate.Component {
        static get props() {
            return {
                myProp: {
                    default(elem: B, data: string) {
                        return [];
                    }
                }
            };
        }
    });
}
{ // https://github.com/skatejs/skatejs#deserialize
    customElements.define('my-component', class extends skate.Component {
        static get props() {
            return {
                myProp: {
                    deserialize(value: string) {
                        return value.split(',');
                    }
                }
            };
        }
    });
}
{ // https://github.com/skatejs/skatejs#get
    customElements.define('my-component', class A extends skate.Component {
        static get props(): { [name: string]: skate.PropAttr<any, any>; } {
            return {
                myProp: {
                    get(elem: A, data: any) {
                        return `prefix_${data.internalValue}`;
                    }
                }
            };
        }
    });
}
{ // https://github.com/skatejs/skatejs#initial
    customElements.define('my-component', class extends skate.Component {
        static get props() {
            return {
                myProp: {
                    initial: 'initial value'
                }
            };
        }
    });

    customElements.define('my-component', class extends skate.Component {
        static get props(): { [name: string]: skate.PropAttr<any, string>; } {
            return {
                myProp: {
                    initial(elem, data) {
                        return 'initial value';
                    }
                }
            };
        }
    });
}
{ // https://github.com/skatejs/skatejs#serialize
    customElements.define('my-component', class extends skate.Component {
        static get props(): { [name: string]: skate.PropAttr<any, string[]>; } {
            return {
                myProp: {
                    serialize(value: string[]) {
                        return value.join(',');
                    }
                }
            };
        }
    });

    customElements.define('my-component', class extends skate.Component {
        static get props(): { [name: string]: skate.PropAttr<any, any>; } {
            return {
                myProp: {
                    set(elem: any, data: any) {
                        // do something
                    }
                }
            };
        }
    });
}
{ // https://github.com/skatejs/skatejs#prototype
    customElements.define('my-component', class extends skate.Component {
        get someProperty() { return 1; }
        set someProperty(v: number) { }
        someMethod() { }
    });
}
{ // https://github.com/skatejs/skatejs#updatedcallback---supersedes-static-updated
    customElements.define('x-component', class extends skate.Component implements skate.OnUpdatedCallback {
        updatedCallback(previousProps: any) {
            // The previous props will not be defined if it is the initial render.
            if (!previousProps) {
                return true;
            }

            // The previous props will always contain all of the keys.
            for (let name in previousProps) {
                if (previousProps[name] !== (this as any)[name]) {
                    return true;
                }
            }
        }
    });

    class Elem extends skate.Component implements skate.OnRenderCallback {
        static get props() {
            return {
                str: skate.prop.string(),
                arr: skate.prop.array()
            }
        }

        str: string;
        arr: string[];

        renderCallback() {
            return skate.h('div', 'testing');
        }
    }

    customElements.define('x-element', Elem);

    const elem = new Elem();

    // Re-renders:
    elem.str = 'updated';

    // Will not re-render:
    elem.arr.push('something');

    // Will re-render:
    elem.arr = elem.arr.concat('something');

    function myCustomCheck(el: skate.Component, prev: any): boolean {
        return true;
    }

    customElements.define('my-component', class extends skate.Component implements skate.OnUpdatedCallback {
        updatedCallback(prev: any) {
            // You can reuse the original check if you want as part of your new check.
            // You could also call it directly if not extending: skate.Component().
            return super.updated(prev) && myCustomCheck(this, prev);
        }
    });

    customElements.define('my-component', class extends skate.Component implements skate.OnUpdatedCallback {
        static get props() {
            return {
                name: skate.prop.string()
            };
        }

        name: string;

        updatedCallback(prev: any) {
            if (prev.name !== this.name) {
                skate.emit(this, 'name-changed', { detail: prev });
            }
        }
    });
}
{ // https://github.com/skatejs/skatejs#rendercallback---supersedes-static-render
    customElements.define('my-component', class extends skate.Component implements skate.OnRenderCallback {
        renderCallback() {
            return skate.h('p', `My name is ${this.tagName}.`);
        }
    });

    customElements.define('my-component', class extends skate.Component implements skate.OnRenderCallback {
        renderCallback() {
            return [
                skate.h('paragraph 1'),
                skate.h('paragraph 2'),
            ];
        }
    });
}
{ // https://github.com/skatejs/skatejs#renderedcallback---supersedes-static-rendered
    // NONE
}
{ // https://github.com/skatejs/skatejs#emit-elem-eventname-eventoptions--
    customElements.define('x-tabs', class extends skate.Component implements skate.OnRenderCallback {
        renderCallback() {
            return skate.h('x-tab', { onSelect: () => { } });
        }
    });

    customElements.define('x-tab', class extends skate.Component implements skate.OnRenderCallback {
        renderCallback() {
            return skate.h('a', { onClick: () => skate.emit(this, 'select') });
        }
    });
}
{ // https://github.com/skatejs/skatejs#preventing-bubbling-or-canceling
    let elem: skate.Component = null as any;
    skate.emit(elem, 'event', {
        composed: false,
        bubbles: false,
        cancelable: false
    });
}
{ // https://github.com/skatejs/skatejs#passing-data
    let elem: skate.Component = null as any;
    skate.emit(elem, 'event', {
        detail: {
            data: 'my-data'
        }
    });
}
{ // https://github.com/skatejs/skatejs#link-elem-propspec
    customElements.define('my-input', class extends skate.Component {
        static get props() {
            return {
                value: { attribute: true }
            };
        }
        renderCallback() {
            return skate.h('input', { onChange: skate.link(this), type: 'text' });
        }
    });

    customElements.define('my-input', class extends skate.Component {
        static get props() {
            return {
                value: { attribute: true }
            };
        }
        renderCallback() {
            skate.h('input', { name: 'someValue', onChange: skate.link(this), type: 'text' });

            skate.link(this, 'someValue');

            skate.link(this, 'obj.someValue');

            skate.h('input', { name: 'someValue', onChange: skate.link(this, 'obj.'), type: 'text' });

            const linkage = skate.link(this, 'obj.');
            skate.h('input', { name: 'someValue1', onChange: linkage, type: 'text' });
            skate.h('input', { name: 'someValue2', onChange: linkage, type: 'checkbox' });
            skate.h('input', { name: 'someValue3', onChange: linkage, type: 'radio' });
            skate.h('select', { name: 'someValue4', onChange: linkage },
                skate.h('option', { value: 2 }, 'Option 2'),
                skate.h('option', { value: 1 }, 'Option 1'),
            );
        }
    });
}
{ // https://github.com/skatejs/skatejs#prop
    skate.prop.boolean();

    const myNewProp = skate.prop.create({});
    myNewProp({});

    skate.prop.boolean({
        coerce() {
            // coerce it differently than the default way
        },
        set() {
            // do something when set
        }
    });
}
{ // https://github.com/skatejs/skatejs#props-elem-props
    const { define, props } = skate;

    class Elem extends skate.Component {
        static get props() {
            return {
                prop1: {}
            };
        }
    }
    customElements.define('my-element', Elem);
    const elem = new Elem();

    // Set any property you want.
    props(elem, {
        prop1: 'value 1',
        prop2: 'value 2'
    });

    // Only returns props you've defined on your component.
    // { prop1: 'value 1' }
    props(elem);
}
{ // https://github.com/skatejs/skatejs#ready-element-callback
    // NONE
}
{ // https://github.com/skatejs/skatejs#h
    customElements.define('my-component', class extends skate.Component implements skate.OnRenderCallback {
        renderCallback() {
            return skate.h('p', { style: { fontWeight: 'bold' } }, 'Hello!');
        }
    });
}
{ // https://github.com/skatejs/skatejs#jsx
    customElements.define('my-component', class extends skate.Component {
        renderCallback() {
            return <p>Hello!</p>;
        }
    });
}
{ // https://github.com/skatejs/skatejs#other-ways-to-use-jsx
    customElements.define('my-component', class extends skate.Component {
        static get props() {
            return {
                title: skate.prop.string()
            };
        }
        renderCallback() {
            return (
                <div>
                    <h1>{this.title}</h1>
                    <slot name="description" />
                    <article>
                        <slot />
                    </article>
                </div>
            );
        }
    });
}
{ // https://github.com/skatejs/skatejs#vdombuilder-
    const {vdom} = skate;
    const h = vdom.builder();
    customElements.define('my-component', class extends skate.Component {
        renderCallback() {
            return h('div', { id: 'test', }, h('p', 'test'));
        }
    });
}
{ // https://github.com/skatejs/skatejs#vdombuilder-elements
    const [div, p] = skate.vdom.builder('div', 'p');
    customElements.define('my-component', class extends skate.Component {
        renderCallback() {
            return div({ id: 'mydiv' }, p('test'));
        }
    });
}
{ // https://github.com/skatejs/skatejs#component-constructor
    class MyElement extends skate.Component { }
    customElements.define('my-element', MyElement);

    // Renders <my-element />
    skate.h(MyElement);

    skate.vdom.elementOpen(MyElement);

    // for https://github.com/Microsoft/TypeScript/issues/7004
    const anyProps: any = {};
    <MyElement {...anyProps} />;
}
{ // https://github.com/skatejs/skatejs#function-helper
    {
        const MyElement = () => skate.h('div', 'Hello, World!');

        // Renders <div>Hello, World!</div>
        skate.h(MyElement);
    }
    {
        const MyElement = (props: any) => skate.h('div', `Hello, ${props.name}!`);

        // Renders <div>Hello, Bob!</div>
        skate.h(MyElement, { name: 'Bob' });
    }
    {
        const MyElement = (props: any, chren: any) => skate.h('div', 'Hello, ', chren, '!');

        // Renders <div>Hello, Mary!</div>
        skate.h(MyElement, 'Mary');
    }
    {
        const MyElement = (props: any, chren: any) => <div>Hello, {chren}!</div>;

        // Renders <div>Hello, Mary!</div>
        <MyElement>Mary</MyElement>
    }
}
{ // https://github.com/skatejs/skatejs#special-attributes
    {
        skate.h('ul',
            skate.h('li', { key: 0 }),
            skate.h('li', { key: 1 }),
        );
    }

    {
        const onClick = console.log;
        skate.h('button', { onClick });
        skate.h('button', { 'on-click': onClick });

        skate.h('button', { onclick: onClick });
    }

    {

        customElements.define('my-element', class extends skate.Component {
            constructor() {
                super();
                this.addEventListener('change', this.handleChange);
            }

            handleChange(e: any) {
                // `this` is the element.
                // The event is passed as the only argument.
            }
        });
    }

    {
        const ref = (button: HTMLButtonElement) => button.addEventListener('click', console.log);
        skate.h('button', { ref });
    }
    {
        const ref = console.log;
        customElements.define('my-element', class extends skate.Component {
            renderCallback() {
                return skate.h('div', { ref });
            }
        });
    }
    {
        customElements.define('my-element', class extends skate.Component {
            renderCallback() {
                const ref = console.log;
                return skate.h('div', { ref });
            }
        });
    }
    {
        skate.h('div', { ref: (e: HTMLElement) => (e.innerHTML = '<p>oh no you didn\'t</p>'), skip: true });
    }
    {
        skate.h('div', { statics: ['attr1', 'prop2'] });
    }
}
