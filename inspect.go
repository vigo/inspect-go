package inspect

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

// Element displays the attributes of incoming object
func Element(input interface{}) (out string) {
	pad := func(input string, length int) string {
		if utf8.RuneCountInString(input) < length {
			return input + strings.Repeat(".", length-utf8.RuneCountInString(input))
		}
		return input
	}

	rt := reflect.TypeOf(input)
	rv := reflect.ValueOf(input).Elem()

	out += fmt.Sprintf("Inspection of: [%v] object\n\n", rt.Elem().Name())

	out += fmt.Sprintf(
		"[%v] %v\n",
		pad(fmt.Sprintf("%v", rt), 24),
		pad(fmt.Sprintf("%v", rv), 24),
	)

	if rv.Kind().String() == "struct" {
		out += fmt.Sprintf("# of Fields: %v\n\n", rv.NumField())
		for i := 0; i < rv.NumField(); i++ {
			f := rv.Type().Field(i)
			out += fmt.Sprintf(
				"%v %v %v %v\n",
				pad(f.Name, 20),
				pad(fmt.Sprintf("%s", f.Type), 10),
				pad(fmt.Sprintf("%v", rv.Field(i)), 30),
				// pad(fmt.Sprintf("%v", rv.Field(i).Interface()), 30),
				pad(fmt.Sprintf("%s", f.Tag), 30),
			)
		}

		if rt.NumMethod() > 0 {
			out += fmt.Sprintf("\n# of Methods: %v\n\n", rt.NumMethod())
			for i := 0; i < rt.NumMethod(); i++ {
				m := rt.Method(i)
				out += fmt.Sprintf(
					"%v %v\n",
					pad(m.Name, 20),
					m.Func.Type(),
				)
			}
		}
	}
	out += fmt.Sprintf("\n%s\n", strings.Repeat("-", 116))
	return out
}
