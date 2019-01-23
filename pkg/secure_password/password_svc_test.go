package secure_password

import (
	"testing"
)

func TestPasswordService_Create(t *testing.T) {
	testCases := []struct {
		name        string
		lettersLen  int
		specChLen   int
		numbersLen  int
		expectedLen int
	}{
		{name: "only letters", lettersLen: 5, specChLen: 0, numbersLen: 0, expectedLen: 5},
		{name: "only special characters", lettersLen: 0, specChLen: 7, numbersLen: 0, expectedLen: 7},
		{name: "only numbers", lettersLen: 0, specChLen: 0, numbersLen: 9, expectedLen: 9},
		{name: "letters + special characters + numbers", lettersLen: 3, specChLen: 7, numbersLen: 4,
			expectedLen: 14},
	}

	passSvc := PasswordService{}

	for _, tc := range testCases {
		t.Log("Test case name: ", tc.name)

		pass := passSvc.Create(tc.lettersLen, tc.specChLen, tc.numbersLen)
		if tc.expectedLen != len([]rune(pass)) {
			t.Errorf("expected %d and got %d results are not equal", tc.expectedLen,
				len([]rune(pass)))
		}
	}
}
