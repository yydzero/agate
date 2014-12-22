package cmd

import (
	"fmt"
	"os/exec"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCommandNotFind(t *testing.T) {
	cmd := exec.Command("notexistcommand")
	sout, serr, err := cmd.Exec()
	assert.Nil(t, err)
	fmt.Println(sout)
	fmt.Println(serr)
}
