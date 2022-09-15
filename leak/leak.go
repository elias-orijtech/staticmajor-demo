package leak

import (
	"os"
)

func LeakFile() error {
	f, err := os.Open("somefile.txt")
	if err != nil {
		return err
	}
	_ = f
	// Leak f, thus triggering a leak warning.
	return nil
}

func FuzzyLeak() {
	_ = newResource()
}

type Resource struct {
}

func (r *Resource) Close() {
}

func newResource() *Resource {
	return new(Resource)
}
