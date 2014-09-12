package Sockets

import (
	"encoding/json"

	"github.com/devinceble/BaseServer/Helpers"
	"github.com/igm/pubsub"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

//Message Structure
type Message struct {
	Msg   string `json:"msg"`
	Error int64  `json:"error"`
}

var chat pubsub.Publisher

//Chat Socket Handler
func Chat(session sockjs.Session) {
	Helpers.BaseLog("SOCKET", "NEW", "/echo", "SESSION", 206, 0, nil)
	var closedSession = make(chan struct{})
	Helpers.BaseLog("SOCKET", "INFO", "/echo", "NEW PARTICIPANT", 200, 0, nil)
	defer Helpers.BaseLog("SOCKET", "INFO", "/echo", "LEAVING PARTICIPANT", 304, 1, nil)
	go func() {
		reader, _ := chat.SubChannel(nil)
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				if err := session.Send(msg.(string)); err != nil {
					return
				}
			}
		}
	}()
	for {
		if msg, err := session.Recv(); err == nil {
			var msgr Message
			err := json.Unmarshal([]byte(msg), &msgr)
			if err == nil {
				Helpers.BaseLog("SOCKET", "PUB", "/echo", "PARTIAL CONTENT", 206, 0, err)
			} else {
				Helpers.BaseLog("SOCKET", "PUB", "/echo", "NO CONTENT", 204, 3, err)
			}
			chat.Publish(msg)
			continue
		}
		break
	}
	close(closedSession)
	Helpers.BaseLog("SOCKET", "CLOSED", "/echo", "SESSION", 100, 1, nil)
}
