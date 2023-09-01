package AutoPoC

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func ReadPocFromFile(filePath string) []byte {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func MakeRequest(text string,Ips []string) {
	for _,v := range Ips {
		pocReq := AnalyzeRequest(text)
		body,_ := ioutil.ReadAll(pocReq.Body)
		target := fmt.Sprintf("http://%s%s",v,pocReq.URL.Path)
		req, _ := http.NewRequest(pocReq.Method,target, strings.NewReader(string(body)))
		req.Header = pocReq.Header
		req.Host = pocReq.Host
		req.ContentLength = int64(len(body))
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		respBytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(respBytes))
	}

}

func AnalyzeRequest(text string) *http.Request {

	reqbuf := bufio.NewReader(strings.NewReader(text))

	myReq, err := http.ReadRequest(reqbuf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return myReq

}


func traverse(target interface{}) {
	sVal := reflect.ValueOf(target)
	sType := reflect.TypeOf(target)
	if sType.Kind() == reflect.Ptr {
		//用Elem()获得实际的value
		sVal = sVal.Elem()
		sType = sType.Elem()
	}
	num := sVal.NumField()
	for i := 0; i < num; i++ {
		f := sType.Field(i)
		val := sVal.Field(i).Interface()
		fmt.Printf("%5s %v = %v\n", f.Name, f.Type, val)
	}
}
