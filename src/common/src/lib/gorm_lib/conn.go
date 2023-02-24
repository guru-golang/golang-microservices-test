package gorm_lib

import (
	"database/sql"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"

	"car-rent-platform/backend/common/src/lib/config_lib"
)

type (
	Interface interface {
		Init() *Gorm
	}

	Conn struct {
		Gorm Gorm `json:"gorm"`
	}

	Conf struct {
		DSN         string      `json:"dsn"`
		Options     gorm.Config `json:"options"`
		ConnOptions ConnOptions `json:"connOptions"`
	}

	ConnOptions struct {
		MaxLifetime  time.Duration
		MaxIdleTime  time.Duration
		MaxOpenConns int
		MaxIdleConns int
	}

	Gorm struct {
		Db   *sql.DB  `json:"db"`
		Base *gorm.DB `json:"base"`
		Conf Conf     `json:"conf"`
	}
)

func NewConn() Interface {
	var i = Conn{}
	return &i
}

func (c *Conn) Init() *Gorm {
	c.Gorm.Conf.Load()
	if db, err := gorm.Open(postgres.Open(c.Gorm.Conf.DSN), &c.Gorm.Conf.Options); err != nil {
		log.Fatal().Msg(err.Error())
	} else {
		c.Gorm.Base = db
	}
	return &c.Gorm
}

func (c *Gorm) Run() error {
	if db, err := c.Base.DB(); err != nil {
		return err
	} else {
		db.SetConnMaxLifetime(c.Conf.ConnOptions.MaxLifetime)
		db.SetConnMaxIdleTime(c.Conf.ConnOptions.MaxIdleTime)
		db.SetMaxOpenConns(c.Conf.ConnOptions.MaxOpenConns)
		db.SetMaxIdleConns(c.Conf.ConnOptions.MaxIdleConns)
		c.Db = db
	}
	return nil
}

func (c *Conf) Load() {
	conf := config_lib.Config.Get("db_gorm").(map[string]any)
	c.DSN = conf["dsn"].(string)
	options := conf["options"].(any)
	connOptions := conf["connOptions"].(any)

	var optionsSwap gorm.Config
	if err := mapstructure.Decode(options, &optionsSwap); err != nil {
		log.Fatal().Msg(err.Error())
	}
	c.Options = optionsSwap
	c.Options.Logger = logger.Default

	var connOptionsSwap ConnOptions
	if err := mapstructure.Decode(connOptions, &connOptionsSwap); err != nil {
		log.Fatal().Msg(err.Error())
	}
	c.ConnOptions = connOptionsSwap
}
