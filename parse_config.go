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


// Parse flags
func parseFlags() map[string]string {
    args := make(map[string]string)

    confFile := flag.String("conf", "config.json", "Config file path")
    
    flag.Parse()

    args["confFile"] = *confFile

    return args
}

// Parse config
func parseConfig(args map[string]string) *Configuration {
    file, _ := os.Open(args["confFile"])
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)

    if err != nil {
      fmt.Println("error:", err)
    }

    return &configuration
}
