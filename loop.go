package main

import (
    "fmt"
)

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }

    // key value pairs
    kvs := map[string]string{
        "name":    "Daniel Raban",
        "website": "http://danielraban.com",
    }

    for key, value := range kvs {
        fmt.Println(key, value)
    }
}
