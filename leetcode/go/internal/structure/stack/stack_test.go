package stack

import (
	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}

	t.Run("IsEmpty returns true", func(t *testing.T) {
		got := s.IsEmpty()
		want := true

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("Peek returns error", func(t *testing.T) {
		_, err := s.Peek()

		utilityGoTest.AssertNotNil(t, err)
	})

	t.Run("Push adds to stack", func(t *testing.T) {
		s.Push(1)

		got, err := s.Peek()
		want := 1

		utilityGoTest.AssertNil(t, err)
		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("Pop removes from stack", func(t *testing.T) {
		got, err := s.Pop()
		want := 1

		utilityGoTest.AssertNil(t, err)
		utilityGoTest.AssertEquality(t, got, want)

		_, err = s.Peek()

		utilityGoTest.AssertNotNil(t, err)
		utilityGoTest.AssertEquality(t, s.IsEmpty(), true)
	})
}
