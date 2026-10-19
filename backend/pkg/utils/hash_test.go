package utils

import "testing"

func TestHash_Valid(t *testing.T) {
	t.Helper()

	testPassword := "TestString!23"

	hashedBytes, err := HashPassword(testPassword, 12)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	isHashValid := ComparePassword(hashedBytes, testPassword)
	if !isHashValid {
		t.Errorf("isHashValid = false, want true")
	}
}
