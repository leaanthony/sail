package fs

import "os"

// FileExists returns true if the given file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
