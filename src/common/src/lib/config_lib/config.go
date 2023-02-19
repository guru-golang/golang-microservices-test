package config_lib

import (
	"encoding/json"
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

	Config = NewConf()
	config = make(map[string]any)

	readDir(path + "/src/configs/")

	Config.SetData(config)

	if config["app_log"] == "debug" {
		//fmt.Println(json_lib.Encode(Config.data))
	}
}

func readDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() == true {
			readDir(path + f.Name() + "/")
		} else {
			file, err := ioutil.ReadFile(path + f.Name())
			if err != nil {
				panic(err)
			}
			var data = make(map[string]any)
			if err := json.Unmarshal(file, &data); err != nil {
				panic(err)
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
				rr = json_lib.Decode(make(map[string]any), swap).(map[string]any)
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
