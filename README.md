# Kubernetix (K8x)
Deploy and manage reusable apps with typescript and javascript

## Most basic chart definition

This example shows the most basic, non configurable version of a chart. It closly follows the example yaml definition of
a kubernetes deployment

``touch chart.js``

```js
// https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#creating-a-deployment
const deployment = {
  apiVersion: 'apps/v1',
  kind: 'Deployment',
  metadata: { 
    name: "hello-world-deployment",
    labels: { app: "hello-world" }
  },
  spec: {
    replicas: 1,
    selector: {
      matchLabels: { app: "hello-world" }
    },
    template: {
      metadata: { labels: { app: "hello-world" } },
      spec: {
        containers: [
          { name: 'nginx', image: "nginx:1.14.2", ports: [{ containerPort: 80 }] }
        ]
      }
    }
  }
}

export default () => ({
  name: "hello-world",
  namespace: "default",
  components: [pod],
});
```

Deploy it with:

```
k8x install chart.js
```

## Most basic chart definition with types

You can use typescript with proper type definitions.

``npm install -D @kubernetix/types && mv chart.js chart.ts``

```diff ts
+/// <reference types="@kubernetix/types" />

// https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#creating-a-deployment
- const deployment = {
+ const deployment: k8x.Deployment = {
  apiVersion: 'apps/v1',
  kind: 'Deployment',
  metadata: { 
    name: "hello-world-deployment",
    labels: { app: "hello-world" }
  },
  spec: {
    replicas: 1,
    selector: {
      matchLabels: { app: "hello-world" }
    },
    template: {
      metadata: { labels: { app: "hello-world" } },
      spec: {
        containers: [
          { name: 'nginx', image: "nginx:1.14.2", ports: [{ containerPort: 80 }] }
        ]
      }
    }
  }
}

- export default () => ({
+ export default (): k8x.Chart => ({
  name: "hello-world",
  namespace: "default",
  components: [pod],
});
```

Deploy it with:

```
k8x install chart.ts
```

## Most basic chart definition with types and env integration

Every environment variable prefixed with ``K8X_`` is availabe for the chart to use

```diff ts
/// <reference types="@kubernetix/types" />

// https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#creating-a-deployment
const deployment: k8x.Deployment = {
  apiVersion: 'apps/v1',
  kind: 'Deployment',
  metadata: { 
    name: "hello-world-deployment",
    labels: { app: "hello-world" }
  },
  spec: {
    replicas: 1,
    selector: {
      matchLabels: { app: "hello-world" }
    },
    template: {
      metadata: { labels: { app: "hello-world" } },
      spec: {
        containers: [
          { name: 'nginx', image: "nginx:1.14.2", ports: [{ containerPort: 80 }] }
        ]
      }
    }
  }
}

export default (): k8x.Chart => ({
  name: "hello-world",
-  namespace: "default",
+  namespace: k8x.$env["NAMESPACE"] ?? "default",
  components: [pod],
});
```

Deploy it with:

```
K8X_NAMESPACE=backend-staging k8x install chart.ts
```

## Most basic chart definition with types and env integration and imported ingress component

You can use leverage js imports to reuse components or variables

```diff ts
/// <reference types="@kubernetix/types" />

+import ingress from "./components/ingress"

// https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#creating-a-deployment
const deployment: k8x.Deployment = {
  apiVersion: 'apps/v1',
  kind: 'Deployment',
  metadata: { 
    name: "hello-world-deployment",
    labels: { app: "hello-world" }
  },
  spec: {
    replicas: 1,
    selector: {
      matchLabels: { app: "hello-world" }
    },
    template: {
      metadata: { labels: { app: "hello-world" } },
      spec: {
        containers: [
          { name: 'nginx', image: "nginx:1.14.2", ports: [{ containerPort: 80 }] }
        ]
      }
    }
  }
}

export default (): k8x.Chart => ({
  name: "hello-world",
  namespace: k8x.$env["NAMESPACE"] ?? "default",
-  components: [pod],
+  components: [pod, ingress],
});
```

That's it, deploy your apps with typescript or javascript. Happy deploying!

## Medium complicated chart definition

```ts
/// <reference types="@kubernetix/types" />

import MyIngress, { MyIngressProps } from "./components/ingress"

const values: MyIngressProps = {
  name: k8x.$env["INGRESS_NAME"] ?? "my-ingress",
  appRoot: k8x.$env["INGRESS_NAME"] ?? "/var/www/html",
  additionalPaths:
    Object.keys(k8x.$env)
      .filter((key) => key.startsWith("INGRESS_PATH"))
      .map((key) => k8x.$env[key]) ?? [],
}

export default (): k8x.Chart => ({
  name: "default",
  namespace: "default",
  components: [MyIngress(values)],
})
```

## Features:

- .env integration
  - K8X_MY_VARIABLE
- Automatic namespace handling
  - Auto create/upgrade namespaces
- Sharing
  - `npm install -D @charts/wordpress`
  - `import Wordpress from "@charts/wordpress"`
- Packaging/Versioning
  - `npm version patch -m "Upgrade to 1.0.1 for reasons"`
  - `npm pack @charts/wordpress`
  - `npm publish wordpress.tgz`
- Typescript
- Single binary
- Safe sandboxing
- Proper IDE support
  ![Proper intellisense support](assets/images/proper_intellisense_support.png "Proper intellisense support")
- Single installation definition
  - Specify `chart.name` and run `k8x install` without name parameter
- Interactive chart inspection
  - Load and inspect a file interactively with k8x inspect. It will display all information rendered based on the
  - input it has. 
- Reusable components
  - Props
- Hooks
  - `<Wordpress beforeInstall={slackMessage} afterInstall={slackMessage} onError={handleError} />`
  - `beforeInstall` `afterInstall` `onInstallError` `beforeUpdate` `afterUpdate` `onUpdateError` 

## Usage

```
k8x install <file>
k8x update <file>
k8x inspect <file>
k8x ls
k8x rm
```

## Goals
Reuse existing infrastructure and code features for enhanced developer experience

## Non Goals
- Replace helm

## Helm differentiation

I feel like helm was built by the ops side of devops people. k8x is built by the dev side of devops people.

In general k8x is pretty similar to helm. It also took a lot of inspiration from it. But where helm is reinventing the wheel, k8x just falls back to already used mechanisms and infrastructure. (npm/typescript/configuration)

| Topic | helm     | k8x   |
| -------- |----------|-------| 
| Packaging | custom   | npm   |
| Templating | gotmpl   | js/ts |
| Configuration | --set servers.foo.port=80 | .env  |
| Scripting | custom   | js/ts |
| Code sharing | custom   | js/ts |

By custom I mean either a custom implementation, or a existing template language with limited or changed features.

## Terminology

### Component

A piece of reusable/configurable code, that is typically a kubernetes object, for example ingress, pod, service
