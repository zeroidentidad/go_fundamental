package main

import "testing"

//go test -run TestReverse -v
func TestReverse(t *testing.T) {
	tests := map[string]struct {
		given    string
		expected string
	}{
		"test_par": {
			given:    "uno dos tres cuatro",
			expected: "cuatro tres dos uno",
		},
		"test impar": {
			given:    "uno dos tres",
			expected: "tres dos uno",
		},
		"test_one_word": {
			given:    "tom",
			expected: "tom",
		},
	}

	for testName, testData := range tests {
		t.Run(testName, func(t *testing.T) {
			reversed := reverseStringArrayFields(testData.given)
			/*reversed := reverseStringFieldsBuilder(testData.given)
			reversed := mirrorStringBytes(testData.given)
			reversed := mirrorStringRunes(testData.given)
			reversed := mirrorStringLikeRunes(testData.given)*/

			if reversed != testData.expected {
				t.Logf("got:%v expected:%v", reversed, testData.expected)
				t.Fail()
			} else {
				t.Logf("got:%v expected:%v", reversed, testData.expected)
			}
		})
	}
}
