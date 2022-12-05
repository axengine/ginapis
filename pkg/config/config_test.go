package config

import (
	"github.com/BurntSushi/toml"
	"github.com/bbdshow/bkit/conf"
	"os"
	"testing"
)

func TestConfigToStdout(t *testing.T) {
	cfg := &Config{}
	if err := conf.UnmarshalDefaultVal(cfg); err != nil {
		t.Fatal(err)
	}
	toml.NewEncoder(os.Stdout).Encode(cfg)
}
