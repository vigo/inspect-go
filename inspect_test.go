package inspect

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// TestMain runs tests
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestInteger(t *testing.T) {
	el := 5
	result := Element(&el)

	if !strings.Contains(result, "*int") {
		t.Errorf("result must contain [*int]")
	}
}

func TestFloat(t *testing.T) {
	el := 5.1
	result := Element(&el)

	if !strings.Contains(result, "*float64") {
		t.Errorf("result must contain [*float64]")
	}
}

func TestString(t *testing.T) {
	el := "hello"
	result := Element(&el)

	if !strings.Contains(result, "*string") {
		t.Errorf("result must contain [*string]")
	}
}

func TestArray(t *testing.T) {
	el := []string{"hello", "world"}
	result := Element(&el)

	if !strings.Contains(result, "*[]string") {
		t.Errorf("result must contain [*[]string]")
	}
}

func TestMap(t *testing.T) {
	el := map[string]string{
		"foo": "bar",
	}
	result := Element(&el)

	if !strings.Contains(result, "*map[string]string") {
		t.Errorf("result must contain [*map[string]string]")
	}

	if !strings.Contains(result, "map[foo:bar]") {
		t.Errorf("result must contain [map[foo:bar]]")
	}
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func TestStruct(t *testing.T) {
	el := &User{"Uğur", "Özyılmazel", "vigo@example.com"}
	result := Element(el)

	if !strings.Contains(result, "*inspect.User") {
		t.Errorf("result must contain [*inspect.User]")
	}
	if !strings.Contains(result, "# of Fields: 3") {
		t.Errorf("result must contain [# of Fields: 2]")
	}
	if !strings.Contains(result, "json:\"first_name\"") {
		t.Errorf("result must contain [json:\"first_name\"]")
	}
	if !strings.Contains(result, "json:\"last_name\"") {
		t.Errorf("result must contain [json:\"last_name\"]")
	}
	if !strings.Contains(result, "json:\"email\"") {
		t.Errorf("result must contain [json:\"email\"]")
	}
	if !strings.Contains(result, "# of Methods: 1") {
		t.Errorf("result must contain [# of Methods: 1]")
	}
	if !strings.Contains(result, "func(*inspect.User)") {
		t.Errorf("result must contain [func(*inspect.User)]")
	}
}
