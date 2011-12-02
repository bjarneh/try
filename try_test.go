// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package try_test

import (
	"errors"
	"github.com/bjarneh/try"
	"testing"
)

func TestTry(t *testing.T) {

	var err error
	var ok bool
	var runtimeError *try.Error

	var fn []func(bool) error

	fn = append(fn, dangarous)
	fn = append(fn, dangarous2)

	for _, f := range fn {

		err = f(true)

		runtimeError, ok = err.(*try.Error)

		if !ok {
			t.Fatalf("err != try.Error\n")
		} else {
			t.Logf("[error] %s\n", runtimeError)
			for _, s := range runtimeError.Trace {
				t.Logf("[trace] %s\n", s)
			}
		}

		err = f(false)

		_, ok = err.(*try.Error)

		if ok {
			t.Fatalf("err == try.Error\n")
		}
	}

}

func dangarous(fail bool) (e error) {
	defer try.Catch(&e)
	if fail {
		panic("I've failed you")
	} else {
		e = errors.New("Plain old os.Error")
	}
	return
}

func dangarous2(fail bool) (e error) {
	defer try.Catch(&e)
	if fail {
		panic(errors.New("panic with os.Error"))
	} else {
		e = errors.New("Plain old os.Error")
	}
	return
}
