package models

import (
	"fmt"
	"regexp"

	"github.com/revel/revel"
)

var userRegex = regexp.MustCompile("^\\w*s")

// User strut
type User struct {
	UserID                   int
	Name, Username, Password string
	HashedPassword           []byte
}

func (u *User) String() string {
	return fmt.Sprintf("User (%s)", u.Username)
}

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{3},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
