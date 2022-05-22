package luvsqli

import (
	"fmt"
	"strings"
)

func ConvertStringToChars(payload string) string {
	charArr := make([]string, len(payload))
	for index, char := range payload {
		charArr[index] = fmt.Sprintf("CHR(%d)", char)
	}
	return strings.Join(charArr, "||")
}
