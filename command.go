package main

import (
	"fmt"
)

// Command interface
type Command interface {
	Execute()
}

// Receiver
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	fmt.Println("Light is off")
}

// Concrete Command
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(l *Light) Command {
	return &LightOnCommand{light: l}
}

func (lc *LightOnCommand) Execute() {
	lc.light.On()
}

// Concrete Command
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(l *Light) Command {
	return &LightOffCommand{light: l}
}

func (lc *LightOffCommand) Execute() {
	lc.light.Off()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) SetCommand(cmd Command) {
	rc.command = cmd
}

func (rc *RemoteControl) PressButton() {
	rc.command.Execute()
}

func main() {
	light := &Light{}
	lightOn := NewLightOnCommand(light)
	lightOff := NewLightOffCommand(light)

	remote := &RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()
}
