package common

import (
	"github.com/gorilla/websocket"
)

type ConnClient struct {
	Conn   *websocket.Conn
	UserId int64
}

type globesConn struct {
	DictConn map[int64]*ConnClient
}

func SetConn(event string, connect *websocket.Conn, userId int64) *ConnClient {
	data := &ConnClient{
		Conn:   connect,
		UserId: userId,
	}
	dictConn := MapConn[event].DictConn
	dictConn[userId] = data
	return data
}

func (q *globesConn) DelConn(userId int64) {
	delete(q.DictConn, userId)
}

func (q *ConnClient) GetUser() int64 {
	return q.UserId
}

func (q *ConnClient) GetConn() *websocket.Conn {
	return q.Conn
}

const (
	ListEvent = "list" //聊天列表事件
	ChatEvent = "chat" //聊天触发事件
)

var MapConn = make(map[string]*globesConn)

func InitMapConn() {
	MapConn[ListEvent] = &globesConn{}
	MapConn[ChatEvent] = &globesConn{}
}
