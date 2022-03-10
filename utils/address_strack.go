package utils

type AddressStack struct {
	addresses []Addressed
}

func NewAddressStack() *AddressStack {
	return &AddressStack{addresses: []Addressed{}}
}

func (s *AddressStack) Push(address Addressed) {
	s.addresses = append(s.addresses, address)
}

func (s *AddressStack) Pop() Addressed {
	if len(s.addresses) == 0 {
		return nil
	}

	address := s.addresses[len(s.addresses)-1]
	s.addresses = s.addresses[:len(s.addresses)-1]

	return address
}

func (s *AddressStack) Empty() bool {
	return len(s.addresses) == 0
}
