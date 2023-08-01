package email

import "fmt"

type From struct {
	Email string
	Name  string
}

func (r From) String() string {
	return fmt.Sprintln("%s <%s>", r.Name, r.Email)
}
