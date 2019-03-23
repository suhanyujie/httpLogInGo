package tests

import (
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
)

func LogClient(logContent string) (isOk bool,err error) {
	url := "http://127.0.0.1:8002/log"
	payload := strings.NewReader("value="+logContent)
	req, err := http.NewRequest("POST", url, payload)
	if err!= nil {
		return false,err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "1b73d955-fcb6-0440-bdbc-c35725115bcc")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err!= nil {
		return false,err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err!= nil {
		return false,err
	}

	fmt.Println(res)
	fmt.Println(string(body))
	return true,err
}
