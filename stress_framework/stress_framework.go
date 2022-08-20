package stress_framework

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/rcrowley/go-metrics"
)

const (
	interval time.Duration = 10 * time.Millisecond
)

var (
	printInterval = time.Second
	httpClients   []*http.Client
)

type StressAction func(httpClient *http.Client)

type Logger interface {
	Printf(format string, v ...interface{})
}

func initClient(clientN int) {
	for i := 0; i < clientN; i++ {
		httpClients = append(httpClients, creteHttpClient())
	}
}

func SetPrintInterval(interval time.Duration) {
	printInterval = interval
}

func HttpStressTest(qps int, clientN int, duration time.Duration, logger Logger, action StressAction) {
	initClient(clientN)
	rand.Seed(time.Now().UnixNano())

	stopTicker := time.NewTicker(duration)
	defer stopTicker.Stop()

	runTicker := time.NewTicker(interval)
	defer runTicker.Stop()

	m := metrics.NewTimer()
	metrics.GetOrRegister("Stress Test Record", m)
	go metrics.Log(metrics.DefaultRegistry, printInterval, logger)

	for {
		select {
		case <-stopTicker.C:
			return
		case <-runTicker.C:
			runCount := qps * int(interval) / int(time.Second)
			for i := 0; i < runCount; i++ {
				httpClient := httpClients[rand.Intn(len(httpClients))]
				go func() {
					begin := time.Now()
					action(httpClient)
					m.UpdateSince(begin)
				}()
			}
		}
	}
}

func creteHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost:     500,
			MaxIdleConnsPerHost: 500,
			DisableKeepAlives:   false,
			DisableCompression:  false,
		},
		Timeout: 5 * time.Second,
	}
}
