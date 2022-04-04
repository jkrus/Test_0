package config

type (
	// DB represent database configuration.
	DB struct {
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Name    string `yaml:"name"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		SSLMode string `yaml:"ssl-mode"`
	}
)
