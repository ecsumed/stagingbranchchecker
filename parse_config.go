package main

import (
    "flag"
)



func parseFlags() map[string]string {
    args := make(map[string]string)

    confFile := flag.String("conf", "config.json", "Config file path")
    
    flag.Parse()

    args["confFile"] = *confFile

    return args
}
