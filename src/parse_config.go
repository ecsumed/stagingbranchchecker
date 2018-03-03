package main

import (
    "flag"
    "encoding/json"
    "os"
    "fmt"
)

type Configuration struct {
    Repos    []string
    Port    int
}

type Args struct {
	conf string
}

// Parse flags
func parseFlags() *Args {

    confFile := flag.String("conf", "config.json", "Config file path")
    
    flag.Parse()

	args := Args{conf: *confFile}

    return &args
}

// Parse config
func parseConfig(args *Args) *Configuration {
    file, _ := os.Open(args.conf)
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)

    if err != nil {
      fmt.Println("error:", err)
    }

    return &configuration
}
