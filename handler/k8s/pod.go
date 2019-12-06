package k8s

import (
	. "gogin/handler"
	"gogin/model"

	"github.com/gin-gonic/gin"
)

type PodListResponse struct {
	Pods []string
}

func PodList(c *gin.Context) {

	list := model.PodList(c.Param("ns"))

	//ns := c.Param("ns")
	//pods, err := Clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//podstatus := make(map[string]string)
	//for _, pod := range pods.Items {
	//	podstatus[pod.Name] = string(pod.Status.Reason) // ?
	//}

	SendResponse(c, nil, list)
}
