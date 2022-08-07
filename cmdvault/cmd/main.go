package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rromulos/command-vault"
)

const (
	cmdFile = "data/commands.json"
)

func main() {
	add := flag.Bool("a", false, "add a new command")
	del := flag.Int("d", 0, "delete a command")
	list := flag.Bool("l", false, "list all commands")
	searchCommand := flag.Bool("scom", false, "search for command")
	searchCategory := flag.Bool("scat", false, "search for category")
	searchDescription := flag.Bool("sdes", false, "search for Description")
	flag.Parse()

	commands := &command.Commands{}

	if err := commands.Load(cmdFile); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		readInstructionFromTerminal(commands)
	case *del > 0:
		deleteInstruction(commands, *del)
	case *list:
		commands.Print()
	case *searchCommand:
		doSearch("Instruction")
	case *searchCategory:
		doSearch("Category")
	case *searchDescription:
		doSearch("Description")
	default:
		fmt.Println(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty command is not allowed")
	}

	return text, nil

}

func doSearch(searchBy string) {
	commands := &command.Commands{}
	cmd, err := getInput(os.Stdin, flag.Args()...)
	cmd = strings.TrimPrefix(cmd, "=")

	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	commands.Search(searchBy, cmd)
}

func readInstructionFromTerminal(cmd *command.Commands) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("[Enter the command] => ")
	scanner.Scan()
	iInstruction := scanner.Text()
	fmt.Print("[Enter the category] => ")
	scanner.Scan()
	iCategory := scanner.Text()
	fmt.Print("[Enter the description] => ")
	scanner.Scan()
	iDescription := scanner.Text()
	cmd.Add(iInstruction, iCategory, iDescription)
	err := cmd.Save(cmdFile)

	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func deleteInstruction(cmd *command.Commands, del int) {
	err := cmd.Delete(del)

	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = cmd.Save(cmdFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
