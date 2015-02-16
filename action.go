package main

// Action represents a single in-game action, such as "go north".
type Action struct {
	Command string
	Func    ActionFunc
}
