package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunanxiang/charlotter/cache"
)

type WechatMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
}

func HandleWechat(c *gin.Context) {
	// 获取请求参数
	var (
		msgReceive WechatMessage
		msgChan    = make(chan string, 10)
		errChan    = make(chan error, 10)
		msgReply   string
	)
	if err := c.Bind(&msgReceive); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		log.Printf("Bad request, bind params failed, err:%v\n", err)
		return
	}

	// 判断消息类型
	if msgReceive.MsgType != "text" {
		c.String(http.StatusOK, "暂不支持该操作哦，等后续升级~")
		return
	}

	// 先从缓存获取
	val, found := cache.GlobalCache.Get(msgReceive.Content)
	if found {
		msgReply = "正在处理中，请2s后再试~"
		if val != "" {
			msgReply = val.(string)
		}
	} else {
		cache.GlobalCache.Add(msgReceive.Content, "", time.Minute*5)
		// 接入chatGPT处理文本消息
		go Completions(msgReceive.Content, msgChan, errChan)

		// 处理
		select {
		case msgReply = <-msgChan:
		case errReply := <-errChan:
			log.Printf("get reply message error:%v\n", errReply)
			msgReply = "小安开小差了，请你稍后再次请求~"
		case <-time.After(time.Second * 4):
			msgReply = "网络交通堵塞~请你5s后复制粘贴上次问题再次询问获取答案，在这之前请不要重复提问哦~谢谢理解"
		}
	}
	// 构造返回消息
	resp := fmt.Sprintf(`
<xml>
  <ToUserName><![CDATA[%s]]></ToUserName>
  <FromUserName><![CDATA[%s]]></FromUserName>
  <CreateTime>%d</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[%s]]></Content>
</xml>`, msgReceive.FromUserName, msgReceive.ToUserName, msgReceive.CreateTime, msgReply)
	c.String(http.StatusOK, resp)
}
