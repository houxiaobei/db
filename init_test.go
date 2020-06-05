package db

import (
	"github.com/houxiaobei/config"
	"testing"
)

func TestInitWithConfig(t *testing.T) {
	config.Init(config.FromFile())
	InitWithConfig("test.yml")

	if err := Read().DB().Ping(); err != nil {
		t.Error(err)
	}

	if err := Write().DB().Ping(); err != nil {
		t.Error(err)
	}
}
