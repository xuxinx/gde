package gde

import (
	"os"
	"testing"
)

func TestSetenv(t *testing.T) {
	if v := os.Getenv("MYENV"); v != "" {
		t.Fatalf("MYENV is set: %s\n", v)
	}
	Setenv("./dev_env")
	if v := os.Getenv("MYENV"); v != "123" {
		t.Fatalf("MYENV is not set correctly: %s\n", v)
	}
}
