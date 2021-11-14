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

func executeAuthenticatedRequest(method string, path string, params map[string]string, body []byte) []byte {
	client := http.DefaultClient
	reader := bytes.NewReader(body)
	address := fmt.Sprintf("%s%s", url, path)

	r, err := http.NewRequest(method, address, reader)
	handleError("create auth request error", err)

	if params != nil {
		appendParams(r, params)
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := createMessage(timestamp, method, path, body)
	signed := signMessage(message)

	appendHeaders(r, signed, timestamp)
	res, err := client.Do(r)
	handleError("execute client request error", err)

	data, err := ioutil.ReadAll(res.Body)
	handleError("request body parsing error", err)

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
	handleError("error base64 decoding secret error", err)

	hmac := hmac.New(sha256.New, key)
	_, err = hmac.Write([]byte(message))
	handleError("hmac message write error", err)

	sha := base64.StdEncoding.EncodeToString(hmac.Sum(nil))
	return sha
}

func appendHeaders(r *http.Request, signed string, timestamp string) {
	key := getEnvVar("COINBASE_KEY")
	passphrase := getEnvVar("COINBASE_PASSPHRASE")

	r.Header.Add("Accept", "application/json; charset=utf-8")
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("CB-ACCESS-KEY", key)
	r.Header.Add("CB-ACCESS-SIGN", signed)
	r.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	r.Header.Add("CB-ACCESS-PASSPHRASE", passphrase)
}

func appendParams(r *http.Request, params map[string]string) {
	q := r.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()
}
