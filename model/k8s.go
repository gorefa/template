package model

import (
	"github.com/lexkong/log"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodStat struct {
	Name   string
	Status string
}

type Ingress struct {
	NS   string
	Name string
}

func PodList(ns string) []PodStat {

	pods, err := Clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	podstatus := []PodStat{}

	for _, pod := range pods.Items {
		var item PodStat
		item.Name = pod.Name
		item.Status = string(pod.Status.Phase)
		podstatus = append(podstatus, item)

	}
	return podstatus
}

func IngressList(ns string) []string {
	ingress, err := Clientset.ExtensionsV1beta1().Ingresses(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf(err, "list %s ingress err", ns)
	}
	var ingresses []string
	for _, ing := range ingress.Items {
		ingresses = append(ingresses, ing.Name)
	}

	return ingresses
}

func DeploymentCreate(ns string) (string, error) {
	deploymentsClient := Clientset.AppsV1().Deployments(ns)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	log.Info("Create deployment....")
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		log.Fatalf(err, "create deployment failed")
		return "", err
	}
	return result.GetObjectMeta().GetName(), nil
}

func int32Ptr(i int32) *int32 { return &i }

func NodeList() []string {
	nodelist, err := Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf(err, "list node err")
	}

	var nodes []string
	for _, node := range nodelist.Items {
		nodes = append(nodes, node.GetName())
	}
	return nodes
}
