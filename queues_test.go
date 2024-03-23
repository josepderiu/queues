package queues

import (
	"io"
	"os"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	if len(q.elements) != 1 {
		t.Errorf("Enqueue was incorrect, got: %d, want: %d.", len(q.elements), 1)
	}
}

func TestDequeue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	deq := q.Dequeue()
	if deq != 1 || len(q.elements) != 1 {
		t.Errorf("Dequeue was incorrect, got: %d, want: %d.", deq, 1)
	}
}

func TestPrintElements(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Redirect stdout to capture the output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	q.PrintElements()

	// Restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Read the captured output
	out, _ := io.ReadAll(r)

	expected := "1\n2\n3\n"
	if string(out) != expected {
		t.Errorf("PrintElements was incorrect, got: %s, want: %s.", string(out), expected)
	}
}
