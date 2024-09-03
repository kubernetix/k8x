# Kubernetix (K8x)
Type safe, reusable kubernetes components

## Example chart

```tsx
import Wordpress from "@charts/wordpress"

const replicas = Number(process.env["VARIABLE"]);

export default () => (
  <cluster config="~/.kube/config">
    <namespace name="default">

      <Wordpress containerName="wordpress"></Wordpress>

      <deployment>
        <spec replicas={replicas}>
          <selector>
            <match-label key="app">wordpress</match-label>
          </selector>
        </spec>

        <template>
          <spec>
            <container
              image="registry.k8s.io/serve_hostname"
              imagePullPolicy="Always"
              name="wordpress"
            ></container>
          </spec>
        </template>
      </deployment>

    </namespace>
  </cluster>
);
```

## Features:

- .env integration
  - K8X_MY_VARIABLE
- Automatic namespace handling
- Chart sharing via npm
  - `npm install -D @charts/wordpress`
  - `import Wordpress from "@charts/wordpress"`
- Better templating
- Statically typed
- IDE support
  ![Proper intellisense support](assets/images/proper_intellisense_support.png "Proper intellisense support")
- Multi cluster definition
- Single installation definition
  - Sometimes a chart is not used to install multiple instances of an app, you just want to install it once in your cluster?
  - Specify `chart.installation` on `package.json` and run `k8x install` without any name
- Reusable components
- Versioning
- Packaging
- Chart/Component Hooks
  - `<Wordpress beforeInstall={slackMessage} afterInstall={slackMessage} onError={handleError} />`

## Usage

```
k8x install my-wordpress
k8x install my-wordpress -u
k8x update my-wordpress
k8x compile
k8x ls
k8x rm
```

## Goals

- Single file app
- Inventory management
- Mature tooling
- Easy sharing
- Easy reusability
- Flexibility

## Non Goals

## Acknowledgements

My experience with real world k8s apps is limited. I work professionally as a developer, so my perspective is biased. I draw a lot of impressions from coding and managing simple applications with docker and docker compose. I am using helm and k8s at work and some of my frustrations might come from inexperience or seem perfectly fine for experienced k8s admins or ops people.

## Terminology

### Component

A piece of tsx code, that is typically a kubernetes entity, for example ingress, pod, service

### Tag

A tsx directive for example <cluster> or <namespace>

## FAQ

### Why JSX/TSX

JSX and TSX already are very mature tooling built directly into the tsc toolchain. IDE support is superior compared to simple yaml or other templating engines. Node has a very mature package and easy accessible code/chart sharing mechanism.

## Helm differentiation

In general k8x is pretty similar to helm. It also took a lot of inspiration from it. But where helm is reinventing the wheel, k8x just falls back to already used mechanisms and infrastructure. (npm/typescript/configuration)

| Topic | helm | k8x |
| -------- | ------- | ------- | 
| Packaging | reinvented | npm |
| Templating | reinvented | tsx |
| Configuration | reinvented | .env |
| Scripting | reinvented | tsx |
| Code sharing | reinvented | tsx |

By reinvented I mean either a custom implementation, or a existing template language with limited or changed features.
