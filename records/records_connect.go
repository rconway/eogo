package records

import "fmt"

type RecordsConnect struct {
	Server string
}

func (rc *RecordsConnect) Execute() error {
	fmt.Printf("running records connect with server=%v\n", rc.Server)
	return nil
}
