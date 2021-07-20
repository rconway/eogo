package process

import "fmt"

type ProcessConnect struct {
	Server string
}

func (pc *ProcessConnect) Execute() error {
	fmt.Printf("running process connect with server=%v\n", pc.Server)
	return nil
}
