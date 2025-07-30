package input


import (
	"testing"
	"reflect"
)


func TestCleanInput(t *testing.T) {

	cases := []struct {
	    input	string
	    expected	[]string
	}{
	    {
		input:		"   hello world   ",
		expected:	[]string{"hello", "world"},
	    },
	    {
		input:		"Charmander Bulbasaur PIKACHU",
		expected:	[]string{"charmander", "bulbasaur", "pikachu"},
	    },
	    {
		input:		"a   b",
		expected:	[]string{"a", "b"},
	    },
    	    {
		input:		"",
		expected:	[]string{},
	     },
	}
	
	for _, c := range cases {
		actual := CleanInput(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("CleanInput(%q) == %v, want %v", c.input, actual, c.expected)
		}
	}
}
