package db

import (
	"github.com/jinzhu/gorm"
	"math/rand"
)

type RWManager struct {
	master *gorm.DB
	slaves []*gorm.DB
}

func (r *RWManager) Write() *gorm.DB {
	return r.master
}

func (r *RWManager) Read() *gorm.DB {
	if len(r.slaves) == 0 {
		return r.master
	}
	n := rand.Intn(len(r.slaves)-1)
	return r.slaves[n]
}
