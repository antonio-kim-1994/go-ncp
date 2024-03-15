package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestInfo struct {
	Method string
	Path   string
	Query  string
}

type RequestResult struct {
	Timestamp string
	Signature string
}

func GenerateHmac(r RequestInfo) RequestResult {
	space := " "
	newLine := "\n"

	accessKey := os.Getenv("NCP_ACC")
	secretKey := os.Getenv("NCP_SEC")

	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("error loading location:", err)
	}

	timestamp := time.Now().In(loc).UnixNano() / 1e6

	var builder strings.Builder
	builder.WriteString(r.Method)
	builder.WriteString(space)
	builder.WriteString(fmt.Sprintf("%s%s", r.Path, r.Query))
	builder.WriteString(newLine)
	builder.WriteString(strconv.FormatInt(timestamp, 10))
	builder.WriteString(newLine)
	builder.WriteString(accessKey)

	messages := builder.String()

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(messages))

	result := RequestResult{
		Signature: base64.StdEncoding.EncodeToString(h.Sum(nil)),
		Timestamp: strconv.FormatInt(timestamp, 10),
	}

	return result
}
