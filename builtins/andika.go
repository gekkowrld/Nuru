package builtins

import (
	"fmt"
	"github.com/AvicennaJr/Nuru/object"
	"strings"
)

func Andika(args ...object.Object) object.Object {
	if len(args) == 0 {
		fmt.Println("")
		return nil
	}

	var arr []string
	var types []string
	for _, arg := range args {
		if arg == nil {
			return newError("Hauwezi kufanya operesheni hii")
		}
		arr = append(arr, arg.Inspect())
		types = append(types, Aina(arg).Inspect())
	}

	passedType := types[0]
	if passedType == "NENO" {
		nStr := arr[1:]
		types = types[1:]
		drrr := formatString(arr[0], nStr, types)
		fmt.Println(drrr)
		return nil
	}

	str := strings.Join(arr, " ")
	fmt.Println(str)
	return nil
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func formatString(fmt string, vals []string, types []string) string {
	var newStr strings.Builder
	pos := 0
	lv := len(vals)
	for i := 0; i < len(fmt); i++ {
		if fmt[i] == '%' && i+1 < len(fmt) {
			switch fmt[i+1] {
			case 's':
				if pos < lv && (types[pos] == "NENO") {
					newStr.WriteString(vals[pos])
					pos++
					i++
					continue
				}
			case 'd':
				if pos < lv && (types[pos] == "NAMBA") {
					newStr.WriteString(vals[pos])
					pos++
					i++
					continue
				}
			}
		} else {
			newStr.WriteByte(fmt[i])
		}
	}
	return newStr.String()
}
