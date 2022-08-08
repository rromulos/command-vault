package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/rromulos/command-vault"
)

const (
	cmdFile = "data/commands.json"
)

const (
	sequenceIdFile = "data/sequence.dat"
)

const (
	red     = "\033[0;91m%s\033[0m"
	cyan    = "\033[0;96m%s\033[0m"
	magenta = "\033[0;95m%s\033[0m"
	yellow  = "\033[0;93m%s\033[0m"
	green   = "\033[0;92m%s\033[0m"
	gray    = "\033[0;90m%s\033[0m"

	Underlined = "\033[4m%s\033[0m"
	Blink      = "\033[5m%s\033[0m"
	Bold       = "\033[1m%s\033[0m"
	Dim        = "\033[2m%s\033[0m"
	Reverse    = "\033[7m%s\033[0m"
	Hidden     = "\033[8m%s\033[0m"
)

func main() {
	add := flag.Bool("a", false, "add a new command")
	del := flag.Int("d", 0, "delete a command")
	list := flag.Bool("l", false, "list all commands")
	version := flag.Bool("v", false, "Shows the application version")
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
		cleanTerminal()
		readInstructionFromTerminal(commands)
	case *del > 0:
		deleteInstruction(commands, *del)
		cleanTerminal()
		commands.Print()
	case *list:
		cleanTerminal()
		commands.Print()
	case *version:
		getVersion()
	case *searchCommand:
		cleanTerminal()
		doSearch("Instruction")
	case *searchCategory:
		cleanTerminal()
		doSearch("Category")
	case *searchDescription:
		cleanTerminal()
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
	fmt.Printf(cyan, "[Enter the command]* => ")
	scanner.Scan()
	iInstruction := scanner.Text()
	fmt.Printf(magenta, "[Enter the category]* => ")
	scanner.Scan()
	iCategory := scanner.Text()
	fmt.Printf(yellow, "[Enter the description]* => ")
	scanner.Scan()
	iDescription := scanner.Text()
	validationResult := doValidation(iInstruction, iCategory, iDescription)

	idSequence := cmd.GenerateSequence()

	if validationResult == true {
		cmd.Add(idSequence, iInstruction, iCategory, iDescription)
		err := cmd.Save(cmdFile)

		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}

func deleteInstruction(cmd *command.Commands, del int) {
	idPosition := cmd.FindIdPosition(del)

	if idPosition == -1 {
		fmt.Println("ID not found")
		os.Exit(1)
	}

	err := cmd.Delete(idPosition)

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

func doValidation(instruction string, category string, description string) bool {

	if len(instruction) == 0 {
		fmt.Println("[Validation Failed] => type the COMMAND")
		return false
	}

	if len(category) == 0 {
		fmt.Println("[Validation Failed] => type the CATEGORY")
		return false
	}

	if len(description) == 0 {
		fmt.Println("[Validation Failed] => type the DESCRIPTION")
		return false
	}

	return true
}

func cleanTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getVersion() {
	fmt.Printf(Underlined, "Version => 1.1.2 \n")
	fmt.Printf(cyan, "Author  => Rômulo Santos \n")
	fmt.Printf(gray, "E-mail  => @rromulosp@gmail.com \n")
	fmt.Printf(gray, "Github  => @rromulos \n")
}
