package wsKit

import (
	"compress/flate"
	"github.com/gorilla/websocket"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"sync"
)

type (
	Connection struct {
		disposed bool

		// 锁
		lock *sync.Mutex

		// gorilla/websocket的连接
		conn *websocket.Conn

		// 唯一id
		uniqueId string
		// 所属于的群组（有且仅有一个）
		group string
		// 所属于的用户（有且仅有一个）
		user  string
		token string
		data  map[string]interface{}
	}
)

func (c *Connection) Dispose() error {
	if c == nil {
		return nil
	}

	conn := c.conn
	c.conn = nil
	if conn == nil {
		return NoConnError
	}
	return c.conn.Close()
}

func (c *Connection) Push(messageType int, data []byte) error {
	if c == nil {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	ws := c.conn
	if err := ws.WriteMessage(messageType, data); err != nil {
		return err
	}
	return nil
}

func NewConnection(conn *websocket.Conn) (*Connection, error) {
	lock := new(sync.Mutex)
	if conn == nil {
		return nil, errorKit.Simple("conn == nil")
	}

	// 写（推送）的压缩 compress
	conn.EnableWriteCompression(true)
	if err := conn.SetCompressionLevel(flate.BestSpeed); err != nil {
		return nil, err
	}

	return &Connection{
		conn: conn,
		lock: lock,
	}, nil
}