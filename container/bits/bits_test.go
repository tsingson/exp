package bits

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(3)
	if !bits.IsBitSet(3) {
		t.Errorf("Expected bit 3 to be set")
	}
	bits.SetBit(129)
	if !bits.IsBitSet(129) {
		t.Errorf("Expected bit 129 to be set")
	}
}

func TestClearBit(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(3)
	bits.ClearBit(3)
	if bits.IsBitSet(3) {
		t.Errorf("Expected bit 3 to be clear")
	}
	bits.SetBit(129)
	bits.ClearBit(129)
	if bits.IsBitSet(129) {
		t.Errorf("Expected bit 129 to be clear")
	}
}

func TestIsBitSet(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(3)
	if !bits.IsBitSet(3) {
		t.Errorf("Expected bit 3 to be set")
	}
	if bits.IsBitSet(4) {
		t.Errorf("Expected bit 4 to be clear")
	}
	bits.SetBit(129)
	if !bits.IsBitSet(129) {
		t.Errorf("Expected bit 129 to be set")
	}
}

func TestCountOnes(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(3)
	bits.SetBit(64)
	bits.SetBit(129)
	if bits.CountOnes() != 3 {
		t.Errorf("Expected 3 bits to be set, got %d", bits.CountOnes())
	}
}

func TestLeftShift(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(3)
	bits.LeftShift(1)
	if bits.IsBitSet(3) {
		t.Errorf("Expected bit 3 to be clear after left shift")
	}
	if !bits.IsBitSet(4) {
		t.Errorf("Expected bit 4 to be set after left shift")
	}
	bits.SetBit(129)
	bits.LeftShift(1)
	if bits.IsBitSet(129) {
		t.Errorf("Expected bit 129 to be clear after left shift")
	}

	bits.LeftShift(130)
	if bits.CountOnes() != 0 {
		t.Errorf("Expected all bits to be clear after left shift beyond size")
	}
}

func TestRightShift(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(3)
	bits.RightShift(1)
	if bits.IsBitSet(3) {
		t.Errorf("Expected bit 3 to be clear after right shift")
	}
	if !bits.IsBitSet(2) {
		t.Errorf("Expected bit 2 to be set after right shift")
	}
	bits.SetBit(129)
	bits.RightShift(1)
	if bits.IsBitSet(129) {
		t.Errorf("Expected bit 129 to be clear after right shift")
	}

	bits.RightShift(130)
	if bits.CountOnes() != 0 {
		t.Errorf("Expected all bits to be clear after right shift beyond size")
	}
}

func TestLeftShiftAcrossWords(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(63)
	bits.LeftShift(1)
	if bits.IsBitSet(63) {
		t.Errorf("Expected bit 63 to be clear after left shift")
	}
	if !bits.IsBitSet(64) {
		t.Errorf("Expected bit 64 to be set after left shift")
	}
}

func TestRightShiftAcrossWords(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(64)
	bits.RightShift(1)
	if bits.IsBitSet(64) {
		t.Errorf("Expected bit 64 to be clear after right shift")
	}
	if !bits.IsBitSet(63) {
		t.Errorf("Expected bit 63 to be set after right shift")
	}
}

func TestLeftShiftEdgeCases(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(128)
	bits.LeftShift(2)
	if bits.IsBitSet(128) || bits.IsBitSet(129) {
		t.Errorf("Expected bits 128 and 129 to be clear after left shift")
	}
}

func TestRightShiftEdgeCases(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(1)
	bits.RightShift(2)
	if bits.IsBitSet(0) || bits.IsBitSet(1) {
		t.Errorf("Expected bits 0 and 1 to be clear after right shift")
	}
}

func TestSetAndClearOutOfRange(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(130)
	if bits.IsBitSet(130) {
		t.Errorf("Expected bit 130 to be out of range and not set")
	}
	bits.ClearBit(130)
	// No need to check this, as ClearBit does nothing out of range
}

func TestRightShiftBeyondSize(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(129)
	bits.RightShift(130)
	if bits.CountOnes() != 0 {
		t.Errorf("Expected all bits to be clear after right shift beyond size")
	}
}

func TestLeftShiftBeyondSize(t *testing.T) {
	bits := NewBits(130)
	bits.SetBit(0)
	bits.LeftShift(130)
	if bits.CountOnes() != 0 {
		t.Errorf("Expected all bits to be clear after left shift beyond size")
	}
}

func TestLeftShiftFull(t *testing.T) {
	bits := NewBits(2560)
	for i := 0; i < 2560; i++ {
		bits.SetBit(i)
	}
	n := bits.CountOnes()
	if n != 2560 {
		t.Errorf("Expected all bits to be set, got %d", n)
	}

	for i := 0; i < 3000; i++ {
		bits.LeftShift(1)
		if bits.CountOnes() != n-1 {
			t.Errorf("Expected %d bits to be set, got %d", n-1, bits.CountOnes())
		}
		if i < 2559 {
			n--
		}

	}
}

func TestRightShiftFull(t *testing.T) {
	bits := NewBits(2560)
	for i := 0; i < 2560; i++ {
		bits.SetBit(i)
	}
	n := bits.CountOnes()
	if n != 2560 {
		t.Errorf("Expected all bits to be set, got %d", n)
	}

	for i := 0; i < 3000; i++ {
		bits.RightShift(1)
		if bits.CountOnes() != n-1 {
			t.Errorf("Expected %d bits to be set, got %d", n-1, bits.CountOnes())
		}
		if i < 2559 {
			n--
		}

	}
}
