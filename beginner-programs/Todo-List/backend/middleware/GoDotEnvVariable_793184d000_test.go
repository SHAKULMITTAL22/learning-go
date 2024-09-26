package middleware

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGoDotEnvVariable_793184d000(t *testing.T) {
	t.Run("env variable loaded correctly", func(t *testing.T) {
		variable := goDotEnvVariable("TEST_VARIABLE")
		assert.Equal(t, "test", variable, fmt.Sprintf("Failed to load environment variable. Expected 'test', got '%s'", variable))
	})

	t.Run("non-existing env variable", func(t *testing.T) {
		variable := goDotEnvVariable("NON_EXISTING_VARIABLE")
		assert.Equal(t, "", variable, fmt.Sprintf("Function should return empty string for non-existing variable. Got '%s'", variable))
	})
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
