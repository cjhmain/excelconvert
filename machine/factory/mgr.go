package factory

import (
	"excelconvert/config"
	"excelconvert/machine"
	"excelconvert/queue"

	// "fmt"
	"time"
)

type MachineFactory struct {
	queue *queue.Queue
	signal chan *machine.Machiner
	stop chan bool
}

var factory MachineFactory

func New() {
	factory.signal = make(chan *machine.Machiner, 1)
	factory.stop = make(chan bool, 1)
	factory.queue = queue.New()
	for i := 0; i < config.MachinerNum; i++ {
		factory.queue.Push(machine.New(i))
	}

	startSignalListen()

	factory.queue.DebugInfo()
}

func Stop() {
	factory.stop <- true
}

func GetIdleMachiner() *machine.Machiner {
	for {
		m := factory.queue.Poll()
		if m != nil {
			machiner, ok := m.(*machine.Machiner)
			if ok {
				return machiner
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func GetSignalChan() chan *machine.Machiner {
	return factory.signal
}

func releaseBusyMachiner(m *machine.Machiner) {
	factory.queue.Push(m)
}

func startSignalListen() {
	go func() {
		for {
			select {
			case m := <-factory.signal:
				releaseBusyMachiner(m)
			case <- factory.stop:
				return
			default:
				continue
			}
		}
	}()
}

func DebugInfo() {
	factory.queue.DebugInfo()
}