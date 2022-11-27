package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecCommandline(command string) (err error, out bytes.Buffer, stderr bytes.Buffer) {
	fmt.Println(command)
	cmd := exec.Command("powershell", command)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	fmt.Println("Result: " + out.String())
	return
}
