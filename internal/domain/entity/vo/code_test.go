package vo

import "testing"

type data struct {
	Cep string
}

func TestGivenAnInvalidCepCodeWhenValidateShouldTriggerAnError(t *testing.T) {
	expectedCep := data{
		Cep: "234-567",
	}
	expectedError := ErrCodeIsNotValid

	_, err := NewCode(expectedCep.Cep)
	if err == nil {
		t.Errorf("expecting an error but no error was ocurred")
	}

	if expectedError != err {
		t.Errorf("Expected '%s', but received '%s'", expectedError.Error(), err.Error())
	}

}

func TestGivenAnEmptyCepCodeWhenValidateShouldTriggerAnError(t *testing.T) {
	expectedCep := data{
		Cep: "",
	}
	expectedError := ErrCodeNotRetrieved

	_, err := NewCode(expectedCep.Cep)
	if err == nil {
		t.Errorf("expecting an error but no error was ocurred")
	}

	if expectedError != err {
		t.Errorf("Expected '%s', but received '%s'", expectedError.Error(), err.Error())
	}
}
