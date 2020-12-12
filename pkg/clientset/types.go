package clientset

import "k8s.io/client-go/kubernetes"

type KubeClient struct {
	Clientset *kubernetes.Clientset
}
