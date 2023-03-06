package net_lib

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net"
	"time"
)

type (
	RpcInterface interface {
		Service(name string) *RpcClient
	}
	Rpc struct {
	}
	RpcResult struct {
		Err  *string `json:"err,omitempty"`
		Body gin.H   `json:"body"`
	}
	RpcClient struct {
		pool    *ConnPool
		current net.Conn
		conf    Conf
	}
)

func New() RpcInterface {
	var i Rpc

	return &i
}

func (r *Rpc) Service(name string) *RpcClient {
	conf := config_lib.Config.Get(fmt.Sprintf("services_%v_tcp", name)).(map[string]any)

	client := new(RpcClient)
	client.conf.Host, client.conf.Pool = conf["host"].(string), int(conf["pool"].(float64))
	client.pool, _ = NewConnPool(r.factory(client.conf.Host), client.conf.Pool)

	return client
}

func (r *Rpc) factory(addr string) func() (net.Conn, error) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)

	return func() (conn net.Conn, err error) {
		conn, err = net.DialTimeout("tcp", tcpAddr.String(), time.Second*2)
		return
	}
}

func (rc *RpcClient) Send(cmd string, data gin.H) (res gin.H, err error) {
	conn, err := rc.pool.Get()
	if err := conn.SetReadDeadline(time.Unix(int64(time.Second*2), 0)); err != nil {
		return nil, err
	}
	defer rc.release(conn)

	ctx := Context{Conn: conn, Msg: &Message{}}
	if err != nil {
		return nil, err
	}

	if err = ctx.Response(gin.H{"cmd": cmd, "body": data}); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(conn)
	for {
		if err = decoder.Decode(&res); err != nil {
			return nil, err
		}
		break
	}

	return rc.response(res)
}

func (rc *RpcClient) Emit(cmd string, data gin.H) (res gin.H, err error) {
	if err = rc.lazyLoad(); err != nil {
		return nil, err
	}

	ctx := Context{Conn: rc.current, Msg: &Message{}}
	if err != nil {
		return nil, err
	}

	if err = ctx.Emit(gin.H{"cmd": cmd, "body": data}); err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(rc.current)
	for {
		if err = decoder.Decode(&res); err != nil {
			return nil, err
		}
		break
	}

	return rc.response(res)
}

func (rc *RpcClient) lazyLoad() (err error) {
	if rc.current == nil {
		rc.current, err = rc.pool.Get()
		err = rc.current.SetReadDeadline(time.Unix(int64(time.Second*2), 0))
	}
	return
}

func (rc *RpcClient) Release() (err error) {
	defer func() { rc.current = nil }()
	return rc.pool.Put(rc.current)
}

func (rc *RpcClient) release(conn net.Conn) {
	if err := rc.pool.Put(conn); err != nil {
		log.Error().Msgf("RPC-error TCP dialer release error: %v", err.Error())
	}
}

func (rc *RpcClient) response(data gin.H) (res gin.H, err error) {
	if data["status"] == nil {
		return nil, errors.New("status is missing")
	}
	if data["status"].(bool) == true {
		return data, nil
	} else {
		fmt.Println(data)
		return nil, errors.New(data["reason"].(string))
	}
}
