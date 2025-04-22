package val

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVal(t *testing.T) {
	t.Run("default on nil", func(t *testing.T) {
		assert.Equal(t, false, Val[bool](nil))

		assert.Equal(t, 0, Val[int](nil))
		assert.Equal(t, int8(0), Val[int8](nil))
		assert.Equal(t, int16(0), Val[int16](nil))
		assert.Equal(t, int32(0), Val[int32](nil))
		assert.Equal(t, int64(0), Val[int64](nil))

		assert.Equal(t, uint(0), Val[uint](nil))
		assert.Equal(t, uint8(0), Val[uint8](nil))
		assert.Equal(t, uint16(0), Val[uint16](nil))
		assert.Equal(t, uint32(0), Val[uint32](nil))
		assert.Equal(t, uint64(0), Val[uint64](nil))
		assert.Equal(t, uintptr(0), Val[uintptr](nil))

		assert.Equal(t, float32(0.0), Val[float32](nil))
		assert.Equal(t, 0.0, Val[float64](nil))

		assert.Equal(t, *new(complex64), Val[complex64](nil))
		assert.Equal(t, *new(complex128), Val[complex128](nil))

		assert.Equal(t, byte(0), Val[byte](nil))
		assert.Equal(t, rune(0), Val[rune](nil))
		assert.Equal(t, "", Val[string](nil))
		assert.Equal(t, time.Time{}, Val[time.Time](nil))
	})

	t.Run("general use case", func(t *testing.T) {
		assert.Equal(t, true, Val[bool](Bool(true)))

		assert.Equal(t, 123, Val[int](Int(123)))
		assert.Equal(t, int8(123), Val[int8](Int8(123)))
		assert.Equal(t, int16(123), Val[int16](Int16(123)))
		assert.Equal(t, int32(123), Val[int32](Int32(123)))
		assert.Equal(t, int64(123), Val[int64](Int64(123)))

		assert.Equal(t, uint(123), Val[uint](UInt(123)))
		assert.Equal(t, uint8(123), Val[uint8](UInt8(123)))
		assert.Equal(t, uint16(123), Val[uint16](UInt16(123)))
		assert.Equal(t, uint32(123), Val[uint32](UInt32(123)))
		assert.Equal(t, uint64(123), Val[uint64](UInt64(123)))
		assert.Equal(t, uintptr(123), Val[uintptr](UIntptr(123)))

		assert.Equal(t, float32(12.34), Val[float32](Float32(12.34)))
		assert.Equal(t, 12.34, Val[float64](Float64(12.34)))

		assert.Equal(t, complex64(1+4i), Val[complex64](Complex64(1+4i)))
		assert.Equal(t, 1+4i, Val[complex128](Complex128(1+4i)))

		assert.Equal(t, byte(1), Val[byte](Byte(1)))
		assert.Equal(t, 'a', Val[rune](Rune('a')))
		assert.Equal(t, "foo", Val[string](String("foo")))

		now := time.Now()
		assert.Equal(t, now, Val[time.Time](Time(now)))
	})
}
