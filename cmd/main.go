package main

import (
	"fmt"
	"github.com/ycsk02/client-go-samples/pkg/clientset"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/utils/pointer"
)

func main() {
	cli, err := clientset.GetClientset()
	if err != nil {
		fmt.Printf("failed to create client to k8s %+v", err)
	}
	kubeclient := clientset.KubeClient{ Clientset: cli }
	_, err = kubeclient.CreateNamespace("sukai2021")
	if err != nil {
		fmt.Printf("failed to create namespace, errors: %+v\n", err)
	}
	namespacelist, err := kubeclient.ListNamespace()
	for _, namespace := range namespacelist.Items {
		fmt.Printf("namespace： %+v\n", namespace.Name)
	}

	// time.Sleep(10 * time.Second)
	// kubeclient.RemoveNamespace("sukai2021")

	pod := &apiv1.Pod{
		TypeMeta:	metav1.TypeMeta{
			Kind: 		"Pod",
			APIVersion:	"v1",
		},
		ObjectMeta:	metav1.ObjectMeta{
			Name: 		"sukai",
			Namespace: 	"sukai2021",
			Labels: 	map[string]string{"creator": "sukai"},
		},
		Spec:		apiv1.PodSpec{
			Containers: []apiv1.Container{
				apiv1.Container{
					Name: 			"nginx",
					Image: 			"nginx",
					ImagePullPolicy: apiv1.PullIfNotPresent,
				},
			},
		},
	}
	_, err = kubeclient.CreatePod(pod, "sukai2021")
	if err != nil {
		fmt.Printf("failed to create pod, error: %+v\n", err)
	}

	podlist, err := kubeclient.ListPod("sukai2021")
	for _, pod := range podlist.Items {
		fmt.Printf("Pod: %+v Status: %+v\n", pod.Name, pod.Status.Phase)
	}

	// err = kubeclient.RemovePod("sukai", "sukai2021")
	// if err != nil {
	// 	fmt.Printf("failed to delete pod, error: %+v\n", err)
	// }

	deployment := &appsv1.Deployment{
		TypeMeta:	metav1.TypeMeta{
			Kind: "Deployment",
			APIVersion: "v1",
		},
		ObjectMeta:	metav1.ObjectMeta{
			Name: "sukai",
			Namespace: "sukai2021",
		},
		Spec: 	appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "sukai",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta:	metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "sukai",
					},
				},
				Spec: 	apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "nginx",
							Image: "nginx",
							ImagePullPolicy: apiv1.PullIfNotPresent,
							Ports: []apiv1.ContainerPort{
								{
									Name: 			"web",
									Protocol: 		apiv1.ProtocolTCP,
									ContainerPort: 	80,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := kubeclient.CreateDeployment(deployment, "sukai2021")
	if err != nil {
		fmt.Printf("failed to create deployment, error: %+v\n", err)
	} else {
		fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	}

	patchTemplate := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"creator": "sukai",
			},
		},
	}
	patchdata, _ := json.Marshal(patchTemplate)
	result, err = kubeclient.PatchDeployment("sukai2021", "sukai", types.StrategicMergePatchType, patchdata)
	if err != nil {
		fmt.Printf("failed to patch deployment, error: %+v\n",err)
	} else {
		fmt.Printf("patched deployment: %+v\n", result.Name)
	}

	scale, err := kubeclient.GetScaleDeployment("sukai2021", "sukai")
	scale.Spec.Replicas = 3
	_, err = kubeclient.ScalaDeployment("sukai2021", "sukai", scale)
	if err != nil {
		fmt.Printf("failed to scale deployment, error: %+v\n",err)
	} else {
		fmt.Printf("scaled deployment")
	}
}

