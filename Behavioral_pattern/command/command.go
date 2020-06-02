package command

import (
	"fmt"
	"github.com/pkg/errors"
)

// ========= Command ================
type Command interface {
	Execute()
}

type BaseCommand struct{}

func (cmd *BaseCommand) Execute() {
	fmt.Println("execute base command")
}

type CopyCommand struct {
	BaseCommand
}

func (cmd *CopyCommand) Execute() {
	fmt.Println("execute copy command")
}

func NewCopyCommand() *CopyCommand {
	return &CopyCommand{}
}

type PasteCommand struct {
	BaseCommand
	Receiver Receiver
}

func (cmd *PasteCommand) Execute() {
	fmt.Println("ready to execute paste command")
	fmt.Println("do some initialization before pasting")
	fmt.Println("send params to pasting Receiver")
	cmd.Receiver.Operate()
}

func NewPasteCommand() *PasteCommand {
	return &PasteCommand{Receiver: Receiver{}}
}

// ========== Receiver (optional)==========
type Receiver struct{}

func (rcv *Receiver) Operate() {
	fmt.Println("receiver operate")
}

// ============== Invoker ================
type Invoker struct {
	command map[string]Command
}

func (invoker *Invoker) SetCommand(shortCut string, command Command) error {
	if _, ok := invoker.command[shortCut]; !ok {
		invoker.command[shortCut] = command
	}
	return errors.Errorf("Conflict, short cut %s has been token", shortCut)
}

func (invoker *Invoker) Execute(shortCut string) {
	if invoker == nil || invoker.command == nil {
		return
	}
	command := invoker.command[shortCut]
	if command == nil {
		return
	}
	command.Execute()
}

func NewInvoker() *Invoker {
	return &Invoker{command: map[string]Command{}}
}

//type History struct {
//	command []Command
//}
//
//func (his *History) Push(command Command) {
//	if command == nil {
//		return
//	}
//	his.command = append(his.command, command)
//}
//
//func (his *History) Pop() Command {
//	if len(his.command ) <= 0 {
//		return nil
//	}
//	cmdNum := len(his.command)
//	command := his.command[cmdNum - 1]
//	his.command = his.command[:cmdNum - 1]
//	return command
//}
