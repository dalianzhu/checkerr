package checkerr

import (
	"errors"
	"log"
	"reflect"
	"testing"
)

func TestCheckErr(t *testing.T) {
	err := testError1()
	log.Printf("err %v %v", err, reflect.TypeOf(err))
	if err == nil {
		t.Fail()
	}
	if err.Error() != "hello" {
		t.Fail()
	}

	err = testError2()
	if err != nil {
		t.Fail()
	}

	err = testError3()
	if err.Error() != "check hello" {
		t.Fail()
	}
}

func testError1() (err error) {
	defer MarkPanic(&err)

	err = errors.New("hello")
	CheckError(err)

	return nil
}

func testError2() (err error) {
	defer MarkPanic(&err)
	CheckError(nil)
	return nil
}

func testError3() (err error) {
	defer MarkPanic(&err)

	CheckErrorf(nil, "check %v", "haha")

	err = errors.New("hello")
	CheckErrorf(err, "check %v", err)
	return nil
}
