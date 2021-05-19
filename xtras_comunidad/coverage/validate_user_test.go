package validate

import "testing"

var usernameTests = []struct {
	in        string
	wantValid bool
}{
	{"", false},
	{"gopher", true},
	{"gopher$", false},
	{"abcdefghijklmnopqrstuvwxyzabcde", false},
}

func TestUsername(t *testing.T) {
	for _, tt := range usernameTests {
		valid, err := Username(tt.in)
		if err != nil && tt.wantValid {
			t.Fatal(err)
		}

		if valid != tt.wantValid {
			t.Errorf("Username(%q) = %t, want %t", tt.in, valid, tt.wantValid)
		}

	}

}
