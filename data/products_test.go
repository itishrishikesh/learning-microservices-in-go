package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "nics",
		Price: 1.00,
		SKU:   "abc-abc-def",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
