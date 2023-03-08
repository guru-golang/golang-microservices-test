package proxy

import (
	"car-rent-platform/backend/common/src/lib/builtin_lib"
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/proxy/src/tasks/common"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Backend struct {
	Path    *regexp.Regexp
	Service *httputil.ReverseProxy
}

type Frontend struct {
	Backend
	Header *regexp.Regexp
}

type Services struct {
	sync.Mutex
	Frontends map[string]*Frontend
	Backends  map[string]*Backend
}

func (i *Services) GetFrontends() map[string]*Frontend {
	i.Lock()
	defer i.Unlock()
	return i.Frontends
}

func (i *Services) GetBackends() map[string]*Backend {
	i.Lock()
	defer i.Unlock()
	return i.Backends
}

type Task struct {
	common.Task
	services Services

	limiter *rate.Limiter
}

func New() common.TaskInterface {
	var i = Task{}
	i.limiter = rate.NewLimiter(rate.Limit(time.Second*1), 100)
	i.parseConfig()
	return &i
}

func (i *Task) Start() error {
	defer builtin_lib.Recovery()

	port := config_lib.Config.Get("services_proxy_port").(string)
	log.Info().Msgf("Starting proxy %v", port)

	http.HandleFunc("/", i.RequestHandler())
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal().Err(err)
	}
	return nil
}

func (i *Task) Stop() error {
	return nil
}

func (i *Task) NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	var tr = httputil.NewSingleHostReverseProxy(u)
	//tr.Transport = http_lib.DT()
	return tr, nil
}

func (i *Task) RequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.ToLower(r.URL.Path)
		ua := strings.ToLower(r.Header.Get("User-Agent"))
		backends := i.services.GetBackends()
		frontends := i.services.GetFrontends()

		for key, backend := range backends {
			if backend.Path.MatchString(path) {
				log.Debug().Msgf("Request:backend destination: %v -> service: %v -> base path: :%v", path, key, backend.Path.String())
				backend.Service.ServeHTTP(w, r)
				return
			}
		}
		for key, frontend := range frontends {
			if frontend.Header.MatchString(ua) {
				log.Debug().Msgf("Request:backend destination: %v -> service: %v", path, key)
				frontend.Service.ServeHTTP(w, r)
				return
			}
		}
	}
}

func (i *Task) parseConfig() {
	backends, ok := config_lib.Config.Get("services_backends").([]any)
	if ok {
		i.services.Backends = make(map[string]*Backend)
		for _, backend := range backends {
			if service, ok := config_lib.Config.Get(fmt.Sprintf("services_%v_gin", backend)).(map[string]any); ok {
				proxySrv, _ := i.NewProxy(service["host"].(string))
				bInstance := new(Backend)
				bInstance.Path = regexp.MustCompile(service["version"].(string))
				bInstance.Service = proxySrv
				i.services.Backends[backend.(string)] = bInstance
				log.Debug().Msgf("[PROXY-debug] backend-service: %v, host: %v, version: %v", backend, service["host"].(string), service["version"].(string))
			}
		}
	}

	frontends, ok := config_lib.Config.Get("services_frontends").([]any)
	if ok {
		i.services.Frontends = make(map[string]*Frontend)
		for _, frontend := range frontends {
			if service, ok := config_lib.Config.Get(fmt.Sprintf("services_%v_gin", frontend)).(map[string]any); ok {
				proxySrv, _ := i.NewProxy(service["host"].(string))
				fInstance := new(Frontend)
				fInstance.Header = regexp.MustCompile(service["header"].(string))
				fInstance.Path = regexp.MustCompile(service["path"].(string))
				fInstance.Service = proxySrv
				i.services.Frontends[frontend.(string)] = fInstance
				log.Debug().Msgf("[PROXY-debug] frontend-service: %v, host: %v, version: %v", frontend, service["host"].(string), service["version"].(string))
			}
		}
	}
}
