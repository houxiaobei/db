package db

import (
	"fmt"
	"github.com/houxiaobei/config"
	"github.com/jinzhu/gorm"
)

var defaultCluster string
var managers map[string]*RWManager

type ClusterNotFoundErr struct {
	Default string
	Err     error
}

func (e *ClusterNotFoundErr) Error() string {
	return fmt.Sprintf("config err,cluster %s not found", e.Default)
}

func Init(cfg *Config) (err error) {
	if _, ok := cfg.Clusters[cfg.Default]; !ok {
		return &ClusterNotFoundErr{Default: cfg.Default}
	}
	defaultCluster = cfg.Default
	managers, err = cfg.Build()
	if err != nil {
		return err
	}
	return nil
}

func InitWithConfig(filename string) {
	cnf := &Config{}
	if err := config.Unmarshal(filename, cnf); err != nil {
		panic(err)
	}
	if err := Init(cnf); err != nil {
		panic(err)
	}
}

func Read(cluster ...string) *gorm.DB {
	if len(cluster) == 0 {
		cluster = []string{defaultCluster}
	}

	if c, ok := managers[cluster[0]]; ok {
		return c.Read()
	}
	return nil
}

func Write(cluster ...string) *gorm.DB {
	if len(cluster) == 0 {
		cluster = []string{defaultCluster}
	}

	if c, ok := managers[cluster[0]]; ok {
		return c.Write()
	}
	return nil
}
