package config

import (
	"encoding/json"
	"os"
	"time"
)

type ProxyConfig struct {
	IncomingAddress         string        // incoming connection from the internet
	ForwardIncommingAddress string        // outgoing connection to local intermediary or origin server
	OutGoingServerAddress   string        // incoming connection from local intermediary
	OutgoingAddress         string        // outgoing connection to next intermediary
	Whitelisting            bool          // apply whitelisting
	ConnTimeout             time.Duration // connection read and write timeout
	Origin                  bool          // true, if target is origin server, false if target is intermediary with two endpoints
}

func (proxyConfig *ProxyConfig) Load(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, proxyConfig)
	return err
}
