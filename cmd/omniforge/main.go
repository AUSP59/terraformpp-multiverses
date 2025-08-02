package main

import (
    "fmt"
    "os"

    "omniforge/pkg/engine"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: omniforge <command> [args]")
        os.Exit(1)
    }

    command := os.Args[1]

    switch command {
    case "deploy":
        engine.Deploy()
    case "status":
        engine.Status()
    default:
        fmt.Println("Unknown command:", command)
        os.Exit(1)
    }
}