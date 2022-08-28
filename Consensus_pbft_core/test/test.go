package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var counter int = 0

var lastTime time.Duration

var replyMsgs = make([]*ReplyMsg, 0)

type RequestMsg struct {
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"clientID"`
	Operation string `json:"operation"`
}

type ReplyMsg struct {
	ViewID    int64  `json:"viewID"`
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"clientID"`
	NodeID    string `json:"nodeID"`
	Result    string `json:"result"`
}

func main() {
	requestMsg := RequestMsg{
		Timestamp: 102938172,
		ClientID:  "kabigon",
		Operation: "GetMyName",
	}

	startTime := time.Now()
	for i := 0; i < 100; i++ {

		pbytes, _ := json.Marshal(requestMsg)

		data := bytes.NewBuffer(pbytes)

		time.Sleep(time.Millisecond * 14)

		go func() {
			http.Post("http://192.168.0.128:10000/req", "application/json", data)
		}()

		fmt.Printf("%d번째 합의 완료\n", i+1)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("경과 시간 : %f초\n", elapsedTime.Seconds())
}
