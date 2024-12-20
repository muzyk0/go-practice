package fullname

import (
	"fmt"
	"strings"
)

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", u.FirstName, u.LastName))
}
