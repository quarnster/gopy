// Copyright 2011 Julian Phillips.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package py

// #include "utils.h"
// static inline void decref(PyObject *obj) { Py_DECREF(obj); }
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

type BaseObject struct {
	AbstractObject
	o C.PyObject
}

var BaseType = (*Type)(unsafe.Pointer(&C.PyBaseObject_Type))

func newBaseObject(obj *C.PyObject) *BaseObject {
	return (*BaseObject)(unsafe.Pointer(obj))
}

// HasAttr returns true if "obj" has the attribute "name".  This is equivalent
// to the Python "hasattr(obj, name)".
func (obj *BaseObject) HasAttr(name Object) bool {
	ret := C.PyObject_HasAttr(c(obj), c(name))
	if ret == 1 {
		return true
	}
	return false
}

// HasAttrString returns true if "obj" has the attribute "name".  This is
// equivalent to the Python "hasattr(obj, name)".
func (obj *BaseObject) HasAttrString(name string) bool {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	ret := C.PyObject_HasAttrString(c(obj), s)
	if ret == 1 {
		return true
	}
	return false
}

// GetAttr returns the attribute of "obj" with the name "name".  This is
// equivalent to the Python "obj.name".
//
// Return value: New Reference.
func (obj *BaseObject) GetAttr(name Object) (Object, os.Error) {
	ret := C.PyObject_GetAttr(c(obj), c(name))
	return obj2ObjErr(ret)
}

// GetAttrString returns the attribute of "obj" with the name "name".  This is
// equivalent to the Python "obj.name".
//
// Return value: New Reference.
func (obj *BaseObject) GetAttrString(name string) (Object, os.Error) {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	ret := C.PyObject_GetAttrString(c(obj), s)
	return obj2ObjErr(ret)
}

// PyObject_GenericGetAttr : This is an internal helper function - we shouldn't
// need to expose it ...

// SetAttr sets the attribute of "obj" with the name "name" to "value".  This is
// equivalent to the Python "obj.name = value".
func (obj *BaseObject) SetAttr(name, value Object) os.Error {
	ret := C.PyObject_SetAttr(c(obj), c(name), c(value))
	return int2Err(ret)
}

// SetAttrString sets the attribute of "obj" with the name "name" to "value".
// This is equivalent to the Python "obj.name = value".
func (obj *BaseObject) SetAttrString(name string, value Object) os.Error {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	ret := C.PyObject_SetAttrString(c(obj), s, c(value))
	return int2Err(ret)
}

// PyObject_GenericSetAttr : This is an internal helper function - we shouldn't
// need to expose it ...

// DelAttr deletes the attribute with the name "name" from "obj".  This is
// equivalent to the Python "del obj.name".
func (obj *BaseObject) DelAttr(name Object) os.Error {
	ret := C.PyObject_SetAttr(c(obj), c(name), nil)
	return int2Err(ret)
}

// DelAttrString deletes the attribute with the name "name" from "obj".  This is
// equivalent to the Python "del obj.name".
func (obj *BaseObject) DelAttrString(name string) os.Error {
	s := C.CString(name)
	defer C.free(unsafe.Pointer(s))
	ret := C.PyObject_SetAttrString(c(obj), s, nil)
	return int2Err(ret)
}

// RichCompare compares "obj" with "obj2" using the specified operation (LE, GE
// etc.), and returns the result.  The equivalent Python is "obj op obj2", where
// op is the corresponding Python operator for op.
//
// Return value: New Reference.
func (obj *BaseObject) RichCompare(obj2 Object, op Op) (Object, os.Error) {
	ret := C.PyObject_RichCompare(c(obj), c(obj2), C.int(op))
	return obj2ObjErr(ret)
}

// RichCompare compares "obj" with "obj2" using the specified operation (LE, GE
// etc.), and returns true or false.  The equivalent Python is "obj op obj2",
// where op is the corresponding Python operator for op.
func (obj *BaseObject) RichCompareBool(obj2 Object, op Op) (bool, os.Error) {
	ret := C.PyObject_RichCompareBool(c(obj), c(obj2), C.int(op))
	return int2BoolErr(ret)
}

// PyObject_Cmp : Thanks to multiple return values, we don't need this function
// to be available in Go.

// Compare returns the result of comparing "obj" and "obj2".  This is equivalent
// to the Python "cmp(obj, obj2)".
func (obj *BaseObject) Compare(obj2 Object) (int, os.Error) {
	ret := C.PyObject_Compare(c(obj), c(obj2))
	return int(ret), exception()
}

// Repr returns a String representation of "obj".  This is equivalent to the
// Python "repr(obj)".
//
// Return value: New Reference.
func (obj *BaseObject) Repr() (Object, os.Error) {
	ret := C.PyObject_Repr(c(obj))
	return obj2ObjErr(ret)
}

// Str returns a String representation of "obj".  This is equivalent to the
// Python "str(obj)".
//
// Return value: New Reference.
func (obj *BaseObject) Str() (Object, os.Error) {
	ret := C.PyObject_Str(c(obj))
	return obj2ObjErr(ret)
}

// Bytes returns a Bytes representation of "obj".  This is equivalent to the
// Python "bytes(obj)".  In Python 2.x this method is identical to Str().
//
// Return value: New Reference.
func (obj *BaseObject) Bytes() (Object, os.Error) {
	ret := C.PyObject_Bytes(c(obj))
	return obj2ObjErr(ret)
}

