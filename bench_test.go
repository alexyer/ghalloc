package ghalloc

import (
	"testing"
	"unsafe"
)

type TestStruct struct {
	Field [128]byte
}

type TestStructLarge struct {
	Field [1024]byte
}

const NUM int = 1000

var (
	structs      []*TestStruct
	structsLarge []*TestStructLarge
)

func init() {
	structs = make([]*TestStruct, NUM)
	structsLarge = make([]*TestStructLarge, NUM)
}

func BenchmarkGhallocAlloc(b *testing.B) {
	b.StopTimer()
	gh, _ := New(&Options{})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		str := (*TestStruct)(gh.Alloc(unsafe.Sizeof(TestStruct{})))
		structs[i%NUM] = str
	}
}

//func BenchmarkGhallocAlloc8(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(8)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := (*TestStruct)(gh.Alloc(unsafe.Sizeof(TestStruct{})))
//str.Field = 5
//structs[i%NUM] = str
//}
//}

//func BenchmarkGhallocAlloc64(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(64)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := (*TestStruct)(gh.Alloc(unsafe.Sizeof(TestStruct{})))
//str.Field = 5
//structs[i%NUM] = str
//}
//}

//func BenchmarkGhallocAlloc128(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(128)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := (*TestStruct)(gh.Alloc(unsafe.Sizeof(TestStruct{})))
//str.Field = 5
//structs[i%NUM] = str
//}
//}

//func BenchmarkNativeAlloc(b *testing.B) {
//for i := 0; i < b.N; i++ {
//str := &TestStruct{}
//if str != nil {
//str.Field = 5
//structs[i%NUM] = str
//}
//}
//}

//func BenchmarkNativeAlloc8(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(8)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := &TestStruct{}
//str.Field = 5
//structs[i%NUM] = str
//}
//}

//func BenchmarkNativeAlloc64(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(64)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := &TestStruct{}
//str.Field = 5
//structs[i%NUM] = str
//}
//}

//func BenchmarkNativeAlloc128(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(128)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := &TestStruct{}
//str.Field = 5
//structs[i%NUM] = str
//}
//}

//func BenchmarkGhallocAllocLarge(b *testing.B) {
//b.StopTimer()
//gh, _ := New(&Options{})
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := (*TestStructLarge)(gh.Alloc(unsafe.Sizeof(TestStructLarge{})))
//structsLarge[i%NUM] = str
//}
//}

//func BenchmarkGhallocAllocLarge8(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(8)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := (*TestStructLarge)(gh.Alloc(unsafe.Sizeof(TestStructLarge{})))
//structsLarge[i%NUM] = str
//}
//}

//func BenchmarkGhallocAllocLarge64(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(64)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := (*TestStructLarge)(gh.Alloc(unsafe.Sizeof(TestStructLarge{})))
//structsLarge[i%NUM] = str
//}
//}

//func BenchmarkGhallocAllocLarge128(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(128)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := (*TestStructLarge)(gh.Alloc(unsafe.Sizeof(TestStructLarge{})))
//structsLarge[i%NUM] = str
//}
//}

//func BenchmarkNativeAllocLarge(b *testing.B) {
//for i := 0; i < b.N; i++ {
//str := &TestStructLarge{}
//structsLarge[i%NUM] = str
//}
//}

//func BenchmarkNativeAllocLarge8(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(8)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := &TestStructLarge{}
//structsLarge[i%NUM] = str
//}
//}

//func BenchmarkNativeAllocLarge64(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(64)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := &TestStructLarge{}
//structsLarge[i%NUM] = str
//}
//}

//func BenchmarkNativeAllocLarge128(b *testing.B) {
//b.StopTimer()
//b.SetParallelism(128)
//b.StartTimer()

//for i := 0; i < b.N; i++ {
//str := &TestStructLarge{}
//structsLarge[i%NUM] = str
//}
//}
