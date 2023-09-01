package AutoPoC

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var FOFA_EMAIL = ""
var FOFA_KEY = ""

func SearchByQuery(query string) (results []string) {
	var nextId = ""
	for {
		searchUrl := fmt.Sprintf("https://fofa.info/api/v1/search/next?qbase64=%s&email=%s&key=%s&next=%s",url.PathEscape(base64.StdEncoding.EncodeToString([]byte(query))),FOFA_EMAIL,FOFA_KEY,nextId)
		//fmt.Println(searchUrl)
		resp,err := http.Get(searchUrl)
		if err != nil {
			fmt.Println(err.Error())
		}
		content, _ := ioutil.ReadAll(resp.Body)
		resultsFromQuery ,err := GetResults(content)
		nextId = resultsFromQuery.Next
		results = append(results,resultsFromQuery.Results...)
		//fmt.Println(nextId)
		if nextId == "" {
			break
		}
	}
	fmt.Println(len(results))
	return results
}

func GetResults(respBody []byte) (results QueryResults,err error) {
	var resp map[string]interface{}
	err = json.Unmarshal(respBody,&resp)
	if err != nil {
		return
	}
	//fmt.Println(string(respBody))
	results.Next = resp["next"].(string)
	resp_results := resp["results"].([]interface{})
	for _,result:= range resp_results {
		r := result.([]interface{})
		results.Results = append(results.Results,r[0].(string))
	}
	//results.Results = resp["results"].([]interface{})
	results.Size = int(resp["size"].(float64))
	//fmt.Println("[33] results :",results.Results)
	if len(results.Results) == 0  {
		err = fmt.Errorf("null")
	}
	return
}

func SetKey(email string,key string) {
	FOFA_EMAIL = email
	FOFA_KEY = key
}
