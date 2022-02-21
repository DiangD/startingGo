package paxos

type MsgType = uint8

const (
	Prepare = iota
	Promise
	Propose
	Accept
)

type message struct {
	tp     MsgType
	from   int
	to     int
	number int    //propose number
	value  string //propose value
}
