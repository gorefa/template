package k8s

import (
	. "gogin/handler"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodListResponse struct {
	Pods []string
}

func PodList(c *gin.Context) {
	ns := c.Param("ns")
	pods, err := Clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	podstatus := make(map[string]string)
	for _, pod := range pods.Items {
		podstatus[pod.Name] = string(pod.Status.Reason) // ?
	}

	SendResponse(c, nil, podstatus)
}
