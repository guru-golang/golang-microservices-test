package config_lib

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"car-rent-platform/backend/common/src/lib/json_lib"
)

type Conf struct {
	sync.Mutex
	data map[string]any
}

func NewConf() *Conf {
	return &Conf{}
}

func (c *Conf) Data() map[string]any {
	c.Lock()
	defer c.Unlock()
	return c.data
}

func (c *Conf) Get(k string) any {
	c.Lock()
	defer c.Unlock()
	return c.data[k]
}

func (c *Conf) Set(k string, v any) {
	c.Lock()
	defer c.Unlock()
	c.data[k] = v
}

func (c *Conf) SetData(data map[string]any) {
	c.Lock()
	defer c.Unlock()
	c.data = data
}

var Config *Conf
var config map[string]any

func Init() {
	path, _ := os.Getwd()

	path = strings.ReplaceAll(path, "\\", "/")
	pathSlice := strings.Split(path, "/")
	pathSlice = pathSlice[:len(pathSlice)-2]
	path = strings.Join(pathSlice, "/")

	Config = NewConf()
	config = make(map[string]any)

	readDir(fmt.Sprintf("%v/%v", path, "conf/"))

	Config.SetData(config)

	if config["app_log"] == "debug" {
		//fmt.Println(json_lib.Encode(Config.data))
	}
}

func readDir(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Panic().Msg(err.Error())
	}

	for _, f := range files {
		if f.IsDir() == true {
			readDir(path + f.Name() + "/")
		} else {
			file, err := ioutil.ReadFile(path + f.Name())
			if err != nil {
				log.Panic().Msg(err.Error())
			}
			var data = make(map[string]any)
			if err := json.Unmarshal(file, &data); err != nil {
				log.Panic().Msg(err.Error())
			}
			setConfigs(&data, strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
		}
	}
}

func setConfigs(data *map[string]any, key string) {
	for k, v := range *data {
		skey := key + "_" + k
		swap := os.Getenv(skey)
		config[skey] = v
		switch v.(type) {
		case float64:
			setFloat(skey, swap)
			break
		case bool:
			setBoll(skey, swap)
			break
		case map[string]any:
			var rr = v.(map[string]any)
			if swap != "" {
				rr = json_lib.Decode[map[string]any](make(map[string]any), swap)
				config[skey] = rr
			}
			setConfigs(&rr, skey)
			break
		case []any:
			config[skey] = v.([]any)
			setArr(skey, swap)
			break
		default:
			if swap != "" {
				config[skey] = swap
			}
			break
		}
	}
}

func setFloat(skey, swap string) {
	if swap != "" {
		val, _ := strconv.Atoi(swap)
		config[skey] = float64(val)
	}
}

func setBoll(skey, swap string) {
	if swap != "" {
		val, _ := strconv.ParseBool(swap)
		config[skey] = val
	}
}

func setArr(skey, swap string) {
	if swap != "" {
		config[skey] = json_lib.Decode(*new([]any), swap)
	}
}
