package main

import (
	"fmt"
	"github.com/YukihiroTaniguchi/laun/internal/cmd"
	"os"

)



func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd := cmd.NewLaunCommand(wd)

	return cmd.Execute()
}