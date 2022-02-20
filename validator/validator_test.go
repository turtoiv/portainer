package validator

import "testing"

func TestValidatorWithSuccess(t *testing.T) {

	res := Validate("{{}()[{}]}")
	if res != true {
		t.Logf("Wanted true but got %t", res)
		t.Fail()
	}

	res = Validate("{()[]{}}")
	if res != true {
		t.Logf("Wanted true but got %t", res)
		t.Fail()
	}
}

func TestValidatorWithError(t *testing.T) {

	res := Validate("{]}()")
	if res != false {
		t.Logf("Wanted false but got %t", res)
		t.Fail()
	}
}
