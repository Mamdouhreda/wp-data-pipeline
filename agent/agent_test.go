package agent

import (
	"testing"

	"github.com/joho/godotenv"
)


func TestAgent_Success(t *testing.T) {
	_ = godotenv.Load(".env")
	agent("how are you")
}

