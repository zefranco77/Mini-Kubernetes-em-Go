package main

import (
    "fmt"
    "time"
)

type Container struct {
    Name string
    Running bool
}

func startContainer(name string) {
    fmt.Printf("Iniciando container: %s\n", name)
    time.Sleep(2 * time.Second)
    fmt.Printf("Container %s rodando\n", name)
}

func main() {
    containers := []string{"web", "db", "cache"}

    for _, c := range containers {
        go startContainer(c)
    }

    fmt.Println("Mini Kubernetes em execução...")
    time.Sleep(5 * time.Second)
}
