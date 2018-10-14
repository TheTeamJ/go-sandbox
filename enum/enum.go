package main

import (
	"fmt"
)

type ServiceWorkerState int

const (
	Installing ServiceWorkerState = iota
	Installed
	Activating
	Activated
	Redundant
)

func main() {
	fmt.Println(Activating) // 2
}
