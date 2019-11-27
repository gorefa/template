package k8s

import (
	"fmt"
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
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	var nspods []string
	for _, pod := range pods.Items {
		nspods = append(nspods, pod.Name)
	}

	SendResponse(c, nil, nspods)
}
