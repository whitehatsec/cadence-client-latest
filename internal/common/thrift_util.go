// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package common

import (
	"reflect"

	// "github.com/apache/thrift/lib/go/thrift"
	thriftV0100 "github.com/apache/thrift/thriftV0100/lib/go/thrift"
)

// TSerialize is used to serialize thrift TStruct to []byte
func TSerialize(t thriftV0100.TStruct) (b []byte, err error) {
	return thriftV0100.NewTSerializer().Write(t)
}

// TListSerialize is used to serialize list of thrift TStruct to []byte
func TListSerialize(ts []thriftV0100.TStruct) (b []byte, err error) {
	if ts == nil {
		return
	}

	t := thriftV0100.NewTSerializer()
	t.Transport.Reset()

	// NOTE: we don't write any markers as thrift by design being a streaming protocol doesn't
	// recommend writing length.

	for _, v := range ts {
		if e := v.Write(t.Protocol); e != nil {
			err = thriftV0100.PrependError("error writing TStruct: ", e)
			return
		}
	}

	if err = t.Protocol.Flush(); err != nil {
		return
	}

	if err = t.Transport.Flush(); err != nil {
		return
	}

	b = t.Transport.Bytes()
	return
}

// TDeserialize is used to deserialize []byte to thrift TStruct
func TDeserialize(t thriftV0100.TStruct, b []byte) (err error) {
	return thriftV0100.NewTDeserializer().Read(t, b)
}

// TListDeserialize is used to deserialize []byte to list of thrift TStruct
func TListDeserialize(ts []thriftV0100.TStruct, b []byte) (err error) {
	t := thriftV0100.NewTDeserializer()
	err = nil
	if _, err = t.Transport.Write(b); err != nil {
		return
	}

	for i := 0; i < len(ts); i++ {
		if e := ts[i].Read(t.Protocol); e != nil {
			err = thriftV0100.PrependError("error reading TStruct: ", e)
			return
		}
	}

	return
}

// IsUseThriftEncoding checks if the objects passed in are all encoded using thriftV0100.
func IsUseThriftEncoding(objs []interface{}) bool {
	// NOTE: our criteria to use which encoder is simple if all the types are serializable using thrift then we use
	// thrift encoder. For everything else we default to gob.

	if len(objs) == 0 {
		return false
	}

	for i := 0; i < len(objs); i++ {
		if !IsThriftType(objs[i]) {
			return false
		}
	}
	return true
}

// IsUseThriftDecoding checks if the objects passed in are all de-serializable using thriftV0100.
func IsUseThriftDecoding(objs []interface{}) bool {
	// NOTE: our criteria to use which encoder is simple if all the types are de-serializable using thrift then we use
	// thrift decoder. For everything else we default to gob.

	if len(objs) == 0 {
		return false
	}

	for i := 0; i < len(objs); i++ {
		rVal := reflect.ValueOf(objs[i])
		if rVal.Kind() != reflect.Ptr || !IsThriftType(reflect.Indirect(rVal).Interface()) {
			return false
		}
	}
	return true
}

// IsThriftType checks whether the object passed in is a thrift encoded object.
func IsThriftType(v interface{}) bool {
	// NOTE: Thrift serialization works only if the values are pointers.
	// Thrift has a validation that it meets thift.TStruct which has Read/Write pointer receivers.

	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return false
	}
	t := reflect.TypeOf((*thriftV0100.TStruct)(nil)).Elem()
	return reflect.TypeOf(v).Implements(t)
}
