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
)

type RequestConfig struct {
	Method  string
	Path    string
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

func executeAuthenticatedRequest(config RequestConfig) ([]byte, error) {
	client := http.DefaultClient
	reader := bytes.NewReader(config.Body)
	address := fmt.Sprintf("%s%s", url, config.Path)

	r, err := http.NewRequest(config.Method, address, reader)
	if err != nil {
		return nil, err
	}

	if config.Params != nil {
		appendParams(r, config.Params)
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := createMessage(timestamp, config.Method, config.Path, config.Body)

	signed, err := signMessage(message)
	if err != nil {
		return nil, err
	}

	appendHeaders(r, signed, timestamp, config.Headers)

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func createMessage(timestamp string, method string, path string, body []byte) string {
	builder := strings.Builder{}
	builder.WriteString(timestamp)
	builder.WriteString(method)
	builder.WriteString(path)
	builder.Write(body)
	return builder.String()
}

func signMessage(message string) (string, error) {
	secret, err := getEnvVar("COINBASE_SECRET")
	if err != nil {
		return "", err
	}

	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	hmac := hmac.New(sha256.New, key)

	_, err = hmac.Write([]byte(message))
	if err != nil {
		return "", err
	}

	sha := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	return sha, nil
}

func appendHeaders(r *http.Request, signed, timestamp string, headers map[string]string) error {
	key, err := getEnvVar("COINBASE_KEY")
	if err != nil {
		return err
	}

	passphrase, err := getEnvVar("COINBASE_PASSPHRASE")
	if err != nil {
		return err
	}

	r.Header.Add("Accept", "application/json; charset=utf-8")
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("CB-ACCESS-KEY", key)
	r.Header.Add("CB-ACCESS-SIGN", signed)
	r.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	r.Header.Add("CB-ACCESS-PASSPHRASE", passphrase)

	for k, v := range headers {
		r.Header.Add(k, v)
	}

	return nil
}

func appendParams(r *http.Request, params map[string]string) {
	q := r.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
}
