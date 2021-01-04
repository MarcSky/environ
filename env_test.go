package environ_test

import (
	"environ"
	"fmt"
	"os"
	"testing"
)

func TestIntGetEnv(t *testing.T) {
	name := "ENV_TEST_GET_INT"
	expectedValue := 100

	err := os.Setenv(name, fmt.Sprintf("%d", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.IntGetEnv(name, 1)
	if value != expectedValue {
		t.Errorf("environ.IntGetEnv test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}

	value = environ.IntGetEnv("ENV_TEST_GET_INT_OTHER", expectedValue)
	if value != expectedValue {
		t.Errorf("environ.IntGetEnv test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}
}

func TestMustGet(t *testing.T) {
	name := "ENV_TEST_GET_INT"
	expectedValue := 100

	err := os.Setenv(name, fmt.Sprintf("%d", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.MustGetInt(name)
	if value != expectedValue {
		t.Errorf("environ.MustGetInt test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("environ.MustGetInt test failed. Did not panic for missing value.")
		}
	}()

	environ.MustGetInt("ENV_TEST_GET_INT_OTHER")
}
