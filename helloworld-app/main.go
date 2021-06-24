package main

import (
    "fmt"

    "github.com/naduvat/go-module-basic/hello-module/hello"
    "github.com/naduvat/go-module-basic/world-module/world"
)

func main() {
    g1 := hello.Greet()
    g2 := world.Greet()
    fmt.Printf("%s %s\n", g1, g2)
}
