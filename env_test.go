package environ_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/MarcSky/environ"
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

func TestMustGetInt(t *testing.T) {
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

func TestInt64GetEnv(t *testing.T) {
	name := "ENV_TEST_GET_INT64"
	var expectedValue int64 = 100

	err := os.Setenv(name, fmt.Sprintf("%d", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.Int64GetEnv(name, 1)
	if value != expectedValue {
		t.Errorf("environ.Int64GetEnv test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}

	value = environ.Int64GetEnv("ENV_TEST_GET_INT64_OTHER", expectedValue)
	if value != expectedValue {
		t.Errorf("environ.Int64GetEnv test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}
}

func TestMustGetInt64(t *testing.T) {
	name := "ENV_TEST_GET_INT64"
	var expectedValue int64 = 100

	err := os.Setenv(name, fmt.Sprintf("%d", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.MustGetInt64(name)
	if value != expectedValue {
		t.Errorf("environ.MustGetInt64 test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("environ.MustGetInt64 test failed. Did not panic for missing value.")
		}
	}()

	environ.MustGetInt64("ENV_TEST_GET_INT64_OTHER")
}

func TestUInt64GetEnv(t *testing.T) {
	name := "ENV_TEST_GET_UINT64"
	var expectedValue uint64 = 100

	err := os.Setenv(name, fmt.Sprintf("%d", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.UInt64GetEnv(name, 1)
	if value != expectedValue {
		t.Errorf("environ.UInt64GetEnv test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}

	value = environ.UInt64GetEnv("ENV_TEST_GET_UINT64_OTHER", expectedValue)
	if value != expectedValue {
		t.Errorf("environ.UInt64GetEnv test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}
}

func TestMustGetUInt64(t *testing.T) {
	name := "ENV_TEST_GET_UINT64"
	var expectedValue uint64 = 100

	err := os.Setenv(name, fmt.Sprintf("%d", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.MustGetUInt64(name)
	if value != expectedValue {
		t.Errorf("environ.MustGetUInt64 test failed. Expected: [%d] Got: [%d]", expectedValue, value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("environ.MustGetUInt64 test failed. Did not panic for missing value.")
		}
	}()

	environ.MustGetUInt64("ENV_TEST_GET_UINT64_OTHER")
}

func TestFloatGetEnv(t *testing.T) {
	name := "ENV_TEST_GET_FLOAT64"
	var expectedValue = 100.0

	err := os.Setenv(name, fmt.Sprintf("%1.f", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.FloatGetEnv(name, 1)
	if value != expectedValue {
		t.Errorf("environ.FloatGetEnv test failed. Expected: [%f] Got: [%f]", expectedValue, value)
	}

	value = environ.FloatGetEnv("ENV_TEST_GET_FLOAT64_OTHER", expectedValue)
	if value != expectedValue {
		t.Errorf("environ.FloatGetEnv test failed. Expected: [%f] Got: [%f]", expectedValue, value)
	}
}

func TestMustGetFloat64(t *testing.T) {
	name := "ENV_TEST_GET_FLOAT64"
	var expectedValue = 100.0

	err := os.Setenv(name, fmt.Sprintf("%1.f", expectedValue))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.MustGetFloat(name)
	if value != expectedValue {
		t.Errorf("environ.MustGetFloat test failed. Expected: [%f] Got: [%f]", expectedValue, value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("environ.MustGetFloat test failed. Did not panic for missing value.")
		}
	}()

	environ.MustGetFloat("ENV_TEST_GET_FLOAT64_OTHER")
}

func TestBoolGetEnv(t *testing.T) {
	name := "ENV_TEST_GET_BOOL"

	err := os.Setenv(name, strconv.FormatBool(false))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.BoolGetEnv(name, false)
	if value != false {
		t.Errorf("environ.BoolGetEnv test failed. Expected: [%s] Got: [%s]", strconv.FormatBool(false), strconv.FormatBool(value))
	}

	value = environ.BoolGetEnv("ENV_TEST_GET_BOOL_OTHER", true)
	if value != true {
		t.Errorf("environ.BoolGetEnv test failed. Expected: [%s] Got: [%s]", strconv.FormatBool(false), strconv.FormatBool(value))
	}
}

func TestMustGetBool(t *testing.T) {
	name := "ENV_TEST_GET_BOOL"

	err := os.Setenv(name, strconv.FormatBool(true))
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.MustGetBool(name)
	if value != true {
		t.Errorf("environ.MustGetBool test failed. Expected: [%s] Got: [%s]", strconv.FormatBool(true), strconv.FormatBool(value))
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("environ.MustGetBool test failed. Did not panic for missing value.")
		}
	}()

	environ.MustGetBool("ENV_TEST_GET_BOOL_OTHER")
}

func TestStringGetEnv(t *testing.T) {
	name := "ENV_TEST_GET_STRING"
	expectedValue := "hello"

	err := os.Setenv(name, expectedValue)
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.StrGetEnv(name, "world")
	if value != expectedValue {
		t.Errorf("environ.StrGetEnv test failed. Expected: [%s] Got: [%s]", expectedValue, value)
	}

	value = environ.StrGetEnv("ENV_TEST_GET_STRING_OTHER", "world")
	if value != "world" {
		t.Errorf("environ.StrGetEnv test failed. Expected: [%s] Got: [%s]", "world", value)
	}
}

func TestMustGetString(t *testing.T) {
	name := "ENV_TEST_GET_STRING"
	expectedValue := "hello"

	err := os.Setenv(name, expectedValue)
	if err != nil {
		t.Error("os.Setenv returned unexpected error:", err)
	}

	value := environ.MustGetString(name)
	if value != expectedValue {
		t.Errorf("environ.MustGetString test failed. Expected: [%s] Got: [%s]", expectedValue, value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("environ.MustGetString test failed. Did not panic for missing value.")
		}
	}()

	environ.MustGetString("ENV_TEST_GET_STRING_OTHER")
}
