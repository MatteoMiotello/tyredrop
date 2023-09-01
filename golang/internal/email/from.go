package email

import "fmt"

type From struct {
	Email string
	Name  string
}

func NewNoreplyFrom() From {
	return From{
		Email: "noreply@tyresintheworld.it",
		Name:  "NoReply",
	}
}

func (r From) String() string {
	return fmt.Sprintf("%s <%s>", r.Name, r.Email)
}
