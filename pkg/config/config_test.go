package config

import (
	// "errors"
	"fmt"
	// "io/fs"
	"log"
	"os"
	"testing"
)

func TestConfig(t *testing.T){
	cfg := New()	
	if cfg.Environment.Env != "Development" {
		log.Fatal("Error : file found but inccorect env variable value : value : ", cfg.Environment.Env)
	}

	if f, err := os.Stat("./config.yml");  err == nil {
		e := os.Remove(f.Name())
		if e != nil  {
			fmt.Println(e)
		}
	}
}