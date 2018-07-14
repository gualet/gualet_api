package main

import (
    "log"
    "net/http"
    "./src/router"
    "./config"
    "os"
)

func main() {

    environment := os.Args[1]

    log.Println("Environment: ", environment)

    conf.NewConfig(environment)

    router := router.NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
