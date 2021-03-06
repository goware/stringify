package stringify

import (
	"math/big"
	"reflect"
	"testing"
	// "github.com/kr/pretty"
)

func TestStringifyMap(t *testing.T) {
	src := map[string]interface{}{
		"a": 123,
		"b": big.NewInt(3232332),
		"c": map[string]interface{}{
			"b": 456,
		},
		"d": []byte{1, 2},
		"e": []*big.Int{big.NewInt(1), big.NewInt(2)},
	}

	dest := map[string]interface{}{}

	err := StringifyMap(src, dest)
	assertNoError(t, err)

	assertEqual(t, "123", dest["a"])

	// pretty.Println(dest)
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expecting '%v' but received '%v", expected, actual)
		t.Fail()
	}
}
