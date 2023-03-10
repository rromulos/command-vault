package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/atotto/clipboard"
	"github.com/thedevsaddam/gojsonq/v2"
)

type command struct {
	Id          string
	Instruction string
	Category    string
	Description string
	CreatedAt   time.Time
}

type Commands []command

func (c *Commands) Add(id string, instruction string, category string, description string) {
	cmd := command{
		Id:          id,
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

func (c *Commands) GenerateSequence() string {
	content, err := ioutil.ReadFile("data/sequence.dat")

	if err != nil {
		err = ioutil.WriteFile("data/sequence.dat", []byte("0"), 0644)
	}

	sequence, err := strconv.ParseInt(string(content[:]), 10, 64)
	sequence++
	sequenceString := strconv.FormatInt(sequence, 10)

	err = ioutil.WriteFile("data/sequence.dat", []byte(sequenceString), 0644)

	return sequenceString
}

func (c *Commands) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Command"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Description"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *c {
		idx++
		id := gray(item.Id)
		instruction := cyan(item.Instruction)
		category := magenta(item.Category)
		description := yellow(item.Description)

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: id},
			{Text: instruction},
			{Text: category},
			{Text: description},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (c *Commands) FindIdPosition(id int) int {
	for idx, item := range *c {
		idx++

		if strconv.Itoa(id) == item.Id {
			return idx
		}
	}
	return -1
}

func (c *Commands) Search(kind string, value string) {

	var jq *gojsonq.JSONQ

	if kind == "Id" {
		jq = gojsonq.New().File("data/commands.json").Select("Id", "Instruction", "Category", "Description", "CreatedAt").Where(kind, "=", value)
	} else {
		jq = gojsonq.New().File("data/commands.json").Select("Id", "Instruction", "Category", "Description", "CreatedAt").Where(kind, "contains", value)
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Id"},
			{Align: simpletable.AlignCenter, Text: "Command"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Description"},
		},
	}

	var cells [][]*simpletable.Cell

	if x, ok := jq.Get().([]interface{}); ok {
		for i, e := range x {
			item := e.(map[string]interface{})

			i++
			id := gray(fmt.Sprintf("%v", item["Id"]))
			instruction := cyan(fmt.Sprintf("%v", item["Instruction"]))
			category := magenta(fmt.Sprintf("%v", item["Category"]))
			description := yellow(fmt.Sprintf("%v", item["Description"]))
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: id},
				{Text: instruction},
				{Text: category},
				{Text: description},
			})
		}
		table.Body = &simpletable.Body{Cells: cells}

		table.SetStyle(simpletable.StyleUnicode)

		table.Println()
	}
}

func (c *Commands) CopyToClipboard(id int) {
	jq := gojsonq.New().File("data/commands.json").Select("Instruction").Where("Id", "=", strconv.Itoa(id))

	if x, ok := jq.Get().([]interface{}); ok {
		for _, e := range x {
			item := e.(map[string]interface{})
			instruction := fmt.Sprintf("%v", item["Instruction"])
			clipboard.WriteAll(instruction)
			clipboard.ReadAll()
			fmt.Println("Command has been copied !")
		}
	}
}
