package main

import (
	"awesomeProject/transport"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	r := transport.Router()
	fmt.Println(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
