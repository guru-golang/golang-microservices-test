package net_lib

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/common/src/lib/json_lib"
	"car-rent-platform/backend/common/src/lib/reflect_lib"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
	"net"
	"reflect"
	"runtime"
	"sync"
)

type (
	HandlerFunc func(ctx *Context)

	ServerInterface interface {
		Init() *Net
	}

	Server struct {
		Net  *Net `json:"net"`
		Conf Conf `json:"conf"`
	}

	Conf struct {
		Host string `json:"host"`
		Pool int    `json:"pool"`
	}

	Net struct {
		addr     *net.TCPAddr
		Base     net.Listener
		patterns *patterns
	}
	patterns struct {
		mx   sync.Mutex
		data map[string]HandlerFunc
	}
	Message struct {
		m    sync.Mutex
		cmd  string
		body string
		data any
	}
)

func NewServer() ServerInterface {
	var i Server
	return &i
}

func (p *patterns) GetPattern(ptrn string) HandlerFunc {
	p.mx.Lock()
	defer p.mx.Unlock()
	return p.data[ptrn]
}

func (p *patterns) SetPatter(ptrn string, hf HandlerFunc) {
	p.mx.Lock()
	defer p.mx.Unlock()
	p.data[ptrn] = hf
}

func (s *Server) Init() *Net {
	s.Conf.load()

	s.Net = new(Net)
	s.Net.patterns = &patterns{data: map[string]HandlerFunc{}}
	if addr, err := net.ResolveTCPAddr("tcp", s.Conf.Host); err != nil {
		log.Panic().Msgf("Net: Wrong address %v, err: %v", addr, err)
	} else {
		log.Debug().Msgf("Net: Resolve %v", addr)
		s.Net.addr = addr
	}

	return s.Net
}

func (n *Net) Run() (err error) {
	for cmd, pattern := range n.patterns.data {
		name := runtime.FuncForPC(reflect.ValueOf(pattern).Pointer()).Name()
		log.Debug().Msgf("RPC-debug %v --> %v", cmd, name)
	}

	if n.Base, err = net.Listen("tcp", n.addr.String()); err != nil {
		return err
	}

	defer func(l net.Listener) {
		if err := l.Close(); err != nil {
			log.Error().Msgf("RPC-error TCP listener closed error: %v", err.Error())
		}
	}(n.Base)

	for {
		if conn, err := n.Base.Accept(); err != nil {
			break
		} else {
			go n.handleRequest(conn)
		}
	}
	return err
}

func (n *Net) handleRequest(conn net.Conn) {
	reader := json.NewDecoder(conn)
	for {
		var res gin.H
		if err := reader.Decode(&res); err != nil {
			log.Error().Msgf("RPC-error TCP listener decode error: %v", err.Error())
			break
		} else {
			n.handlePattern(res, conn)
		}
	}
}

func (n *Net) Pattern(cmd string, hf HandlerFunc) {
	n.patterns.SetPatter(cmd, hf)
}

func (n *Net) handlePattern(data gin.H, conn net.Conn) {
	if data["cmd"] == nil || data["body"] == nil {
		return
	}
	cmd, body := data["cmd"].(string), json_lib.Encode(data["body"])

	if hf := n.patterns.GetPattern(cmd); hf == nil {
		log.Error().Msgf(`RPC-error Non-existent RPC message pattern:"%v"`, cmd)
	} else {
		log.Debug().Msgf(`RPC-debug RPC message pattern cmd:"%v"`, cmd)
		hf(&Context{Conn: conn, Msg: &Message{cmd: cmd, body: body}})
	}

}

func (m *Message) ShouldBind(obj any) error {
	return binding.JSON.BindBody([]byte(m.body), obj)
}

func (m *Message) Data() any {
	if m.data == nil {
		json_lib.Decode[any](&m.data, m.body)
	}
	return m.data
}

func (m *Message) Prop(name string) (reflect.Value, error) {
	return reflect_lib.Prop(m.Data(), name)
}

func (m *Message) PropString(name string) (*string, error) {
	return reflect_lib.PropString(m.Data(), name)
}

func (m *Message) PropInt(name string) (*int, error) {
	return reflect_lib.PropInt(m.Data(), name)
}

func (m *Message) PropUint(name string) (*uint, error) {
	return reflect_lib.PropUint(m.Data(), name)
}

func (m *Message) PropBoll(name string) (*bool, error) {
	return reflect_lib.PropBoll(m.Data(), name)
}

func (m *Message) ToByte(obj gin.H) []byte {
	return []byte(m.ToStr(obj))
}

func (m *Message) ToStr(obj gin.H) string {
	return json_lib.Encode(obj)
}

func (c *Conf) load() {
	appName := config_lib.Config.Get(fmt.Sprintf("app_name")).(string)
	conf := config_lib.Config.Get(fmt.Sprintf("services_%v_tcp", appName)).(map[string]any)
	c.Host = conf["host"].(string)
}
