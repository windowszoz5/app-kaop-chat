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

var MapConn = globesConn{}

func SetConn(conn *websocket.Conn, userId int64) *ConnClient {
	data := &ConnClient{
		Conn:   conn,
		UserId: userId,
	}
	MapConn.DictConn[userId] = data
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
