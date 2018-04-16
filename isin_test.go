package isin

import "testing"

var validIsins = []string{
	"SE0000356008",
	"SE0006593919",
	"SE0000101297",
	"SE0000526626",
	"SE0010415281",
	"SE0002216713",
	"SE0009922164",
	"SE0009778848",
	"SE0005794617",
	"SE0009973449",
}

var invalidIsins = []string{
	"SE0008342008",
	"SE0625365919",
	"SE0065464597",
	"SE0075632446",
	"SE0644215281",
	"SE0064531713",
	"SE5463992164",
	"SE0087658848",
	"SE5624354617",
	"SE0745654349",
}

func TestValidate(t *testing.T) {

	for _, isin := range validIsins {
		resp := Validate(isin)
		if resp != true {
			t.Errorf("%s is returned as %v, expected to be false", isin, resp)
		}
	}

	for _, isin := range invalidIsins {
		resp := Validate(isin)
		if resp != false {
			t.Errorf("%s is returned as %v, expected to be invalid", isin, resp)
		}
	}
}
