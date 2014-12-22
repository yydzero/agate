// Package cmd executes external command.
package cmd

import (
	"bytes"
	"errors"
	"os/exec"
)

func (c *exec.Cmd) Exec() ([]byte, []byte, error) {
	if c.Stdout != nil {
		return nil, nil, errors.New("exec: Stdout already set")
	}
	if c.Stderr != nil {
		return nil, nil, errors.New("exec: Stderr already set")
	}

	var sout bytes.Buffer
	var serr bytes.Buffer
	c.Stdout = &sout
	c.Stderr = &serr
	err := c.Run()
	return sout.Bytes(), serr.Bytes(), err
}
