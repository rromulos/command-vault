package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/rromulos/command-vault"
	"io"
	"os"
	"strings"
)

const (
	cmdFile = ".commands.json"
)

func main() {
	add := flag.Bool("add", false, "add a new command")
	del := flag.Int("d", 0, "delete a command")
	flag.Parse()

	// Initializing the struct
	commands := &command.Commands{}

	// tries to load the file that contains the commands
	if err := commands.Load(cmdFile); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		cmd, err := getInput(os.Stdin, flag.Args()...)
		args := strings.Split(cmd, ",")
		if args < 3 {
			fmt.Println("Error: missing arguments")
			os.Exit(1)			
		}
		commands.Add(args[0], args[1], args[2])
		err = commands.Save(cmdFile)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		err := commands.Delete(*del)
		if err != nil {
			fmt.Println(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = commands.Save(cmdFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
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