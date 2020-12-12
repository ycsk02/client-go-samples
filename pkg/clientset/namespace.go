package clientset

import (
	"context"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeClient) CreateNamespace(name string) (*apiv1.Namespace, error) {
	namespace := &apiv1.Namespace{
		TypeMeta:	metav1.TypeMeta{
			Kind: "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	return k.Clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
}

func (k *KubeClient) RemoveNamespace(name string) error {
	return k.Clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (k *KubeClient) ListNamespace() (*apiv1.NamespaceList, error) {
	return k.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
}

