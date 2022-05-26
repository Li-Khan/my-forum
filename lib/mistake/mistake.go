package mistake

import "fmt"

// * Wrap - wraps the error
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// * WrapIfErr - returns nil if there is no error
func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}

	return Wrap(msg, err)
}
