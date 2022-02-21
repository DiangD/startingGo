package paxos

import (
	"log"
	"time"
)

type learner struct {
	id        int
	acceptors map[int]message
	net       network
}

func newLearner(id int, net network, acceptors ...int) *learner {
	l := &learner{id: id, acceptors: make(map[int]message), net: net}
	for _, id := range acceptors {
		l.acceptors[id] = message{tp: Accept}
	}
	return l
}

func (l *learner) learn() string {
	for {
		msg, ok := l.net.receive(time.Hour)
		if !ok {
			continue
		}
		if msg.tp != Accept {
			log.Panicf("learner: %d receivd unexpected message %+v", l.id, msg)
		}
		l.handleAccepted(msg)
		accept, ok := l.chosen()
		if !ok {
			continue
		}
		return accept.value
	}
}

func (l *learner) handleAccepted(args message) {
	a := l.acceptors[args.from]
	if a.number < args.number {
		log.Printf("learner: %d received a new accepted proposal %+v", l.id, args)
		l.acceptors[args.from] = args
	}
}

func (l *learner) majority() int {
	return len(l.acceptors)/2 + 1
}

func (l *learner) chosen() (message, bool) {
	acceptedCount := make(map[int]int)
	acceptedMsg := make(map[int]message)

	for _, accepted := range l.acceptors {
		if accepted.number != 0 {
			acceptedCount[accepted.number]++
			acceptedMsg[accepted.number] = accepted
		}
	}

	for k, count := range acceptedCount {
		if count >= l.majority() {
			return acceptedMsg[k], true
		}
	}
	return message{}, false
}
