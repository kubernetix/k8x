// See if one can patch typescript
// https://github.com/microsoft/TypeScript/blob/0e292c441a0e5f27e18803128b7dfb1155ac0f5a/src/compiler/transformers/jsx.ts#L218
// https://github.com/sdegutis/imlib/issues/5
// https://github.com/microsoft/TypeScript/issues/21699


// Need to do that to shut up the compiler
// jsx transforms <div> to React.createElement per default
// That means we need to declare the JSX namespace but ALSO a global variable that just returns stuff
declare var React: any

declare namespace JSX {
  interface IntrinsicElements {
    chart: {
      /**
       * This is the name of the chart
       */
      name: string;
      /**
       * This is the version of the chart
       */
      version: string;
      /**
       * This is the version of contained app
       */
      appVersion?: string;
      /**
       * This is the version of the k8s cluster
       */
      kubeVersion?: string;
      description?: string;
      type?: string;
      keywords?: string[];
      home?: URL;
      sources?: string[];

      maintainers?: string[];
      icon?: URL;
      deprecated?: boolean;
      annotations?: string[];
      //children: IntrinsicElements["k8s"];
    };
    k8s: {
      "config-path"?: string;
      //children?: IntrinsicElements["namespace"];
    };
    namespace: {
      name?: string
      //children?: (
      //  | JSX.IntrinsicElements["metadata"]
      //  | JSX.IntrinsicElements["metadata"]
      //)[];
    };
    ingress: {
      //children: JSX.IntrinsicElements["metadata"];
    };
    deployment: {
      "config-path"?: string;
      //children?: IntrinsicElements["namespace"];
    };
    "match-label"?: {
      key: string;
    };
    "match-expression"?: {
      key: string;
      operator: "In" | "Out";
      values: string[];
    };
    selector: {};
    spec: {
      replicas?: number
    } 
    template: {};
    container: {
      image: string;
      imagePullPolicy: "Always" | "Never";
      name: string;
    };
    metadata: {
      name?: string;
      //children?: JSX.IntrinsicElements["annotation"][];
    };
    annotation: { [key: string]: string };
    label: { [key: string]: string };
    service: {
      kind?: "Service";
      apiVersion?: "v1";
    };
    port: {
      name?: string
      protocol: "TCP" | "UDP";
      port: number;
      targetPort: number;
    };
  }
}




