package gocurl

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func ApiRequest(method string, url string, headers []string)(response string){

	client := &http.Client{}

	req, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil{
		return err.Error()
	}

	for _, headerContent:= range headers{
		headerContentSplit := strings.Split(headerContent, ":")
		req.Header.Add(headerContentSplit[0], headerContentSplit[1])
	}
	
	resp, err := client.Do(req)
	if err != nil{
		return err.Error()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return err.Error()
	}

	if len(body) == 0{
		return resp.Status
	}

	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, body, "", "\t");err!= nil{
		return err.Error()
	}

	return string(prettyJSON.Bytes())
}
