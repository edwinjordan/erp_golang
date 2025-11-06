package models

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	user := &User{}
	password := "testpassword123"

	err := user.HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if user.Password == password {
		t.Error("Password should be hashed, not plain text")
	}

	if len(user.Password) == 0 {
		t.Error("Hashed password should not be empty")
	}
}

func TestCheckPassword(t *testing.T) {
	user := &User{}
	password := "testpassword123"

	err := user.HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Test with correct password
	err = user.CheckPassword(password)
	if err != nil {
		t.Error("CheckPassword should succeed with correct password")
	}

	// Test with incorrect password
	err = user.CheckPassword("wrongpassword")
	if err == nil {
		t.Error("CheckPassword should fail with incorrect password")
	}
	if err != bcrypt.ErrMismatchedHashAndPassword {
		t.Errorf("Expected bcrypt.ErrMismatchedHashAndPassword, got %v", err)
	}
}
