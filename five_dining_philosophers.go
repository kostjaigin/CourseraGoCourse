package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	name            string
	number          int
	leftCS, rightCS *ChopStick
	eatCounter      int
	eatSignals      <-chan int
	ateSignals      chan<- int
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	for p.eatCounter != 3 {
		// request eating from the host
		<-p.eatSignals

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("%s starting to eat %d\n", p.name, p.number)
		p.eatCounter += 1
		time.Sleep(time.Second)
		fmt.Printf("%s finishing eating %d\n", p.name, p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		// signal that this philosopher ate
		p.ateSignals <- 1
	}
	// signal that this philosopher is done eating 3 times
	wg.Done()
}

// eatSignals is channel of size two
func host(eatSignals chan<- int, ateSignals <-chan int, abort <-chan int) {
	// host allows no more than two philosophers to eat concurrently
	eatSignals <- 1
	eatSignals <- 1

	// wait for someone to be fed to send new eat signal
	for {
		select {
		case <-ateSignals:
			eatSignals <- 1
		case <-abort:
			return
		}
	}
}

func main() {
	// initialize chopsticks
	ChopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}

	// communication and synchronization means
	var eatingDone sync.WaitGroup
	eatingDone.Add(5)
	eatSignals := make(chan int, 2)
	ateSignals := make(chan int, 2)

	// initialize philosophers
	names := []string{"Plato", "Aristotle", "Nietzsche", "Kant", "Descartes"}
	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{
			names[i],
			i,
			ChopSticks[i],
			ChopSticks[(i+1)%5],
			0,
			eatSignals,
			ateSignals,
		}
	}

	// initialize host
	// up to 5 phisolophers can wait to be allowed to eat
	hostAbortChannel := make(chan int, 1)
	go host(eatSignals, ateSignals, hostAbortChannel)

	// start the dining
	for i := 0; i < 5; i++ {
		go philos[i].eat(&eatingDone)
	}

	systemDoneChannel := make(chan int, 1)

	// wait for philos to be done and signal this
	go func() {
		eatingDone.Wait()
		hostAbortChannel <- 1
		systemDoneChannel <- 1
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-shutdown:
		fmt.Println("Shutting down on termination command")
	case <-systemDoneChannel:
		fmt.Println("All philosoperhers are fed")
	}

}
