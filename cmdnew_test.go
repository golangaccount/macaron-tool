package main

import (
	"testing"
)

func TestNewcmd(t *testing.T) {
	t.Log("test .................")
	cmdNew.Run(cmdNew, []string{"testnewcmd"})
}