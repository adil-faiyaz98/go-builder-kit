package example

// Person represents a simple person
type Person struct {
	ID        string
	Name      string
	Age       int
	Email     string
	Addresses []Address
}

// Address represents a physical location
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
}
