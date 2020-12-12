package clientset

import (
	"context"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeClient) CreatePod(pod *apiv1.Pod, namespace string) (*apiv1.Pod, error) {
	return k.Clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
}

func (k *KubeClient) RemovePod(name string, namespace string) error {
	return k.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (k *KubeClient) ListPod(namespace string) (*apiv1.PodList, error) {
	return k.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
}
