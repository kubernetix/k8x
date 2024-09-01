type JsxId = "cluster" | "namespace" | Function | Object;

type JsxFactoryProps = { [key: string]: any };

type JsxTree = {
  id: JsxId;
  children: JsxTree[];
  props: JsxFactoryProps
};

type JsxFactory = (
  id: JsxId,
  props: JsxFactory,
  ...children: JsxTree[]
) => Promise<JsxTree>;

export const jsx: JsxFactory = async (
  id: JsxId,
  props: JsxFactoryProps,
  ...children: JsxTree[]
): Promise<JsxTree> => {
  if (typeof id === "function") {
    return await id(props, ...children);
  }

  const element: JsxTree = { id: id, children: [], props: props };

  children = (await Promise.all(children)).map((v) => v);

  children.forEach((child) => {
    element.children.push(child);
  });

  return element;
};

export const __createFragment = (props, ...children) => {
  return children;
};
