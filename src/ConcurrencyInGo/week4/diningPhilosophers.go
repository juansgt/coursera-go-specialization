package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS      *ChopS
	rightCS     *ChopS
	philoNumber int
	host        *Host
}

type Host struct {
	channel chan struct{}
}

func (host *Host) askPermissionToEat() {
	host.channel <- struct{}{}
}

func (host *Host) stopEating() {
	<-host.channel
}

func (p *Philo) pickUpChopSticks() {
	switch rand.Intn(2) {
	case 0:
		p.leftCS.Lock()
		p.rightCS.Lock()
	case 1:
		p.rightCS.Lock()
		p.leftCS.Lock()
	}
}

func (p *Philo) dropChopSticks() {
	p.rightCS.Unlock()
	p.leftCS.Unlock()
}

func (p *Philo) eat() {
	for i := 0; i < 3; i++ {
		p.pickUpChopSticks()
		p.host.askPermissionToEat()
		fmt.Println("starting to eat ", p.philoNumber)

		p.host.stopEating()
		p.dropChopSticks()
		fmt.Println("finishing eating ", p.philoNumber)
	}
}

func philoEat(philo *Philo, waitGroup *sync.WaitGroup) {
	philo.eat()

	waitGroup.Done()
}

func main() {
	var cSticks []*ChopS = make([]*ChopS, 5)
	var philos []*Philo = make([]*Philo, 5)
	var host *Host = new(Host)
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	host.channel = make(chan struct{}, 2)

	for i := 0; i < 5; i++ {
		cSticks[i] = new(ChopS)
	}

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			cSticks[i],
			cSticks[(i+1)%5],
			i + 1,
			host,
		}
	}

	for i := 0; i < 5; i++ {
		go philoEat(philos[i], &waitGroup)
	}

	waitGroup.Wait()
}
