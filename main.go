package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type NsInfo struct {
	Data       []string `json:"data"`
	Httpstatus int64    `json:"httpstatus"`
	Msg        string   `json:"msg"`
}
type Machinelist struct {
	Data []struct {
		ID         string `json:"_id"`
		AgentType  string `json:"agentType"`
		Branch     string `json:"branch"`
		BuildTime  string `json:"buildTime"`
		Commit     string `json:"commit"`
		GoVersion  string `json:"goVersion"`
		Hostname   string `json:"hostname"`
		IP         string `json:"ip"`
		Key        string `json:"key"`
		LastReport string `json:"lastReport"`
		Ns         string `json:"ns"`
		Sleep      string `json:"sleep"`
		Sn         string `json:"sn"`
		Status     string `json:"status"`
		Version    string `json:"version"`
	} `json:"data"`
	Httpstatus int64  `json:"httpstatus"`
	Msg        string `json:"msg"`
}

func main() {
	Secondarynode := GetSecondarynode()
	for i := 0; i < len(Secondarynode); i++ {
		nsmachine := Machinelist{}
		url := fmt.Sprintf("https://registry.monitor.ifengidc.com/api/v1/event/resource?ns=%s&type=machine", Secondarynode[i])
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(body, &nsmachine)
		fmt.Println(Secondarynode[i], len(nsmachine.Data))

	}
}
func GetSecondarynode() []string {
	nsinfo := NsInfo{}
	resp, err := http.Get("http://registry.monitor.ifengidc.com/api/v1/router/ns?format=list")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &nsinfo)

	allns := nsinfo.Data
	var onens []string
	for i := 0; i < len(allns); i++ {
		ns := allns[i]
		test := strings.Split(ns, ".")
		onens = append(onens, test[len(test)-2]+"."+test[len(test)-1])
	}

	return RemoveRepeatedElement(onens)

}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
