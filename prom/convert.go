package prom

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	dto "github.com/prometheus/client_model/go"

	"github.com/prometheus/prom2json"
)

func Promjson(urlpath string) (jsontext string, err error) {
	var input io.Reader

	if url, urlErr := url.Parse(urlpath); urlErr != nil || url.Scheme == "" {
		if input, err = os.Open(urlpath); err != nil {
			return "", err
		}
	}

	mfChan := make(chan *dto.MetricFamily, 1024)

	if input != nil {
		go func() {
			if err := prom2json.ParseReader(input, mfChan); err != nil {
				fmt.Fprintln(os.Stderr, "error reading metrics:", err)
			}
		}()
	} else {
		transport, err := makeTransport("", "", true)
		if err != nil {
			return "", err
		}
		go func() {
			err := prom2json.FetchMetricFamilies(urlpath, mfChan, transport)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}()
	}

	result := []*prom2json.Family{}
	for mf := range mfChan {
		result = append(result, prom2json.NewFamily(mf))
	}
	jsonText, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(jsonText), nil
}

func makeTransport(certificate string, key string, skipServerCertCheck bool) (*http.Transport, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DisableKeepAlives = true
	transport.ResponseHeaderTimeout = time.Minute
	tlsConfig := &tls.Config{InsecureSkipVerify: skipServerCertCheck}
	if certificate != "" && key != "" {
		cert, err := tls.LoadX509KeyPair(certificate, key)
		if err != nil {
			return nil, err
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}
	transport.TLSClientConfig = tlsConfig
	return transport, nil
}
