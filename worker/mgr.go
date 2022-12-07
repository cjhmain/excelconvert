package worker

import (
	"excelconvert/config"
	"fmt"
)

func Start() {
	fmt.Println("worker work start")
	worker := new(Worker)
	worker.Work(config.ExcelPath)
}
