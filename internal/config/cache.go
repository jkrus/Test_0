package config

type (
	Cache struct {
		Size         int  `yaml:"size"`
		RecoverySize uint `yaml:"recoverySize"`
	}
)
