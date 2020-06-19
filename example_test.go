package sampling_test

import (
	"errors"
	"fmt"

	"kr.dev/sampling"
)

func Example_errors() {
	// Make & populate a new sequence of errors with capacity 5.
	seq := sampling.New(error)(5)
	seq.Add(errors.New("first"))
	seq.Add(errors.New("second"))
	seq.Add(errors.New("third"))

	// Read a sample into errs.
	errs := make([]error, seq.Cap())
	n := seq.Sample(errs)

	fmt.Println("our sample:", errs[:n])
	// Output:
	// our sample: [first second third]
}

func Example_many() {
	// Make & populate a new sequence of ints with capacity 3.
	seq := sampling.New(int)(3)
	for i := 0; i < 10_000; i++ {
		seq.Add(i)
	}

	// Read a sample into ints.
	ints := make([]int, seq.Cap())
	n := seq.Sample(ints)

	fmt.Println("our sample:", ints[:n])
	// Output:
	// our sample: [1 2 3]
}
