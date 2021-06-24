package main

import (
    "fmt"

    "github.com/naduvat/go-module-basic/world-module/world"
)

func main() {
    greeting := world.Greet()
    fmt.Println(greeting)
}
