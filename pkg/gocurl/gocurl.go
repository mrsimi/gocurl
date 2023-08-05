package gocurl

import(
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetRequest(url string)(response string){
	resp, err := http.Get(url)
	if err != nil{
		return err.Error()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return err.Error()
	}

	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, body, "", "\t");err!= nil{
		return err.Error()
	}

	return string(prettyJSON.Bytes())
}
