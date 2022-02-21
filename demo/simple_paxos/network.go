package paxos

import (
	"log"
	"time"
)

type network interface {
	send(m message)
	receive(d time.Duration) (message, bool)
}

func (n *nodeNetwork) receive(d time.Duration) (message, bool) {
	return n.receiveFrom(n.id, d)
}

type basicNetwork struct {
	queue map[int]chan message
}

type nodeNetwork struct {
	id int
	*basicNetwork
}

func newBasicNetwork(ids ...int) *basicNetwork {
	bnw := &basicNetwork{
		queue: make(map[int]chan message, 0),
	}
	for _, id := range ids {
		bnw.queue[id] = make(chan message, 1024)
	}
	return bnw
}

func (b *basicNetwork) nodeNetwork(id int) *nodeNetwork {
	return &nodeNetwork{id: id, basicNetwork: b}
}

func (b *basicNetwork) send(m message) {
	log.Printf("net: send %+v", m)
	b.queue[m.to] <- m
}

func (b *basicNetwork) receiveFrom(from int, d time.Duration) (message, bool) {
	select {
	case m := <-b.queue[from]:
		log.Printf("net: recv %+v", m)
		return m, true
	case <-time.After(d):
		return message{}, false
	}
}
