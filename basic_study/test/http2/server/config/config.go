package config

// 程序配置
type Config struct {
	ServicePort         int    `json:"servicePort"`
	ServiceReadTimeout  int    `json:"serviceReadTimeout"`
	ServiceWriteTimeout int    `json:"serviceWriteTimeout"`
	ServerPem           string `json:"serverPem"`
	ServerKey           string `json:"serverKey"`
}

func GetConfig() *Config {
	return &Config{
		ServicePort:         8080,
		ServiceReadTimeout:  3000, // ms
		ServiceWriteTimeout: 3000, // ms
		ServerKey:           "./config/default.key",
		ServerPem:           "./config/default.pem",
	}
}
