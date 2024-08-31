# k8x

## Acknowledgements

My experience with real world k8s apps is limited. I work professionally as a developer, so my perspective is biased. I draw a lot of impressions from coding and managing simple applications with docker and docker compose. I am using helm and k8s at work and some of my frustrations might come from inexperience or seem perfectly fine for experienced k8s admins or ops people.

## Features:

- .env integration
    - K8X_MY_VARIABLE
- Automatic namespace handling
- Chart sharing via npm
- Better templating
- Statically typed
- IDE support
- Single File definition
- Multi cluster definition
- Reusable components
- Versioning
- Packaging

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



## Example chart

```


```


## Terminology

### Component

Is a kubernetes entity, for example ingress, pod, service



## FAQ

### Why JSX/TSX
JSX and TSX already are very mature tooling built directly into the tsc toolchain. IDE support is superior compared to simple yaml or other templating engines. Node has a very mature package and easy accessible code/chart sharing mechanism.

