package utils

import (
	"testing"
	"time"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Initialize JWT with a secret
	InitJWT("test-secret-key")

	userID := uint(1)
	roleID := uint(2)

	// Generate token
	token, err := GenerateToken(userID, roleID)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Error("Generated token should not be empty")
	}

	// Validate token
	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID %d, got %d", userID, claims.UserID)
	}

	if claims.RoleID != roleID {
		t.Errorf("Expected RoleID %d, got %d", roleID, claims.RoleID)
	}

	// Check expiration time
	if claims.ExpiresAt == nil {
		t.Error("Token should have an expiration time")
	} else {
		expiresAt := claims.ExpiresAt.Time
		expectedExpiration := time.Now().Add(24 * time.Hour)
		diff := expiresAt.Sub(expectedExpiration)
		if diff < -time.Minute || diff > time.Minute {
			t.Errorf("Token expiration time is not as expected")
		}
	}
}

func TestValidateInvalidToken(t *testing.T) {
	InitJWT("test-secret-key")

	// Test with invalid token
	_, err := ValidateToken("invalid.token.here")
	if err == nil {
		t.Error("ValidateToken should fail with invalid token")
	}

	// Test with empty token
	_, err = ValidateToken("")
	if err == nil {
		t.Error("ValidateToken should fail with empty token")
	}
}

func TestValidateTokenWithDifferentSecret(t *testing.T) {
	// Generate token with one secret
	InitJWT("secret-1")
	token, err := GenerateToken(1, 1)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Try to validate with a different secret
	InitJWT("secret-2")
	_, err = ValidateToken(token)
	if err == nil {
		t.Error("ValidateToken should fail when secret is different")
	}
}
