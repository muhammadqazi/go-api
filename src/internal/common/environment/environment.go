package environment

import "fmt"

const (
	Development = Name("development")
	Staging     = Name("staging")
	Production  = Name("production")
	Test        = Name("test")
)

type Name string

func (e *Name) EnvDecode(val string) error {
	env := Name(val)
	if !ValidEnvironment(env) {
		return fmt.Errorf("invalid environment name '%s'", val)
	}
	*e = env
	return nil
}

func (e Name) Development() bool {
	return e == Development
}

func (e Name) Staging() bool {
	return e == Staging
}

func (e Name) Production() bool {
	return e == Production
}

func ValidEnvironment(e Name) bool {
	for _, env := range []Name{Development, Staging, Production, Test} {
		if e == env {
			return true
		}
	}
	return false
}
