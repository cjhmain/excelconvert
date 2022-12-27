package main

import (
	"excelconvert/factory"
	"excelconvert/worker"
)

func main() {
	factory.New()
	worker.Start()
	factory.Stop()
}
