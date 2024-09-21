package k8s

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func CreateNamespae(name string) {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// options := metav1.CreateOptions{}
	getOpts := metav1.GetOptions{}

	// ns := v1.Namespace()

	namespace, err := clientset.CoreV1().Namespaces().Get(context.Background(), "default", getOpts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deploying to: %s", namespace.Name)
	fmt.Printf("ID: %s", namespace.UID)

}
