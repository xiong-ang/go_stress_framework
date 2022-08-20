# go_stress_framework
> Simple and stupid go stress testing framework~

## Example
```go
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xiong-ang/go_stress_framework/stress_framework"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	stress_framework.HttpStressTest(500, 8, time.Minute, logrus.StandardLogger(), func(httpClient *http.Client) {
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
```

> output
```
time="2022-08-20T22:20:51+08:00" level=info msg="timer Stress Test Record\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  count:           22931\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  min:         301021600.00ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  max:         855084300.00ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  mean:        415062953.70ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  stddev:       65841952.95ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  median:      412611400.00ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  75%:         465312275.00ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  95%:         507566955.00ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  99%:         587732796.00ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  99.9%:       851288484.20ns\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  1-min rate:        966.51\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  5-min rate:        944.41\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  15-min rate:       940.18\n"
time="2022-08-20T22:20:51+08:00" level=info msg="  mean rate:         997.22\n"
```
