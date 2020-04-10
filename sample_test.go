package tddbc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSay(t *testing.T) {
	actual := Say("Wow!")
	expected := "Wow! TDD BootCamp!!"

	if actual != expected {
		t.Errorf("actual=%s, expect=%s", actual, expected)
	}
}

func TestSay_testify(t *testing.T) {
	actual := Say("Hello!")
	assert.Equal(t, "Hello! TDD BootCamp!!", actual, "they should be equal")
}

// IntClosedRange構造体を新しく生成するための関数のテスト
func TestNewValidIntClosedRange(t *testing.T) {
	expected := IntClosedRange{
		upper: 1,
		lower: 3,
	}
	intClosedRange := NewIntClosedRange(1, 3)
	assert.Equal(t, intClosedRange, expected)
}

func TestNewInvalidIntClosedRange(t *testing.T) {
	expected := errors.New("err: uppper < lower")
	_, err := NewIntClosedRange(3, 1)
	assert.Equal(t, err, expected)
}
