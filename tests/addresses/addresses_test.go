package addresses

import "testing"

func TestAddressType(t *testing.T) {
	testAddress := "aaawawaw 123"

	wantedAddressType := "Road"

	receivedAddressType := AddressType(testAddress)

	if receivedAddressType != wantedAddressType {
		t.Errorf("AddressType(%s) = %s; want %s", testAddress, receivedAddressType, wantedAddressType)
	}
}
