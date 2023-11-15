package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(2), answer)
}
func TestSubstraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Substraction(1, 1)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(0), answer)
}
func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multitplication(1, 1)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(1), answer)
}
func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Errorf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, int32(1), answer)
}
