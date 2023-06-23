package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ChopStick struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopStick
}

func (p Philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	// initialize chopsticks
	ChopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}
	// initialize philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{ChopSticks[i], ChopSticks[(i+1)%5]}
	}
	// start the dining
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	// this probably would not work, cause main routine would be killed faster...
	// so we let it run until user cancels that main routine
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	sig := <-shutdown
	fmt.Printf("Shutting down on signal %s..\n", sig)

}
