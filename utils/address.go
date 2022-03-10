package utils

type Addressed interface {
	GetAddress() (column, row int)
}

type Address struct {
	Column, Row int
}

func NewAddress(column int, row int) Address {
	return Address{Column: column, Row: row}
}

func (a Address) GetAddress() (column, row int) {
	return a.Column, a.Row
}
