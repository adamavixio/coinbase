package coinbaseclient

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	logger "github.com/adamavixio/logger"
)

type RequestConfig struct {
	Method  string
	Path    string
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

func executeAuthenticatedRequest(config RequestConfig) []byte {
	client := http.DefaultClient
	reader := bytes.NewReader(config.Body)
	address := fmt.Sprintf("%s%s", url, config.Path)

	r, err := http.NewRequest(config.Method, address, reader)
	logger.Error("create auth request error", err)

	if config.Params != nil {
		appendParams(r, config.Params)
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := createMessage(timestamp, config.Method, config.Path, config.Body)
	signed := signMessage(message)

	appendHeaders(r, signed, timestamp, config.Headers)
	res, err := client.Do(r)
	logger.Error("execute client request error", err)

	data, err := ioutil.ReadAll(res.Body)
	logger.Error("request body parsing error", err)

	return data
}

func createMessage(timestamp string, method string, path string, body []byte) string {
	builder := strings.Builder{}
	builder.WriteString(timestamp)
	builder.WriteString(method)
	builder.WriteString(path)
	builder.Write(body)
	return builder.String()
}

func signMessage(message string) string {
	secret := getEnvVar("COINBASE_SECRET")

	key, err := base64.StdEncoding.DecodeString(secret)
	logger.Error("error base64 decoding secret error", err)

	hmac := hmac.New(sha256.New, key)
	_, err = hmac.Write([]byte(message))
	logger.Error("hmac message write error", err)

	sha := base64.StdEncoding.EncodeToString(hmac.Sum(nil))
	return sha
}

func appendHeaders(r *http.Request, signed, timestamp string, headers map[string]string) {
	key := getEnvVar("COINBASE_KEY")
	passphrase := getEnvVar("COINBASE_PASSPHRASE")

	r.Header.Add("Accept", "application/json; charset=utf-8")
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("CB-ACCESS-KEY", key)
	r.Header.Add("CB-ACCESS-SIGN", signed)
	r.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	r.Header.Add("CB-ACCESS-PASSPHRASE", passphrase)

	for k, v := range headers {
		r.Header.Add(k, v)
	}
}

func appendParams(r *http.Request, params map[string]string) {
	q := r.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
}
