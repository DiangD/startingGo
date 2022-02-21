package paxos

import (
	"log"
	"time"
)

type acceptor struct {
	id             int
	promiseNumber  int
	acceptedNumber int
	acceptedValue  string

	learner []int
	net     network
}

func newAcceptor(id int, net network, learners ...int) *acceptor {
	return &acceptor{
		id:      id,
		net:     net,
		learner: learners,
	}
}

func (a *acceptor) run() {
	log.Printf("acceptor %d run", a.id)
	for {
		msg, ok := a.net.receive(time.Hour)
		if !ok {
			continue
		}
		switch msg.tp {
		case Prepare:
			promise, ok := a.handlePrepare(msg)
			if ok {
				a.net.send(promise)
			}
		case Propose:
			accepted := a.handleAccept(msg)
			if accepted {
				for _, l := range a.learner {
					msg := message{
						tp:     Accept,
						from:   a.id,
						to:     l,
						number: a.acceptedNumber,
						value:  a.acceptedValue,
					}
					a.net.send(msg)
				}
			}
		default:
			log.Panicf("acceptor: %d unexpected message type: %v", a.id, msg.tp)
		}
	}
}

func (a *acceptor) handlePrepare(args message) (message, bool) {
	if a.promiseNumber >= args.number {
		return message{}, false
	}
	a.promiseNumber = args.number
	m := message{
		tp:     Promise,
		from:   a.id,
		to:     args.from,
		number: a.acceptedNumber,
		value:  a.acceptedValue,
	}
	return m, true
}

func (a *acceptor) handleAccept(args message) bool {
	if args.number >= a.acceptedNumber {
		a.acceptedNumber = args.number
		a.acceptedValue = args.value
		a.promiseNumber = args.number
		return true
	}
	return false

}
