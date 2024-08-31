import * as k8s from '@kubernetes/client-node';

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
  const kc = new k8s.KubeConfig();
  kc.loadFromDefault();

  const k8sApi = kc.makeApiClient(k8s.CoreV1Api);

  const namespace = {
    metadata: {
      name: "test",
    },
  };

  try {
    const createNamespaceRes = await k8sApi.createNamespace(namespace);
  } catch (err) { }

  console.log("Hello World");
}

main()