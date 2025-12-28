package agents

type Agent interface {
	Percept()
	Deliberate()
	Act()
}