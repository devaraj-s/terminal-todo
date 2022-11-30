package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

const (
	todoFile = "~/.todo"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
	DelegatedTo string
	DelegatedAt time.Time
	Delegated   bool
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}

	*t = append(*t, todo)
}

func (t *Todos) Delegate(index int, toPerson string) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].DelegatedAt = time.Now()
	ls[index-1].Delegated = true
	ls[index-1].DelegatedTo = toPerson

	return nil
}

func (t *Todos) Complete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true
	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load() error {
	file, err := ioutil.ReadFile(todoFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) List() error {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},

			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
			{Align: simpletable.AlignCenter, Text: "Delegated To"},
			{Align: simpletable.AlignCenter, Text: "Delegated At"},
		},
	}
	for index, row := range *t {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", index+1)},
			{Align: simpletable.AlignLeft, Text: getIcon(row)},
			{Align: simpletable.AlignLeft, Text: row.Task},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%t", row.Done)},
			{Align: simpletable.AlignLeft, Text: row.CreatedAt.Format(time.RFC822)},
			{Align: simpletable.AlignLeft, Text: row.CreatedAt.Format(time.RFC822)},
			{Align: simpletable.AlignLeft, Text: row.DelegatedTo},
			{Align: simpletable.AlignLeft, Text: row.DelegatedAt.Format(time.RFC822)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	// table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())

	return nil
}

func (t *Todos) Store() error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(todoFile, data, 0o644)
}

//""
//"ﱤ"
//"ﱣ"
//""
//""
//""

func getIcon(row item) string {
	doneOption := ""

	if row.Done {
		doneOption = "ﱣ"
	}
	if row.Delegated {
		doneOption = ""
	}

	return doneOption
}
