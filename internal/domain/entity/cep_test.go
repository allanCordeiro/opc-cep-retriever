package entity

import (
	"testing"
)

type data struct {
	Cep      string
	Address  string
	District string
	City     string
	State    string
}
type testTable struct {
	scenarioName   string
	scenarioErrMsg error
	scenarioData   data
}

func TestGivenAMissingFieldWhenValidateShouldTriggerAnError(t *testing.T) {
	tt := []testTable{
		{
			scenarioName:   "Without address",
			scenarioErrMsg: ErrAddressNotRetrieved,
			scenarioData: data{
				Cep:      "01234-567",
				Address:  "",
				District: "Some district",
				City:     "Some city",
				State:    "SS",
			},
		},
		{
			scenarioName:   "Without district",
			scenarioErrMsg: ErrDistrictNotRetrieved,
			scenarioData: data{
				Cep:      "01234-567",
				Address:  "Some address",
				District: "",
				City:     "Some city",
				State:    "SS",
			},
		},
		{
			scenarioName:   "Without city",
			scenarioErrMsg: ErrCityNotRetrieved,
			scenarioData: data{
				Cep:      "01234-567",
				Address:  "Some address",
				District: "Some district",
				City:     "",
				State:    "SS",
			},
		},
		{
			scenarioName:   "Without state",
			scenarioErrMsg: ErrStateNotRetrieved,
			scenarioData: data{
				Cep:      "01234-567",
				Address:  "Some address",
				District: "Some district",
				City:     "Some city",
				State:    "",
			},
		},
	}
	for _, test := range tt {
		t.Run(test.scenarioName, func(t *testing.T) {
			_, err := NewCep(
				test.scenarioData.Cep,
				test.scenarioData.Address,
				test.scenarioData.District,
				test.scenarioData.City,
				test.scenarioData.State,
			)
			if err == nil {
				t.Errorf("expecting an error but no error was ocurred")
			}
			if err != test.scenarioErrMsg {
				t.Errorf("expected '%s' but '%s', was received", test.scenarioErrMsg, err.Error())
			}

		})
	}
}

func TestGivenAllFieldsWhenValidateShouldBeOk(t *testing.T) {
	expectedCep := data{
		Cep:      "01234-567",
		Address:  "Some address",
		District: "Some district",
		City:     "Some city",
		State:    "SS",
	}

	output, err := NewCep(
		expectedCep.Cep,
		expectedCep.Address,
		expectedCep.District,
		expectedCep.City,
		expectedCep.State,
	)
	if err != nil {
		t.Errorf("expected to be sucess but an error was found")
	}

	if expectedCep.Cep != output.Code.ID {
		t.Errorf("Expected '%s', but received '%s'", expectedCep.Cep, output.Code)
	}

	if expectedCep.Address != output.Address {
		t.Errorf("Expected '%s', but received '%s'", expectedCep.Address, output.Address)
	}

	if expectedCep.District != output.District {
		t.Errorf("Expected '%s', but received '%s'", expectedCep.District, output.District)
	}

	if expectedCep.City != output.City {
		t.Errorf("Expected '%s', but received '%s'", expectedCep.City, output.City)
	}

	if expectedCep.State != output.State {
		t.Errorf("Expected '%s', but received '%s'", expectedCep.State, output.State)
	}
}

func TestGivenAnUnmaskedCepCodeWhenValidateShouldBeOk(t *testing.T) {
	inputedCep := data{
		Cep:      "01234567",
		Address:  "Some address",
		District: "Some district",
		City:     "Some city",
		State:    "SS",
	}
	expectedCepCode := "01234-567"

	output, err := NewCep(
		inputedCep.Cep,
		inputedCep.Address,
		inputedCep.District,
		inputedCep.City,
		inputedCep.State,
	)
	if err != nil {
		t.Errorf("expected to be sucess but an error was found: %s", err)
	}

	if expectedCepCode != output.Code.ID {
		t.Errorf("Expected '%s', but received '%s'", expectedCepCode, output.Code)
	}

}
