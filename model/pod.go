package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodStat struct {
	Name   string
	Status string
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
