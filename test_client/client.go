package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"go_stress_framework/stress_framework"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	stress_framework.HttpStressTest(1000, 8, time.Minute, logrus.StandardLogger(), func(httpClient *http.Client) {
		reqest := map[string]string{"Msg": "test"}
		reqBody, _ := json.Marshal(&reqest)
		req, err := http.NewRequest("POST", "http://127.0.0.1:8888/echo", bytes.NewReader(reqBody))
		if err != nil {
			logrus.Errorf("request error: %v", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Accept-Encoding", "gzip, deflate")

		resp, err := httpClient.Do(req)
		if err != nil {
			logrus.Errorf("respond error: %v", err)
			return
		}
		defer resp.Body.Close()

		if _, err = ioutil.ReadAll(resp.Body); err != nil {
			logrus.Errorf("respond error: %v", err)
		}

	})
}
