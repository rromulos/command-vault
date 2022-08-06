package command

import (
	"encoding/json"
	"errors"
	"fmt"

	// "fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type command struct {
	Instruction string
	Category    string
	Description string
	CreatedAt   time.Time
}

type Commands []command

func (c *Commands) Add(instruction string, category string, description string) {
	cmd := command{
		Instruction: instruction,
		Category:    category,
		Description: description,
		CreatedAt:   time.Now(),
	}

	*c = append(*c, cmd)
}

func (c *Commands) Delete(index int) error {
	ls := *c
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*c = append(ls[:index-1], ls[index:]...)

	return nil
}

func (c *Commands) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, c)

	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Save(filename string) error {
	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func (c *Commands) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Command"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Description"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *c {
		idx++
		instruction := cyan(item.Instruction)
		category := magenta(item.Category)
		description := yellow(item.Description)
		createdAt := green(item.CreatedAt.Format(time.RFC822))
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: instruction},
			{Text: category},
			{Text: description},
			{Text: createdAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
