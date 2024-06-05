package controllers
import (
	"testing"
)
func TestValidIdParm(t *testing.T) {
	_, valid := ValidIdParm("")
	if valid {
		t.Error(`(ValidIdParm("") == true`)
	}

	_, valid = ValidIdParm("1A")
	if valid {
		t.Error(`(ValidIdParm("1A") == true`)
	}


	_, valid = ValidIdParm("1")
	if !valid {
		t.Error(`(ValidIdParm("") == false`)
	}
}

func TestValidPassword(t *testing.T) {
	// valid := ValidPassword("")
	if ValidPassword("") {
		t.Error(`(ValidPassword("") == true`)
	}

	if ValidPassword("1") {
		t.Error(`(ValidPassword(1") == true`)
	}

	if ValidPassword("12") {
		t.Error(`(ValidPassword(12") == true`)
	}

	if  ValidPassword("123") {
		t.Error(`(ValidPassword("123") == true`)
	}

	if ValidPassword("1234"){
		t.Error(`(ValidPassword("1234") == true`)
	}
	if ValidPassword("12345") {
		t.Error(`(ValidPassword("12345") == true`)
	}

	
	if !ValidPassword("123456") {
		t.Error(`(ValidPassword("123456") == false`)
	}
	if !ValidPassword("1234567"){
		t.Error(`(ValidPassword("1234567") == false`)
	}

}
func TestValidEmail(t *testing.T) {
	if ValidEmail("") {
		t.Error(`(ValidEmail("") == false`)
	}
	if ValidEmail(" ") {
		t.Error(`(ValidEmail(" ") == false`)
	}
	if ValidEmail("1") {
		t.Error(`(ValidEmail("1") == false`)
	}
	if ValidEmail("a") {
		t.Error(`(ValidEmail("a") == false`)
	}
	if ValidEmail("a.b") {
		t.Error(`(ValidEmail("a.b") == false`)
	}
	if ValidEmail("a.b@") {
		t.Error(`(ValidEmail("a.b@") == false`)
	}
	if ValidEmail("a.b@c") {
		t.Error(`(ValidEmail("a.b@c") == false`)
	}
	if ValidEmail("a.b@c.") {
		t.Error(`(ValidEmail("a.b@c.") == false`)
	}
	if ValidEmail("a.b@c.d") {
		t.Error(`(ValidEmail("a.b@c.d") == false`)
	}
	if !ValidEmail("a.b@c.com") {
		t.Error(`(ValidEmail("a.b@c.com") == false`)
	}

}