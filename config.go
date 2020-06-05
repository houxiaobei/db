package db

import (
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type Config struct {
	Default  string `json:"default" yaml:"default"`
	Clusters map[string]*ClusterConfig
}

type ClusterConfig struct {
	Dialect  string   `json:"dialect"`
	Master   string   `json:"master"`
	Slaves   []string `json:"slaves"`
	MaxConn  int      `json:"max_conn"`
	IdleConn int      `json:"idle_conn"`
}

func (c *Config) Build() (map[string]*RWManager, error) {
	managers := make(map[string]*RWManager)
	for k, v := range c.Clusters {
		master, err := build(v.Dialect, v.Master)
		if err != nil {
			return nil, err
		}
		slaves, err := build(v.Dialect, v.Slaves...)
		if err != nil {
			return nil, err
		}
		managers[k] = &RWManager{
			master: master[0],
			slaves: slaves,
		}
	}
	return managers, nil
}

func build(dialect string, dsn ...string) ([]*gorm.DB, error) {
	dbs := make([]*gorm.DB, len(dsn))
	for k, v := range dsn {
		db, err := gorm.Open(dialect, v)
		if err != nil {
			return nil, err
		}
		dbs[k] = db
	}
	return dbs, nil
}
