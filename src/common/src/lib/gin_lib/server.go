package gin_lib

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"flag"
	"fmt"
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
		Host    string `json:"host"`
		Version string `json:"version"`
	}

	Router struct {
		Base  string           `json:"base"`
		Group *gin.RouterGroup `json:"group"`
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
	flag.Parse()
	s.Gin.Conf.Load()

	if config_lib.Config.Get("app_log").(string) == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

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
	appName := config_lib.Config.Get(fmt.Sprintf("app_name")).(string)
	fmt.Println("appName", appName)
	conf := config_lib.Config.Get(fmt.Sprintf("services_%v_gin", appName)).(map[string]any)
	c.Host, c.Version = conf["host"].(string), conf["version"].(string)
}
