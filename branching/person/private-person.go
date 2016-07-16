package person

import "fmt"

type PrivatePerson struct {
	PrivateName string
}

func init() {
	fmt.Println("initializing person package...")
}
