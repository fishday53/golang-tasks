//go:build !solution

package testequal

import (
	"reflect"
)

func compareSlicesInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareSlicesByte(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareMap(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	for key := range a {
		if _, ok := b[key]; !ok {
			return false
		}
		if a[key] != b[key] {
			return false
		}
	}
	return true
}

func compare(left, right interface{}) bool {

	l := reflect.TypeOf(left)
	r := reflect.TypeOf(right)

	if l != r {
		return false
	}

	switch left.(type) {
	case int:
		return left == right
	case int8:
		return left == right
	case int16:
		return left == right
	case int32:
		return left == right
	case int64:
		return left == right
	case uint8:
		return left == right
	case uint16:
		return left == right
	case uint32:
		return left == right
	case uint64:
		return left == right
	case string:
		return left == right
	case []int:
		return compareSlicesInt(left.([]int), right.([]int))
	case []byte:
		return compareSlicesByte(left.([]byte), right.([]byte))
	case map[string]string:
		return compareMap(left.(map[string]string), right.(map[string]string))
	default:
		return false
	}
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if !compare(expected, actual) {
		switch len(msgAndArgs) {
		case 0:
		case 1:
			t.Errorf(msgAndArgs[0].(string))
		default:
			args := msgAndArgs[1:]
			t.Errorf(msgAndArgs[0].(string), args...)
		}
		return false
	}
	return true
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if compare(expected, actual) {
		switch len(msgAndArgs) {
		case 0:
		case 1:
			t.Errorf(msgAndArgs[0].(string))
		default:
			args := msgAndArgs[1:]
			t.Errorf(msgAndArgs[0].(string), args...)
		}
		t.Helper()
		return false
	}
	return true
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if !compare(expected, actual) {
		switch len(msgAndArgs) {
		case 0:
		case 1:
			t.Errorf(msgAndArgs[0].(string))
		default:
			args := msgAndArgs[1:]
			t.Errorf(msgAndArgs[0].(string), args...)
		}
		//panic(msgAndArgs)
		//os.Exit(1)
		//runtime.Goexit()
		t.FailNow()
	}
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if compare(expected, actual) {
		switch len(msgAndArgs) {
		case 0:
		case 1:
			t.Errorf(msgAndArgs[0].(string))
		default:
			args := msgAndArgs[1:]
			t.Errorf(msgAndArgs[0].(string), args...)
		}
		//panic(msgAndArgs)
		//os.Exit(1)
		//runtime.Goexit()
		t.FailNow()
	}
}
