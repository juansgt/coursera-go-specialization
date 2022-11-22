package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
	host    *Host
}

type Host struct {
	mu    sync.Mutex
	count int
}

func (host *Host) askPermissionToEat() {
	host.mu.Lock()
	if host.count < 2 {
		host.count++
		if host.count < 2 {
			host.mu.Unlock()
		}
	}
}

func (host *Host) stopEating() {
	host.mu.Lock()
	if host.count > 0 {
		host.count--
	}
	host.mu.Unlock()
}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		p.host.askPermissionToEat()
		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		p.host.stopEating()
	}
}

func philoEat(philos []*Philo, i int, waitGroup *sync.WaitGroup) {
	philos[i].eat()
}

func main() {
	var cSticks []*ChopS = make([]*ChopS, 5)
	var philos []*Philo = make([]*Philo, 5)
	var host *Host = new(Host)
	var waitGroup sync.WaitGroup

	for i := 0; i < 5; i++ {
		cSticks[i] = new(ChopS)
	}

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			cSticks[i],
			cSticks[(i+1)%5],
			host,
		}
	}

	for i := 0; i < 5; i++ {
		go philoEat(philos, i, &waitGroup)
	}

	waitGroup.Wait()
}
