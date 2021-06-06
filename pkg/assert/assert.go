package assert

import (
	"reflect"
)

type ErrorAsserter interface {
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
}

func AssertEqual(t ErrorAsserter, a, b interface{}) {
	if a != b {
		t.Fatalf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

func AssertTrue(t ErrorAsserter, predicate bool, errMsg string) {
	if !predicate {
		t.Fatal(errMsg)
	}
}