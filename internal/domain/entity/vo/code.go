package vo

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrCodeNotRetrieved = errors.New("cep code was not retrieved")
	ErrCodeIsNotValid   = errors.New("cep code is not valid")
)

type Code struct {
	ID string
}

func NewCode(cep string) (*Code, error) {
	newCode := &Code{ID: cep}
	err := newCode.validate()
	if err != nil {
		return nil, err
	}
	return newCode, nil
}

func (c *Code) validate() error {
	if c.ID == "" {
		return ErrCodeNotRetrieved
	}
	c.formatCode()

	if !isCepCodeValid(c.ID) {
		return ErrCodeIsNotValid
	}
	return nil
}

func (c *Code) formatCode() {
	if strings.ContainsAny(c.ID, "-") {
		c.ID = fmt.Sprintf("%08s", c.ID)
		return
	}
	c.ID = c.ID[:5] + "-" + c.ID[5:]
}

func isCepCodeValid(code string) bool {
	re := regexp.MustCompile("^\\d{5}-\\d{3}$")
	return re.MatchString(code)
}
