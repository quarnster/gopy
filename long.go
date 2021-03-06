// Copyright 2011 Julian Phillips.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py

// #include "utils.h"
// static inline long longCheck(PyObject *o) { return PyLong_Check(o); }
import "C"

import (
	"fmt"
	"unsafe"
)

type Long struct {
	AbstractObject
	NumberProtocol
	o C.PyLongObject
}

// LongType is the Type object that represents the Long type.
var LongType = (*Type)(unsafe.Pointer(C.getBasePyType(C.GoPyLong_Type)))

func longCheck(obj Object) bool {
	if obj == nil {
		return false
	}
	return C.longCheck(c(obj)) != 0
}

func newLong(obj *C.PyObject) *Long {
	return (*Long)(unsafe.Pointer(obj))
}

func NewLong(i int64) *Long {
	return newLong(C.PyLong_FromLongLong(C.PY_LONG_LONG(i)))
}

func (l *Long) Int64() int64 {
	// TODO: AsLongLong doesn't work for me on windows...
	return int64(C.PyLong_AsLong(c(l)))
}

func (l *Long) String() string {
	if l == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", l.Int64())
}
