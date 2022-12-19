package entity

import (
	"cep-retriever/internal/domain/entity/vo"
	"errors"
)

type Cep struct {
	Code     *vo.Code
	Address  string
	District string
	City     string
	State    string
}

var (
	ErrAddressNotRetrieved  = errors.New("address was not retrieved")
	ErrDistrictNotRetrieved = errors.New("district was not retrieved")
	ErrCityNotRetrieved     = errors.New("city was not retrieved")
	ErrStateNotRetrieved    = errors.New("state was not retrieved")
)

func NewCep(cep string, address string, district string, city string, state string) (*Cep, error) {
	newCode, err := vo.NewCode(cep)
	if err != nil {
		return nil, err
	}

	newCep := &Cep{
		Code:     newCode,
		Address:  address,
		District: district,
		City:     city,
		State:    state,
	}

	err = newCep.Validate()
	if err != nil {
		return nil, err
	}
	return newCep, nil
}

func (c *Cep) Validate() error {

	if c.Address == "" {
		return ErrAddressNotRetrieved
	}

	if c.District == "" {
		return ErrDistrictNotRetrieved
	}

	if c.City == "" {
		return ErrCityNotRetrieved
	}

	if c.State == "" {
		return ErrStateNotRetrieved
	}
	return nil
}
