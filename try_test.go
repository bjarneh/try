// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package try_test

import(
    "testing"
    "github.com/bjarneh/try"
    "os"
)

func TestTry(t *testing.T) {
    var err os.Error
    var ok bool
    var runtimeError *try.Error

    err = dangarous(true)
    runtimeError, ok = err.(*try.Error)

    if !ok {
        t.Fatalf("err != try.Error\n")
    }else{
        t.Logf("[error] %s\n", runtimeError)
        for _, s := range runtimeError.Trace {
            t.Logf("[trace] %s\n", s)
        }
    }

    err = dangarous(false)
    _, ok = err.(*try.Error)

    if ok {
        t.Fatalf("err == try.Error\n")
    }
}

func dangarous(fail bool) (e os.Error) {
    defer try.Catch(&e)
    if( fail ){
        panic("I've failed you")
    }else{
        e = os.NewError("Plain old os.Error")
    }
    return
}
