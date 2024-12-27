package addresses

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AddressType(address string) string {
	caser := cases.Title(language.English)

	validTypes := []string{"road", "street", "avenue", "boulevard"}
	addressSplitted := strings.Split(address, " ")
	addressLastWord := addressSplitted[len(addressSplitted)-1]

	for _, validType := range validTypes {
		if strings.ToLower(addressLastWord) == validType {
			return caser.String(validType)
		}
	}

	return "Invalid address type"
}
