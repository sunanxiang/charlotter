package server

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sunanxiang/charlotter/config"
)

func Verify(c *gin.Context) {
	// 获取参数
	signature := c.Request.URL.Query().Get("signature")
	timestamp := c.Request.URL.Query().Get("timestamp")
	nonce := c.Request.URL.Query().Get("nonce")
	echoStr := c.Request.URL.Query().Get("echostr")
	if VerifySignature(signature, timestamp, nonce) {
		c.String(http.StatusOK, echoStr)
	} else {
		c.String(http.StatusOK, "")
	}
}

func VerifySignature(signature, timestamp, nonce string) bool {
	sortedStr := SortString(timestamp, nonce, config.WechatToken)
	encodedStr := sha1.Sum([]byte(sortedStr))
	encodedStrStr := hex.EncodeToString(encodedStr[:])
	return signature == encodedStrStr
}

func SortString(a, b, c string) string {
	strs := []string{a, b, c}
	sort.Strings(strs)
	return strings.Join(strs, "")
}
