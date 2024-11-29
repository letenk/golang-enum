package main

import "fmt"

type StatusOrder int

const (
	Pending StatusOrder = iota
	Processed
	Shipped
	Delivered
	Cancelled
)

// String method for make representation string value from StatusOrder
func (s StatusOrder) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Processed:
		return "Processed"
	case Shipped:
		return "Shipped"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

// IsValid method for checking enum value is valid
func (s StatusOrder) IsValid() bool {
	switch s {
	case Pending, Processed, Shipped, Delivered, Cancelled:
		return true

	default:
		return false
	}
}

func main() {

	processed := Processed
	fmt.Printf("Order Status: %s (%d)\n", processed, processed)

	pending := Pending
	fmt.Printf("Order Status: %s (%d)\n", pending, pending)

}
