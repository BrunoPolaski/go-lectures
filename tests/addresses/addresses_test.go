package addresses_test

import (
	"testing"

	"github.com/BrunoPolaski/go-lectures/tests/addresses"
)

func TestAddressType(t *testing.T) {
	t.Run("shall_return_road_if_road_is_passed", func(t *testing.T) {
		testString := "Alley Road"
		expected := "Road"
		result := addresses.AddressType(testString)

		if result != expected {
			t.Fatalf("\n Failed: %s returned. Expected: %s. \n", result, expected)
		}
	})

	t.Run("shall_return_invalid_type_if_nothing_is_passed", func(t *testing.T) {
		testString := ""
		expected := "Invalid address type"
		result := addresses.AddressType(testString)

		if result != expected {
			t.Fatalf("\n Failed: %s returned. Expected: %s. \n", result, expected)
		}
	})
}
