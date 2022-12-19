package stack

import "testing"

func TestCopyingStack(t *testing.T) {
	stackA := New()

	stackA.Push("A")

	stackB := *stackA

	stackB.Push("B")

	if stackA.Size() == stackB.Size() {
		t.Errorf("stack size should be different")
	}
}

func TestPopAfterCopy(t *testing.T) {
	stackA := *New()

	stackA.Push("A")
	stackA.Push("B")

	stackB := stackA

	stackB.Pop()

	if stackA.Size() == stackB.Size() {
		t.Errorf("stack size should be different")
	}
}

func TestCopyingStackInArray(t *testing.T) {
	stackA := New()

	stackA.Push("A")

	arrayA := []Stack{*stackA}

	arrayB := append([]Stack(nil), arrayA...)

	arrayB[0].Push("B")

	if arrayA[0].Size() == arrayB[0].Size() {
		t.Errorf("stack size should be different")
	}
}