// PyObject_Unicode : TODO

// IsInstance returns true if "obj" is an instance of "cls", false otherwise.
// If "cls" is a Type instead of a class, then true will be return if "obj" is
// of that type.  If "cls" is a Tuple then true will be returned if "obj" is an
// instance of any of the Objects in the tuple.  This is equivalent to the
// Python "isinstance(obj, cls)".
func (obj *BaseObject) IsInstance(cls Object) (bool, os.Error) {
	ret := C.PyObject_IsInstance(c(obj), c(cls))
	return int2BoolErr(ret)
}

// IsSubclass retuns true if "obj" is a Subclass of "cls", false otherwise.  If
// "cls" is a Tuple, then true is returned if "obj" is a Subclass of any member
// of "cls".  This is equivalent to the Python "issubclass(obj, cls)".
func (obj *BaseObject) IsSubclass(cls Object) (bool, os.Error) {
	ret := C.PyObject_IsSubclass(c(obj), c(cls))
	return int2BoolErr(ret)
}

func (obj *BaseObject) Call(args *Tuple, kwds *Dict) (Object, os.Error) {
	ret := C.PyObject_Call(c(obj), c(args), c(kwds))
	return obj2ObjErr(ret)
}

func (obj *BaseObject) CallObject(args *Tuple) (Object, os.Error) {
	var a *C.PyObject = nil
	if args != nil {
		a = c(args)
	}
	ret := C.PyObject_CallObject(c(obj), a)
	return obj2ObjErr(ret)
}

func (obj *BaseObject) CallFunction(format string, args ...interface{}) (Object, os.Error) {
	t, err := buildTuple(format, args...)
	if err != nil {
		return nil, err
	}
	return obj.CallObject(t)
}

func (obj *BaseObject) CallMethod(name string, format string, args ...interface{}) (Object, os.Error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	f := C.PyObject_GetAttrString(c(obj), cname)
	if f == nil {
		return nil, fmt.Errorf("AttributeError: %s", name)
	}
	defer C.decref(f)

	if C.PyCallable_Check(f) == 0 {
		return nil, fmt.Errorf("TypeError: attribute of type '%s' is not callable", name)
	}

	t, err := buildTuple(format, args...)
	if err != nil {
		return nil, err
	}

	ret := C.PyObject_CallObject(f, c(t))
	return obj2ObjErr(ret)
}

func (obj *BaseObject) CallFunctionObjArgs(args ...Object) (Object, os.Error) {
	t, err := PackTuple(args...)
	if err != nil {
		return nil, err
	}
	return obj.CallObject(t)
}

func (obj *BaseObject) CallMethodObjArgs(name string, args ...Object) (Object, os.Error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	f := C.PyObject_GetAttrString(c(obj), cname)
	if f == nil {
		return nil, fmt.Errorf("AttributeError: %s", name)
	}
	defer C.decref(f)

	if C.PyCallable_Check(f) == 0 {
		return nil, fmt.Errorf("TypeError: attribute of type '%s' is not callable", name)
	}

	t, err := PackTuple(args...)
	if err != nil {
		return nil, err
	}

	ret := C.PyObject_CallObject(f, c(t))
	return obj2ObjErr(ret)
}

// PyObject_Hash : TODO

// PyObject_HashNotImplement : This is an internal function, that we probably
// don't need to export.

// PyObject_IsTrue : Implemented on AbstractObject

// PyObject_Not : Implemented on AbstractObject

// PyObject_Type : Implemented on AbstractObject

// PyObject_TypeCheck : TODO

// Length returns the length of the Object.  This is equivalent to the Python
// "len(obj)".
func (obj *BaseObject) Length() (int64, os.Error) {
	ret := C.PyObject_Length(c(obj))
	return int64(ret), exception()
}

// Size returns the length of the Object.  This is equivalent to the Python
// "len(obj)".
func (obj *BaseObject) Size() (int64, os.Error) {
	ret := C.PyObject_Size(c(obj))
	return int64(ret), exception()
}

// GetItem returns the element of "obj" corresponding to "key".  This is
// equivalent to the Python "obj[key]".
//
// Return value: New Reference.
func (obj *BaseObject) GetItem(key Object) (Object, os.Error) {
	ret := C.PyObject_GetItem(c(obj), c(key))
	return obj2ObjErr(ret)
}

// SetItem sets the element of "obj" corresponding to "key" to "value".  This is
// equivalent to the Python "obj[key] = value".
func (obj *BaseObject) SetItem(key, value Object) os.Error {
	ret := C.PyObject_SetItem(c(obj), c(key), c(value))
	return int2Err(ret)
}

// DelItem deletes the element from "obj" that corresponds to "key".  This is
// equivalent to the Python "del obj[key]".
func (obj *BaseObject) DelItem(key Object) os.Error {
	ret := C.PyObject_DelItem(c(obj), c(key))
	return int2Err(ret)
}

// PyObject_AsFileDescriptor : TODO

func (obj *BaseObject) Dir() (Object, os.Error) {
	ret := C.PyObject_Dir(c(obj))
	return obj2ObjErr(ret)
}

// PyObject_GetIter : TODO