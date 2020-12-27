package util

import "fmt"

type ErrGroup []error

func (e ErrGroup) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}
	return fmt.Sprintf("found %d errors, first one: %s", len(e), e[0])
}
