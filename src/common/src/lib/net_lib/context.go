package net_lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

// broadcast messages
const (
	UserFindAll = "User.FideAll"
	UserFindOne = "User.FindOne"
	UserCreate  = "User.Create"
	UserUpdate  = "User.Create"
	UserRemove  = "User.Remove"
)

type Context struct {
	Conn net.Conn
	Msg  Message
}

func (c *Context) SendEmit(data gin.H) (err error) {
	swap := data
	swap["status"] = true
	_, err = fmt.Fprintln(c.Conn, c.Msg.ToStr(swap))
	return
}

func (c *Context) SendResp(data gin.H) (err error) {
	swap := data
	swap["status"] = true
	_, err = c.Conn.Write(c.Msg.ToByte(swap))
	return
}

func (c *Context) SendErr(err error) (errr error) {
	swap := gin.H{"status": false, "reason": err.Error()}
	_, errr = c.Conn.Write(c.Msg.ToByte(swap))
	return
}

func (c *Context) Close() error {
	return c.Conn.Close()
}
