package luvcheckerlib

import (
	"fmt"
	"sync"
)

type Worker struct {
	Work func()
}

//NewWorker returns a new Worker
func NewWorker() Worker {
	return Worker{}
}

//AddWork function adds a new work item to the worker's work queue
func (w *Worker) AddWork(f func()) {
	w.Work = f
}

//StartWorker starts the worker and performs the work assigned to it
func (w *Worker) StartWorker() {

	//This is the reciever task that the worker will perform on every check
	go func() {
		for {
			i := <-OnCheck
			switch i {
			case 1: // Worker Received a Hit Task
				HitsNumber++

				//Add ui update here

			case 2: // Worker Received a Fail Task
				FailsNumber++

				//Add ui update here
			}
		}
	}()

	WaitGroup = sync.WaitGroup{}

	//Prompt the user to enter a number of bots to start
	fmt.Print("Enter the number of bots to start: ")

	//Read the number of bots to start
	fmt.Scanln(&BotCount)

	Semaphore = make(chan int, BotCount)

	WaitGroup.Add(BotCount)
	for i := 0; i < BotCount; i++ {
		Semaphore <- 0
		go w.WorkerFunc(i)

	}

	WaitGroup.Wait()
}

//WorkerFunc is the function that the worker will perform
func (w *Worker) WorkerFunc(id int) {
	defer func() {
		WaitGroup.Done()
		<-Semaphore
	}()

	fmt.Printf("Worker %d starting\n", id)
	w.Work()
	fmt.Printf("Worker %d done\n", id)
}
