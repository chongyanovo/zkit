package atomicx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string
}

func TestValue_NewValue(t *testing.T) {
	testCases := []struct {
		name  string
		input *Person
	}{
		{
			name: "nil",
		},
		{
			name: "not nil",
			input: &Person{
				Name: "Tom",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val := NewValueOf[*Person](tc.input)
			assert.Equal(t, tc.input, val.Load())
		})
	}
}

func TestValue_Load_Store(t *testing.T) {
	testCases := []struct {
		name    string
		input   *Person
		wantVal *Person
	}{
		{
			name: "nil",
		},
		{
			name: "not nil",
			input: &Person{
				Name: "Tom",
			},
			wantVal: &Person{
				Name: "Tom",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val := NewValue[*Person]()
			val.Store(tc.input)
			v := val.Load()
			assert.Equal(t, tc.wantVal, v)
		})
	}
}

func TestValue_Swap(t *testing.T) {
	testCases := []struct {
		name string
		old  *Person
		new  *Person
	}{
		{
			name: "both nil",
		},
		{
			name: "old nil",
			old:  nil,
			new: &Person{
				Name: "Tom",
			},
		},
		{
			name: "new nil",
			old: &Person{
				Name: "Tom",
			},
			new: nil,
		},
		{
			name: "both not nil",
			old: &Person{
				Name: "Tom",
			},
			new: &Person{
				Name: "Jerry",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val := NewValueOf[*Person](tc.old)
			oldVal := val.Swap(tc.new)
			newVal := val.Load()
			assert.Equal(t, tc.old, oldVal)
			assert.Equal(t, tc.new, newVal)
		})
	}
}

func TestValue_CompareAndSwap(t *testing.T) {
	testCases := []struct {
		name string
		old  *Person
		new  *Person
	}{
		{
			name: "both nil",
		},
		{
			name: "old nil",
			old:  nil,
			new: &Person{
				Name: "Tom",
			},
		},
		{
			name: "new nil",
			old: &Person{
				Name: "Tom",
			},
			new: nil,
		},
		{
			name: "both not nil",
			old: &Person{
				Name: "Tom",
			},
			new: &Person{
				Name: "Jerry",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val := NewValueOf[*Person](tc.old)
			swapped := val.CompareAndSwap(tc.old, tc.new)
			assert.True(t, swapped)
		})
	}
}
