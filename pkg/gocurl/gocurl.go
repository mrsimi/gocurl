package gocurl

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

type LoginModel struct{
	Username string 
	Password string 
	Channel string
}


func ApiRequest(method string, url string, headers []string, requestBody string)(response string){
	
	//requestBody = `{"username":"wends10","password":"Alat001!","channel":1}`
	//requestBody = `{"title":"Post title", "body": "Post description", "userId":1}`
	t := &http.Transport{
		Dial: (&net.Dialer{
				Timeout: 60 * time.Second,
				KeepAlive: 30 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 60 * time.Second,
}
	client := &http.Client{
		Transport: t,
	}
	
	var body io.Reader = nil	
	if len(requestBody) > 0{

		requestJsonBody := []byte(``+requestBody+``)
		body = bytes.NewBuffer(requestJsonBody)
		
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, body)
	if err != nil{
		return err.Error()
	}

	for _, headerContent:= range headers{
		headerContentSplit := strings.Split(headerContent, ":")
		req.Header.Add(headerContentSplit[0], headerContentSplit[1])
	}

	//add default contentype
	if len(headers) == 0{
		req.Header.Add("Content-Type", "application/json")
	}
	
	resp, err := client.Do(req)
	if err != nil{
		return err.Error()
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil{
		return err.Error()
	}

	if len(responseBody) == 0{
		return resp.Status
	}

	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, responseBody, "", "\t");err!= nil{
		return err.Error()
	}

	return string(prettyJSON.Bytes())
}
