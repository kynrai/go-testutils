package testutils

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

type Cases map[string]struct {
	Got, Want interface{}
}

func (c Cases) Assert(t *testing.T, f func(x, y interface{}) bool) {
	for name, tc := range c {
		if got, want := tc.Got, tc.Want; !f(got, want) {
			t.Fatalf("diff for case %q\n%s", name, pretty.Compare(got, want))
		}
	}
}
