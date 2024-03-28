package errorhandling

import (
	"fmt"
	"os"
)

func WriteFile() error {
	file, err := os.CreateTemp(os.TempDir(), "")
	if err != nil {
		return err
	}

	for i := 0; i < 100; i++ {
		res, err := FizzbuzzWithError(i)
		if err != nil {
			return err
		}

		file.WriteString(fmt.Sprintf("%d: %s", i, res))
	}

	return file.Close()
}
