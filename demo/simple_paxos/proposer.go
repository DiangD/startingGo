package paxos

import (
	"log"
	"time"
)

type proposer struct {
	id        int
	round     int
	number    int
	value     string
	acceptors map[int]bool
	net       network
}

func newProposer(id int, value string, net network, acceptors ...int) *proposer {
	p := &proposer{id: id, round: 0, value: value, net: net, acceptors: make(map[int]bool, len(acceptors))}
	for _, acceptor := range acceptors {
		p.acceptors[acceptor] = false
	}
	return p
}

func (p *proposer) run() {
	var (
		msg message
		ok  bool
	)
	for !p.majorityReached() {
		if !ok {
			prepareMsg := p.prepare()
			for _, msg := range prepareMsg {
				p.net.send(msg)
			}
		}
		msg, ok = p.net.receive(time.Hour)
		if !ok {
			continue
		}
		switch msg.tp {
		case Promise:
			p.handlePromise(msg)
		default:
			panic("UnSupport message.")
		}
	}
	proposeMsg := p.accept()
	for _, msg := range proposeMsg {
		p.net.send(msg)
	}
}

func (p *proposer) prepare() []message {
	p.round++
	p.number = p.proposalNumber()
	messages := make([]message, p.majority())
	i := 0
	for to := range p.acceptors {
		messages[i] = message{
			tp:     Prepare,
			from:   p.id,
			to:     to,
			number: p.number,
		}
		i++
		if i == p.majority() {
			break
		}
	}
	return messages
}
func (p *proposer) handlePromise(reply message) {
	log.Printf("proposer: %d received a promise %+v", p.id, reply)
	p.acceptors[reply.from] = true
	if reply.number > 0 {
		p.number = reply.number
		p.value = reply.value
	}
}

func (p *proposer) accept() []message {
	msg := make([]message, p.majority())
	i := 0
	for to, ok := range p.acceptors {
		if ok {
			msg[i] = message{
				tp:     Propose,
				from:   p.id,
				to:     to,
				number: p.number,
				value:  p.value,
			}
			i++
		}
		if i == p.majority() {
			break
		}
	}
	return msg
}

func (p *proposer) proposalNumber() int {
	return p.round<<16 | p.id
}

func (p *proposer) majority() int {
	return len(p.acceptors)/2 + 1
}

func (p *proposer) majorityReached() bool {
	count := 0
	for _, ok := range p.acceptors {
		if ok {
			count++
		}
	}
	if count >= p.majority() {
		return true
	}
	return false
}
