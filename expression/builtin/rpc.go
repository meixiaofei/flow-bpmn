package builtin

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Rpc(method, url string, m map[string]interface{}, uid ...string) map[string]interface{} {
	body, _ := json.Marshal(m)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	req.Header.Set("Content-Type", "application/json")
	if len(uid) > 0 {
		req.Header.Set("Fcheck-Passthrough-Uid", uid[0])
	}
	//req.Header.Set("Cookie", "X-Branch-Forwarded-For=liufei;")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	} else if res.StatusCode != http.StatusOK {
		return map[string]interface{}{"error": "http status code: " + res.Status}
	} else {
		defer res.Body.Close()
		var m map[string]interface{}
		resBody, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(resBody, &m)
		if err != nil {
			return map[string]interface{}{"error": err.Error()}
		}
		return m
	}
}
