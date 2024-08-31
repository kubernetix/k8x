/// <reference no-default-lib="true"/>

declare var React: any;

declare namespace k8x {

  type Chart = {
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

  type Cluster = {
    config?: string;
    //children?: IntrinsicElements["namespace"];
  };
  type Namespace = {
    name?: string;
    //children?: (
    //  | JSX.IntrinsicElements["metadata"]
    //  | JSX.IntrinsicElements["metadata"]
    //)[];
  };
  type Ingress = {
    //children: JSX.IntrinsicElements["metadata"];
  };
  type Deployment = {
    "config-path"?: string;
    //children?: IntrinsicElements["namespace"];
  };
  type MatchLabel = {
    key: string;
  };
  type MatchExpression = {
    key: string;
    operator: "In" | "Out";
    values: string[];
  };
  type Selector = {}
  type Spec = {
    replicas?: number;
  };
  type Template = {}
  type Container = {
    image: string;
    imagePullPolicy: "Always" | "Never";
    name: string;
  };
  type Metadata = {
    name?: string;
    //children?: JSX.IntrinsicElements["annotation"][];
  };
  type Annotation = { [key: string]: string };
  type Label = Annotation
  type Service = {
    kind?: "Service";
    apiVersion?: "v1";
  };
  type Port = {
    name?: string;
    protocol: "TCP" | "UDP";
    port: number;
    targetPort: number;
  };

}

declare namespace JSX {
  interface IntrinsicElements {
    /**
     * Root tag for a k8s application chart
     * 
     * @deprecated Use package.json instead
     * 
     * @example
     * export default () => (
     *   <chart name="wordpress" version="1.0.0" appVersion="6.6.1"></chart>
     * )
     */
    chart: k8x.Chart
    /**
     * @description Use the <cluster> tag to deploy to one or more namespaces at the same time
     * @optional
     */
    cluster: k8x.Cluster
    /*
     * OPTIONAL: Use the <namespace> tag to deploy to one or more namespaces at the same time
     */
    namespace: k8x.Namespace
    ingress: k8x.Ingress
    deployment: k8x.Deployment
    "match-label"?: k8x.MatchLabel
    "match-expression"?: k8x.MatchExpression
    selector: k8x.Selector
    spec: k8x.Spec
    template: k8x.Template
    container: k8x.Container
    metadata: k8x.Metadata
    annotation: k8x.Annotation
    label: k8x.Label
    service: k8x.Service
    port: k8x.Port
  }
}
