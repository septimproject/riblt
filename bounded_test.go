package riblt

import (
	"errors"
	"testing"
)

func TestDecodeBudgetStopsPeeling(t *testing.T) {
	var d Decoder[testSymbol]
	var e Encoder[testSymbol]
	e.AddSymbol(testSymbol{42})
	d.AddCodedSymbol(e.ProduceNextCodedSymbol())
	work := 0
	if err := d.TryDecodeWithBudget(&work); !errors.Is(err, ErrDecodeBudget) {
		t.Fatalf("budget refusal: %v", err)
	}
}
func TestInvalidDegreeReturnsError(t *testing.T) {
	var d Decoder[testSymbol]
	d.cs = []CodedSymbol[testSymbol]{{Count: 2}}
	d.decodable = []int{0}
	work := 4
	if err := d.TryDecodeWithBudget(&work); !errors.Is(err, ErrInvalidDegree) {
		t.Fatalf("invalid input: %v", err)
	}
}
func TestMappingAlwaysAdvances(t *testing.T) {
	for _, hash := range []uint64{0, 1, ^uint64(0), 0x0123456789abcdef} {
		m := randomMapping{prng: hash}
		for i := 0; i < 32; i++ {
			old := m.lastIndex
			next := m.nextIndex()
			if next <= old && old != int(^uint(0)>>1) {
				t.Fatalf("mapping did not advance: %d -> %d", old, next)
			}
		}
	}
}

func TestMappingProgressAfterRoundedUnitRatio(t *testing.T) {
	const multiplier uint64 = 0xda942042e4dd58b5
	inverse := uint64(1)
	for i := 0; i < 6; i++ {
		inverse *= 2 - multiplier*inverse
	}
	m := randomMapping{prng: (^uint64(0)) * inverse}
	if next := m.nextIndex(); next != 1 {
		t.Fatalf("rounded zero increment must advance once, got %d", next)
	}
}
