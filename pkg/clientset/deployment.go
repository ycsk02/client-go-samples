package clientset

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	scalev1 "k8s.io/api/autoscaling/v1"
)

func (k *KubeClient) CreateDeployment(deployment *appsv1.Deployment, namespace string) (*appsv1.Deployment, error) {
	return k.Clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
}

func (k *KubeClient) PatchDeployment(namespace string, name string, pt types.PatchType, data []byte) (*appsv1.Deployment, error) {
	return k.Clientset.AppsV1().Deployments(namespace).Patch(context.TODO(), name, pt, data, metav1.PatchOptions{})
}

func (k *KubeClient) ScalaDeployment(namespace string, name string, scale *scalev1.Scale) (*scalev1.Scale, error) {
	return k.Clientset.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), name, scale, metav1.UpdateOptions{})
}

func (k *KubeClient) GetScaleDeployment(namespace string, name string) (*scalev1.Scale, error) {
	return k.Clientset.AppsV1().Deployments(namespace).GetScale(context.TODO(), name, metav1.GetOptions{})
}
