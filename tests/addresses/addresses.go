package addresses

import "strings"

func AddressType(address string) string {
	validTypes := []string{"road", "street", "avenue", "boulevard"}

	addressFirstWord := strings.Split(address, " ")[0]

	for _, validType := range validTypes {
		if strings.ToLower(addressFirstWord) == validType {
			return strings.Title(validType)
		}
	}

	return "Invalid address type"
}
