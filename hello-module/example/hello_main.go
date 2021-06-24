package main

import (
    "fmt"

    "github.com/naduvat/go-module-basic/hello-module/hello"
)

func main() {
    greeting := hello.Greet()
    fmt.Println(greeting)
}
