import * as k8s from '@kubernetes/client-node';
import * as esbuild from 'esbuild'
import fs from "node:fs/promises"

import { jsx } from './jsx/jsx-factory';

// https://github.com/oclif/oclif

// Todo start by implementing k8x chart.tsx => json

// Idea, replace nodejs with go and use k8s-go, esbuild-go and transpile the esbuild result to gocode via https://github.com/owenthereal/godzilla

// Transpile step
  // transpile
  // trycatch code
  // load tree from jsx factory
// Configure cluster
  // if configured, use different cluster config
  // test cluster config with SELECT 1
  // idea: infer what permissions needed from jsx and check if they are there.
// Configure namespace
  // Check if namespace exists
  // Create if not (opt out)
// Create non operational k8s entities
  // Create ingress
  // Create network policies
  // Create services
// Create operational k8s entities
  // Create pods
  // Create deployments

async function main() {

  // use esbuild.transform instead

  const chart = await fs.readFile("examples/wordpress/chart.tsx")

  const code = await esbuild.transform(chart, {
    loader: 'tsx',
    jsxFactory: "jsx",
    jsxFragment: "jsxFragment",
    jsxImportSource: "k8x",
    platform: "node",
    target: "node18"
  })

  const src = code.code.replace("export default () =>", "export default (jsx) =>")
  let func = await import(`data:text/javascript, ${src}`);

  const struct = await func.default(jsx)

  const kc = new k8s.KubeConfig();
  kc.loadFromDefault();

  const ns = struct.children[0].props.name

  const k8sApi = kc.makeApiClient(k8s.CoreV1Api);

  const namespace = {
    metadata: {
      name: ns,
    },
  };

  try {
    const createNamespaceRes = await k8sApi.createNamespace(namespace);
    const readNamespaceRes = await k8sApi.readNamespace(ns);
  } catch (err) {
    console.error(err)
  }
  
}

main()