package gde

import (
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	if v := os.Getenv("MYENV"); v != "" {
		t.Fatalf("MYENV is set: %s\n", v)
	}
	SetEnv("./dev_env")
	if v := os.Getenv("MYENV"); v != "123" {
		t.Fatalf("MYENV is not set correctly: %s\n", v)
	}
}
