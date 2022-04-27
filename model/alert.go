package model

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Alerts struct {
	Alerts []struct {
		Annotations struct {
			Content string `json:"content"`
		} `json:"annotations"`
		EndsAt       string `json:"endsAt"`
		Fingerprint  string `json:"fingerprint"`
		GeneratorURL string `json:"generatorURL"`
		Labels       struct {
			AlertName string `json:"alertName"`
		} `json:"labels"`
		StartsAt string `json:"startsAt"`
		Status   string `json:"status"`
	} `json:"alerts"`
	CommonAnnotations interface{} `json:"commonAnnotations"`
	CommonLabels      interface{} `json:"commonLabels"`
	ExternalURL       string      `json:"externalURL"`
	GroupLabels       interface{} `json:"groupLabels"`
	Receiver          string      `json:"receiver"`
	Status            string      `json:"status"`
}

func Alertmessage(alert *Alerts, url string) error {

	client := &http.Client{}

	msgStr := fmt.Sprintf(`
	{
		"msgtype": "text",
		"text": {
		  "content": %s
		}
	  }
	`, alert.Alerts[0].Annotations.Content)

	jsonStr := []byte(msgStr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

	return nil
}
