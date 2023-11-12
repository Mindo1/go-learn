package libs

import "testing"

func TestLogin(t *testing.T) {
	Greeting()
	if Login("admin", "234") != true {
		t.Errorf("error test fail")
	}
	if Login("", "234") != false {
		t.Errorf("error test fail")
	}
}
