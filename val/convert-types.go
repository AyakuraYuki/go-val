package val

import "time"

func toSlice[V comparable](src []V) (dst []*V) {
	dst = make([]*V, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return
}

func toValueSlice[V comparable](src []*V) (dst []V) {
	dst = make([]V, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return
}

func toMap[K, V comparable](src map[K]V) (dst map[K]*V) {
	dst = make(map[K]*V, len(src))
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return
}

func toValueMap[K, V comparable](src map[K]*V) (dst map[K]V) {
	dst = make(map[K]V, len(src))
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return
}

// Bool

func Bool(v bool) *bool { return &v }

func BoolValue(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func BoolSlice(src []bool) []*bool                          { return toSlice(src) }
func BoolValueSlice(src []*bool) []bool                     { return toValueSlice(src) }
func BoolMap[K comparable](src map[K]bool) map[K]*bool      { return toMap(src) }
func BoolValueMap[K comparable](src map[K]*bool) map[K]bool { return toValueMap(src) }

// Int

func Int(v int) *int { return &v }

func IntValue(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func IntSlice(src []int) []*int                          { return toSlice(src) }
func IntValueSlice(src []*int) []int                     { return toValueSlice(src) }
func IntMap[K comparable](src map[K]int) map[K]*int      { return toMap(src) }
func IntValueMap[K comparable](src map[K]*int) map[K]int { return toValueMap(src) }

// Int8

func Int8(v int8) *int8 { return &v }

func Int8Value(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

func Int8Slice(src []int8) []*int8                          { return toSlice(src) }
func Int8ValueSlice(src []*int8) []int8                     { return toValueSlice(src) }
func Int8Map[K comparable](src map[K]int8) map[K]*int8      { return toMap(src) }
func Int8ValueMap[K comparable](src map[K]*int8) map[K]int8 { return toValueMap(src) }

// Int16

func Int16(v int16) *int16 { return &v }

func Int16Value(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

func Int16Slice(src []int16) []*int16                          { return toSlice(src) }
func Int16ValueSlice(src []*int16) []int16                     { return toValueSlice(src) }
func Int16Map[K comparable](src map[K]int16) map[K]*int16      { return toMap(src) }
func Int16ValueMap[K comparable](src map[K]*int16) map[K]int16 { return toValueMap(src) }

// Int32

func Int32(v int32) *int32 { return &v }

func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func Int32Slice(src []int32) []*int32                          { return toSlice(src) }
func Int32ValueSlice(src []*int32) []int32                     { return toValueSlice(src) }
func Int32Map[K comparable](src map[K]int32) map[K]*int32      { return toMap(src) }
func Int32ValueMap[K comparable](src map[K]*int32) map[K]int32 { return toValueMap(src) }

// Int64

func Int64(v int64) *int64 { return &v }

func Int64Value(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func Int64Slice(src []int64) []*int64                          { return toSlice(src) }
func Int64ValueSlice(src []*int64) []int64                     { return toValueSlice(src) }
func Int64Map[K comparable](src map[K]int64) map[K]*int64      { return toMap(src) }
func Int64ValueMap[K comparable](src map[K]*int64) map[K]int64 { return toValueMap(src) }

// UInt

func UInt(v uint) *uint { return &v }

func UIntValue(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

func UIntSlice(src []uint) []*uint                          { return toSlice(src) }
func UIntValueSlice(src []*uint) []uint                     { return toValueSlice(src) }
func UIntMap[K comparable](src map[K]uint) map[K]*uint      { return toMap(src) }
func UIntValueMap[K comparable](src map[K]*uint) map[K]uint { return toValueMap(src) }

// UInt8

func UInt8(v uint8) *uint8 { return &v }

func UInt8Value(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

func UInt8Slice(src []uint8) []*uint8                          { return toSlice(src) }
func UInt8ValueSlice(src []*uint8) []uint8                     { return toValueSlice(src) }
func UInt8Map[K comparable](src map[K]uint8) map[K]*uint8      { return toMap(src) }
func UInt8ValueMap[K comparable](src map[K]*uint8) map[K]uint8 { return toValueMap(src) }

// UInt16

func UInt16(v uint16) *uint16 { return &v }

func UInt16Value(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

func UInt16Slice(src []uint16) []*uint16                          { return toSlice(src) }
func UInt16ValueSlice(src []*uint16) []uint16                     { return toValueSlice(src) }
func UInt16Map[K comparable](src map[K]uint16) map[K]*uint16      { return toMap(src) }
func UInt16ValueMap[K comparable](src map[K]*uint16) map[K]uint16 { return toValueMap(src) }

// UInt32

func UInt32(v uint32) *uint32 { return &v }

func UInt32Value(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func UInt32Slice(src []uint32) []*uint32                          { return toSlice(src) }
func UInt32ValueSlice(src []*uint32) []uint32                     { return toValueSlice(src) }
func UInt32Map[K comparable](src map[K]uint32) map[K]*uint32      { return toMap(src) }
func UInt32ValueMap[K comparable](src map[K]*uint32) map[K]uint32 { return toValueMap(src) }

// UInt64

func UInt64(v uint64) *uint64 { return &v }

func UInt64Value(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func UInt64Slice(src []uint64) []*uint64                          { return toSlice(src) }
func UInt64ValueSlice(src []*uint64) []uint64                     { return toValueSlice(src) }
func UInt64Map[K comparable](src map[K]uint64) map[K]*uint64      { return toMap(src) }
func UInt64ValueMap[K comparable](src map[K]*uint64) map[K]uint64 { return toValueMap(src) }

// UIntPtr

func UIntptr(v uintptr) *uintptr { return &v }

func UIntptrValue(v *uintptr) uintptr {
	if v == nil {
		return 0
	}
	return *v
}

func UIntptrSlice(src []uintptr) []*uintptr                          { return toSlice(src) }
func UIntptrValueSlice(src []*uintptr) []uintptr                     { return toValueSlice(src) }
func UIntptrMap[K comparable](src map[K]uintptr) map[K]*uintptr      { return toMap(src) }
func UIntptrValueMap[K comparable](src map[K]*uintptr) map[K]uintptr { return toValueMap(src) }

// Float32

func Float32(v float32) *float32 { return &v }

func Float32Value(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

func Float32Slice(src []float32) []*float32                          { return toSlice(src) }
func Float32ValueSlice(src []*float32) []float32                     { return toValueSlice(src) }
func Float32Map[K comparable](src map[K]float32) map[K]*float32      { return toMap(src) }
func Float32ValueMap[K comparable](src map[K]*float32) map[K]float32 { return toValueMap(src) }

// Float64

func Float64(v float64) *float64 { return &v }

func Float64Value(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func Float64Slice(src []float64) []*float64                          { return toSlice(src) }
func Float64ValueSlice(src []*float64) []float64                     { return toValueSlice(src) }
func Float64Map[K comparable](src map[K]float64) map[K]*float64      { return toMap(src) }
func Float64ValueMap[K comparable](src map[K]*float64) map[K]float64 { return toValueMap(src) }

// Complex64

func Complex64(v complex64) *complex64 { return &v }

func Complex64Value(v *complex64) complex64 {
	if v == nil {
		return *(new(complex64))
	}
	return *v
}

func Complex64Slice(src []complex64) []*complex64                          { return toSlice(src) }
func Complex64ValueSlice(src []*complex64) []complex64                     { return toValueSlice(src) }
func Complex64Map[K comparable](src map[K]complex64) map[K]*complex64      { return toMap(src) }
func Complex64ValueMap[K comparable](src map[K]*complex64) map[K]complex64 { return toValueMap(src) }

// Complex128

func Complex128(v complex128) *complex128 { return &v }

func Complex128Value(v *complex128) complex128 {
	if v == nil {
		return *(new(complex128))
	}
	return *v
}

func Complex128Slice(src []complex128) []*complex128                          { return toSlice(src) }
func Complex128ValueSlice(src []*complex128) []complex128                     { return toValueSlice(src) }
func Complex128Map[K comparable](src map[K]complex128) map[K]*complex128      { return toMap(src) }
func Complex128ValueMap[K comparable](src map[K]*complex128) map[K]complex128 { return toValueMap(src) }

// Byte

func Byte(v byte) *byte { return &v }

func ByteValue(v *byte) byte {
	if v == nil {
		return 0
	}
	return *v
}

func ByteSlice(src []byte) []*byte                          { return toSlice(src) }
func ByteValueSlice(src []*byte) []byte                     { return toValueSlice(src) }
func ByteMap[K comparable](src map[K]byte) map[K]*byte      { return toMap(src) }
func ByteValueMap[K comparable](src map[K]*byte) map[K]byte { return toValueMap(src) }

// Rune

func Rune(v rune) *rune { return &v }

func RuneValue(v *rune) rune {
	if v == nil {
		return 0
	}
	return *v
}

func RuneSlice(src []rune) []*rune                          { return toSlice(src) }
func RuneValueSlice(src []*rune) []rune                     { return toValueSlice(src) }
func RuneMap[K comparable](src map[K]rune) map[K]*rune      { return toMap(src) }
func RuneValueMap[K comparable](src map[K]*rune) map[K]rune { return toValueMap(src) }

// String

func String(v string) *string { return &v }

func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func StringSlice(src []string) []*string                          { return toSlice(src) }
func StringValueSlice(src []*string) []string                     { return toValueSlice(src) }
func StringMap[K comparable](src map[K]string) map[K]*string      { return toMap(src) }
func StringValueMap[K comparable](src map[K]*string) map[K]string { return toValueMap(src) }

// Time

func Time(v time.Time) *time.Time { return &v }

func TimeValue(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}

func SecondsTimeValue(v *int64) time.Time {
	if v == nil {
		return time.Time{}
	}
	return time.Unix(*v/1000, 0)
}

func MillisecondsTimeValue(v *int64) time.Time {
	if v == nil {
		return time.Time{}
	}
	return time.Unix(0, *v*1000000)
}

func TimeUnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond/time.Nanosecond)
}

func TimeSlice(src []time.Time) []*time.Time                          { return toSlice(src) }
func TimeValueSlice(src []*time.Time) []time.Time                     { return toValueSlice(src) }
func TimeMap[K comparable](src map[K]time.Time) map[K]*time.Time      { return toMap(src) }
func TimeValueMap[K comparable](src map[K]*time.Time) map[K]time.Time { return toValueMap(src) }

// Ptr

// Ptr returns a pointer to its argument.
// It can be used to initialize pointer fields.
func Ptr[T any](t T) *T { return &t }

func Val[T any](v *T) T {
	if v == nil {
		return *new(T)
	}
	return *v
}
