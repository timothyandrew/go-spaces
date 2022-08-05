package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

type SetCmd struct {
	Index uint `arg:"" help:"Index of current space."`
}

func (c *SetCmd) Run() error {
	return setCurrentSpace(c.Index)
}

type SwitchCmd struct {
	Index uint `arg:"" help:"Index of space to switch to."`
}

func (c *SwitchCmd) Run() error {
	current, err := currentSpace()
	if err != nil {
		return err
	}

	output := generateSwitch(current, c.Index)

	err = setCurrentSpace(c.Index)
	if err != nil {
		return err
	}
	fmt.Fprint(os.Stdout, output)

	return nil
}

type PrevCmd struct{}

func (c *PrevCmd) Run() error {
	current, err := currentSpace()
	if err != nil {
		return err
	}

	prev, err := prevSpace()
	if err != nil {
		return err
	}

	output := generateSwitch(current, prev)

	err = setCurrentSpace(prev)
	if err != nil {
		return err
	}
	fmt.Fprint(os.Stdout, output)

	return nil
}

var cli struct {
	Set    SetCmd    `cmd:"" help:"Set current space."`
	Switch SwitchCmd `cmd:"" help:"Switch to space by index."`
	Prev   PrevCmd   `cmd:"" help:"Switch to the space you were previously on."`
}

func main() {
	cmd := kong.Parse(&cli)
	err := cmd.Run()
	cmd.FatalIfErrorf(err)
}
