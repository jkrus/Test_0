package config

import "time"

type (
	NatsSubscribeOptions struct {
		AllowReconnect bool          `yaml:"allowReconnect"`
		MaxReconnect   int           `yaml:"maxReconnect"`
		ReconnectWait  time.Duration `yaml:"reconnectWait"`
		Timeout        time.Duration `yaml:"timeout"`
	}
)
