package id

// AccountID defines account id object
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// TripID defines account id object
type TripID string

func (t TripID) String() string {
	return string(t)
}

// IdentifyID defines identify id object
type IdentifyID string

func (i IdentifyID) String() string {
	return string(i)
}

// CarID defines identify car
type CarID string

func (i CarID) String() string {
	return string(i)
}
