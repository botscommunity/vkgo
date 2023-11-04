package conv

import "fmt"

func BooleanToInteger(property bool) int {
	if value := 1; property {
		return value
	}

	return 0
}

func ForParams[element any](array []element) string {
	var (
		list  string
		count = len(array)
	)

	for index, field := range array {
		if list += fmt.Sprintf("%v", field); index != count-1 {
			list += ","
		}
	}

	return list
}
