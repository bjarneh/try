// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package try

import(
    "fmt"
    "os"
    "io"
    "runtime"
)

type Error struct{
    Message string
    Trace []string
}

func (e *Error) String() string {
    return e.Message
}

func (e *Error) Report(w io.Writer) {
    fmt.Fprintf(w, "[error] %s\n", e.Message)
    for _, t := range e.Trace {
        fmt.Fprintf(w, "[trace] %s\n", t)
    }
}

func (e *Error) Die(w io.Writer) {
    e.Report(w)
    os.Exit(1)
}

// Indicate that we are a run time error (see runtime)
func (e *Error) RuntimeError() {}

func Catch(error *os.Error) {
    if r := recover(); r != nil {
        s := make([]string, 0)
        for i := 2; ; i++ {
            _, file, line, ok := runtime.Caller(i)
            if ok {
                s = append(s, fmt.Sprintf("%s: %d",file,line))
            }else{
                break
            }
        }
        *error = &Error{Message: fmt.Sprintf("%v", r), Trace: s}
    }
}
