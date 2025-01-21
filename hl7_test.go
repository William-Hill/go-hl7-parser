package hl7_test

type my7 struct {
	FirstName string `hl7:"PID.5.2"`
	LastName  string `hl7:"PID.5.1"`
}
