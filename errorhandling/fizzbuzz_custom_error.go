package errorhandling

func FizzbuzzWithCustomError(input int) (string, error) {
	return "", nil
}

func MyFizzbuzzProcessWithCustomError(input int) error {
	// Call FizzbuzzWithCustomError
	// If error is an InvalidInput error, return nil
	// Else return the error
	// Otherwise print the result
	return nil
}
