package fs

import "os"

// FileExists returns true if the given file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// DirExists returns true if the given file exists and is a directory
func DirExists(dir string) bool {
	stat, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return stat.IsDir()
}