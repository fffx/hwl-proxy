package config_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"hwl-proxy/config"
)

func TestLoadWhitelist(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	var proxyConfig config.ProxyConfig
	err := proxyConfig.Load(basepath + "/../test/" + "config.json")
	if err != nil {
		t.Error(err)
	}

}
