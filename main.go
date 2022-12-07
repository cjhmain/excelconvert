package main

import (
	"excelconvert/machine/factory"
	"excelconvert/worker"
)

func main() {
	factory.New()
	worker.Start()
	factory.Stop()
}
