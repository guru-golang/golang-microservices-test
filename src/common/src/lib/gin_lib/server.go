package gin_lib

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"github.com/gin-gonic/gin"
)

type (
	Interface interface {
		Init() *Gin
	}

	Server struct {
		Gin Gin `json:"gin"`
	}

	Conf struct {
		Host string `json:"host"`
	}

	Router struct {
		Base    string                      `json:"base"`
		Group   *gin.RouterGroup            `json:"group"`
		Actions map[string]*gin.HandlerFunc `json:"actions"`
	}

	Gin struct {
		Base   *gin.Engine        `json:"base"`
		Conf   Conf               `json:"conf"`
		Router map[string]*Router `json:"group"`
	}
)

func NewServer() Interface {
	var i = Server{}
	return &i
}

func (s *Server) Init() *Gin {
	s.Gin.Conf.Load()
	s.Gin.Base = gin.New()

	s.Gin.Base.Use(gin.Logger())
	s.Gin.Base.Use(gin.Recovery())

	s.Gin.Router = make(map[string]*Router)
	return &s.Gin
}

func (s *Gin) Route(rPath string) *Router {
	if s.Router[rPath] == nil {
		s.Router[rPath] = new(Router)
	}
	s.Router[rPath].Base = rPath
	s.Router[rPath].Group = s.Base.Group(rPath)
	return s.Router[rPath]
}

func (g *Gin) Run() error {
	return g.Base.Run(g.Conf.Host)
}

func (c *Conf) Load() {
	conf := config_lib.Config.Get("server_gin").(map[string]any)
	c.Host = conf["host"].(string)
}
