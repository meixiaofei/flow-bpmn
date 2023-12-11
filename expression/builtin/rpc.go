package builtin

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func FcheckRpc(method, url string, body []byte, uid ...string) (bool, map[string]interface{}) {
	if ok, m := Rpc(method, url, body, uid...); ok {
		if success, ok := m["result"].(bool); ok && success {
			return true, m
		} else {
			return false, m
		}
	} else {
		return false, nil
	}
}

func Rpc(method, url string, body []byte, uid ...string) (bool, map[string]interface{}) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return false, nil
	}
	req.Header.Set("Content-Type", "application/json")
	if len(uid) > 0 {
		req.Header.Set("_fcheck_passthrough_uid", uid[0])
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return false, nil
	} else {
		defer res.Body.Close()
		var m map[string]interface{}
		resBody, _ := io.ReadAll(res.Body)
		err = json.Unmarshal(resBody, &m)
		if err != nil {
			return false, nil
		}
		return true, m
	}
}
