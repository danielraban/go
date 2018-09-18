package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    URL := "https://api.github.com/users/danielraban"
    res, err := http.Get(URL)
    if err != nil {
        log.Println(err)
    }
    defer res.Body.Close() // for garbage collection

    responseBodyBytes, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Println(err)
    }

    fmt.Println(string(responseBodyBytes))

}
