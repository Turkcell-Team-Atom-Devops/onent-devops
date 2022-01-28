package config

type EnvironmentConfiguration struct {
	Env string `default:"Development"`
}


type LogConfiguration struct {
	Path string 
}

type ServerConfiguration struct {
	Port string
	Timeout int64
}

type Configuration struct {
	Environment EnvironmentConfiguration
	Log			LogConfiguration
	Server	    ServerConfiguration
}