package utils

type AddressCollection struct {
	addresses map[int]map[int]Address
}

func NewAddressCollection() *AddressCollection {
	return &AddressCollection{addresses: map[int]map[int]Address{}}
}

func (c *AddressCollection) Set(address Address) {
	column, row := address.GetAddress()

	rows, ok := c.addresses[column]
	if !ok {
		rows = map[int]Address{}
	}

	rows[row] = address

	c.addresses[column] = rows
}

func (c *AddressCollection) Has(column, row int) bool {
	rows, ok := c.addresses[column]
	if !ok {
		return false
	}

	_, ok = rows[row]
	if !ok {
		return false
	}

	return true
}

func (c *AddressCollection) Remove(address Addressed) {
	column, row := address.GetAddress()

	rows, ok := c.addresses[column]
	if !ok {
		rows = map[int]Address{}
	}

	delete(rows, row)

	c.addresses[column] = rows
}
