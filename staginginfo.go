package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
    "net/http"
	"encoding/json"
    "log"
)
    
var args = parseFlags()
var conf = parseConfig(args)

func CheckIfError(err error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	fmt.Fprintf(w, "", err)
}

func Index(w http.ResponseWriter, r *http.Request, vars httprouter.Params) {
    fmt.Fprintf(w, "Index page.")
}

func ReposInfo(w http.ResponseWriter, r *http.Request, vars httprouter.Params) {

	var info []HeadInfo

	for _, repoPath := range conf.Repos {
		repo := Repo{path: repoPath}

		info = append(info, *repo.GetHeadInfo())
	}

	log.Printf("", info)
    
	jsonString, err := json.Marshal(info)
	CheckIfError(err, w)

	fmt.Fprintf(w, string(jsonString))
}

func main() {


    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/repos/info", ReposInfo)

	host := fmt.Sprintf("localhost:%d", conf.Port)	
    log.Fatal(http.ListenAndServe(host, router))
}
