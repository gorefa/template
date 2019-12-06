package model

import (
	"github.com/lexkong/log"
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
		panic(err.Error())
	}
	var ingresses []string
	for _, ing := range ingress.Items {
		ingresses = append(ingresses, ing.Name)
	}

	return ingresses
}
