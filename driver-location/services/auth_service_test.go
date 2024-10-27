package services_test

import (
	"location-service/utils"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := utils.GenerateToken("testUser")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Errorf("Expected a token, got an empty string")
	}
}

func TestValidateToken(t *testing.T) {
	token, _ := utils.GenerateToken("driver")
	_, err := utils.ValidateToken(token)
	if err != nil {
		t.Errorf("Expected token to be valid")
	}
}