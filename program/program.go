package lib

import (
	"syscall"
	"bytes"
	"path/filepath"
	"os/exec"
)

// Program holds information about a program
type Program struct {
	Name string
	Filename string
	Path string
}

// Find will try and find the given filename on the path
func Find(filename string, name string) *Program {
	path, err := exec.LookPath(filename)
	if err != nil {
		return nil
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return nil
	}
	return &Program{
		Name: name,
		Filename: filename,
		Path: path,
	}
}

// Run will execute the program with the given parameters
// Returns stdout + stderr as strings and an error if one occurred
func (p *Program) Run(vars ...string) (stdout, stderr string, exitCode int, err error) {
	cmd := exec.Command(p.Path, vars...)
	var stdo, stde bytes.Buffer
	cmd.Stdout = &stdo
	cmd.Stderr = &stde
	err = cmd.Run()
	stdout = string(stdo.Bytes())
	stderr = string(stde.Bytes())

	// https://stackoverflow.com/questions/10385551/get-exit-code-go
	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			exitCode = 1
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	return
}