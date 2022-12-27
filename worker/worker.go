package worker

import (
	"log"

	"excelconvert/factory"
	"io/ioutil"
	"sync"
)

type Worker struct{}

func (w *Worker) Work(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	ch := factory.GetSignalChan()

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		machiner := factory.GetIdleMachiner()
		go machiner.Run(path+file.Name(), &wg, ch)
	}
	wg.Wait()
}
