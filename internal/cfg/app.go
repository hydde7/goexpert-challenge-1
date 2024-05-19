package cfg

var App = &appConfig{}

type appConfig struct {
	Environment string `json:"APP_ENV"`
	Address     string `json:"APP_ADDRESS"`
	LogLevel    string `json:"APP_LOG_LEVEL"`
}
