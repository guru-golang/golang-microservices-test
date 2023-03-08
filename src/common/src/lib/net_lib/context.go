package net_lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

// broadcast messages
const (
	UserFindAll = "User.FindAll"
	UserFindOne = "User.FindOne"
	UserCreate  = "User.Create"
	UserUpdate  = "User.Update"
	UserRemove  = "User.Remove"

	UserProfileFindAll = "UserProfile.FindAll"
	UserProfileFindOne = "UserProfile.FindOne"
	UserProfileCreate  = "UserProfile.Create"
	UserProfileUpdate  = "UserProfile.Update"
	UserProfileRemove  = "UserProfile.Remove"

	UserNotificationFindAll = "UserNotification.FindAll"
	UserNotificationFindOne = "UserNotification.FindOne"
	UserNotificationCreate  = "UserNotification.Create"
	UserNotificationUpdate  = "UserNotification.Update"
	UserNotificationRemove  = "UserNotification.Remove"

	WorkerNotificationFindAll = "WorkerNotification.FindAll"
	WorkerNotificationFindOne = "WorkerNotification.FindOne"
	WorkerNotificationCreate  = "WorkerNotification.Create"
	WorkerNotificationUpdate  = "WorkerNotification.Update"
	WorkerNotificationRemove  = "WorkerNotification.Remove"

	ConfFindAll = "Conf.FindAll"
	ConfFindOne = "Conf.FindOne"
	ConfCreate  = "Conf.Create"
	ConfUpdate  = "Conf.Update"
	ConfRemove  = "Conf.Remove"

	PortfolioFindAll = "Portfolio.FindAll"
	PortfolioFindOne = "Portfolio.FindOne"
	PortfolioCreate  = "Portfolio.Create"
	PortfolioUpdate  = "Portfolio.Update"
	PortfolioRemove  = "Portfolio.Remove"
)

type Context struct {
	Conn net.Conn
	Msg  *Message
}

func (c *Context) Emit(data gin.H) (err error) {
	swap := data
	swap["status"] = true
	_, err = fmt.Fprintln(c.Conn, c.Msg.ToStr(swap))
	return
}

func (c *Context) Response(data gin.H) (err error) {
	swap := data
	swap["status"] = true
	_, err = c.Conn.Write(c.Msg.ToByte(swap))
	return
}

func (c *Context) Error(err error) (errr error) {
	swap := gin.H{"status": false, "reason": err.Error()}
	_, errr = c.Conn.Write(c.Msg.ToByte(swap))
	return
}

func (c *Context) Close() error {
	return c.Conn.Close()
}
