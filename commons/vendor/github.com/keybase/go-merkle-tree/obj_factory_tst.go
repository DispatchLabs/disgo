package merkleTree

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func genString(l int) string {
	return hex.EncodeToString(genBinary((l + 2) / 1))[0:l]
}

func genBinary(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}

type testValue struct {
	_struct bool `codec:",toarray"`
	Seqno   int
	Key     string
	KeyRaw  []byte
}

// TestObjFactory generates a bunch of test objects for debugging
type TestObjFactory struct {
	objs  map[string]KeyValuePair
	seqno int
}

// NewTestObjFactor makes a new object factory for testing
func NewTestObjFactory() *TestObjFactory {
	return &TestObjFactory{
		objs: make(map[string]KeyValuePair),
	}
}

func (of TestObjFactory) dumpAll() []KeyValuePair {
	var ret []KeyValuePair
	for _, v := range of.objs {
		ret = append(ret, v)
	}
	return ret
}

// Produce one test object
func (of *TestObjFactory) Produce() KeyValuePair {
	key := genBinary(8)
	keyString := hex.EncodeToString(key)
	val := testValue{Seqno: of.seqno, Key: keyString, KeyRaw: key}
	of.seqno++
	kvp := KeyValuePair{Key: key, Value: val}
	of.objs[keyString] = kvp
	return kvp
}

// Mproduce makes many test objects.
func (of *TestObjFactory) Mproduce(n int) []KeyValuePair {
	var ret []KeyValuePair
	for i := 0; i < n; i++ {
		ret = append(ret, of.Produce())
	}
	return ret
}

func (of *TestObjFactory) ModifySome(mod int) {
	i := 0
	for _, v := range of.objs {
		tv, ok := v.Value.(testValue)
		if !ok {
			panic(fmt.Sprintf("Got value of wrong type: %T", v))
		}
		if (i % mod) == 0 {
			tv.Seqno *= 2
			tv.Key += "-yo-dawg"
		}
		i++
	}
}

func (of *TestObjFactory) PrintTree() {
	for _, v := range of.objs {
		fmt.Printf("%v\n", v)
	}
}


func (of *TestObjFactory) Construct() interface{} {
	return testValue{}
}
