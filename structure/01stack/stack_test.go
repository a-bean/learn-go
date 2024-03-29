package stack

import (
	"reflect"
	"testing"
)

func TestStackArray(t *testing.T) {
	t.Run(
		"Stack with Array", func(t *testing.T) {
			Push(2)
			Push(3)
			t.Run(
				"stack Push", func(t *testing.T) {
					if !reflect.DeepEqual([]any{3, 2}, stackArray) {
						t.Errorf("Stack Push is not work we expected %v but got %v", []any{3, 2}, stackArray)
					}
				},
			)

			pop := Pop()
			t.Run(
				"stack Pop", func(t *testing.T) {
					if pop != 3 && Length() == 2 {
						t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
					}
				},
			)
			Push(4)
			Push(5)
			t.Run(
				"stack Peak", func(t *testing.T) {
					if Peak() != 5 {
						t.Errorf("Stack Peak is not work we expected %v but got %v", 5, Peak())

					}
				},
			)

			t.Run(
				"stack Length", func(t *testing.T) {
					if Length() != 3 {
						t.Errorf("Stack Length is not work we expected %v but got %v", 3, Length())

					}
				},
			)

			t.Run(
				"stack IsEmpty", func(t *testing.T) {
					if IsEmpty() != false {
						t.Errorf("Stack IsEmpty is not work we expected %v but got %v", false, IsEmpty())
					}
					Pop()
					Pop()
					Pop()
					if IsEmpty() != true {
						t.Errorf("Stack Peak is not work we expected %v but got %v", true, IsEmpty())
					}
				},
			)

		},
	)
}

func TestStackLinkedList(t *testing.T) {
	var newStack ListStack
	newStack.Push(1)
	newStack.Push(2)

	t.Run(
		"Stack Push", func(t *testing.T) {
			result := newStack.Show()
			expected := []any{2, 1}
			for x := range result {
				if result[x] != expected[x] {
					t.Errorf("Stack Push is not work, got %v but expected %v", result, expected)
				}
			}
		},
	)

	t.Run(
		"Stack IsEmpty", func(t *testing.T) {
			if newStack.IsEmpty() {
				t.Error("Stack IsEmpty is returned true but expected false", newStack.IsEmpty())
			}
		},
	)

	t.Run(
		"Stack Length", func(t *testing.T) {
			if newStack.Length() != 2 {
				t.Error("Stack Length should be 2 but got", newStack.Length())
			}
		},
	)

	newStack.Pop()
	pop := newStack.Pop()

	t.Run(
		"Stack Pop", func(t *testing.T) {
			if pop != 1 {
				t.Error("Stack Pop should return 1 but is returned", pop)
			}
		},
	)

	newStack.Push(52)
	newStack.Push(23)
	newStack.Push(99)

	t.Run(
		"Stack Peak", func(t *testing.T) {
			if newStack.Peak() != 99 {
				t.Error("Stack Peak should return 99 but got ", newStack.Peak())
			}
		},
	)
}
