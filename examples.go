package testutils

import (
	"reflect"
	"testing"

	. "github.com/kynrai/go-testutils"
)

func TestExamplet(t *testing.T) {
	type Stuff struct {
		Name   string
		Type   string
		Number int
	}

	stuff1 := Stuff{Name: "Bob", Type: "Human", Age: 10}
	stuff2 := Stuff{Name: "Bob", Type: "Dog", Age: 10}

	Cases{
		"stuff": {stuff1, stuff2},
	}.Assert(t, reflect.DeepEqual)
}
