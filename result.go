package envdef

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
)

type Result struct {
	InsertSlice   InsertSlice
	UpdateSlice   UpdateSlice
	DeleteSlice   DeleteSlice
	NoChangeSlice NoChangeSlice
}

func (r *Result) Write() error {
	result := append(r.InsertSlice, r.UpdateSlice...)
	result = append(result, r.DeleteSlice...)
	result = append(result, r.NoChangeSlice...)
	sort.Strings(result)
	b := []byte(strings.Join(result, "\n"))

	return ioutil.WriteFile(".env.new", b, os.ModePerm)
}

func (r *Result) Print() {
	r.InsertSlice.Print()
	r.UpdateSlice.Print()
	r.DeleteSlice.Print()
}

type InsertSlice []string

func (es InsertSlice) Print() {
	iconInsert := "+"

	for _, e := range es {
		color.Yellow(iconFormat(iconInsert, e))
	}
}

type UpdateSlice []string

func (es UpdateSlice) Print() {
	iconUpdate := "~"

	for _, e := range es {
		color.Cyan(iconFormat(iconUpdate, e))
	}
}

type DeleteSlice []string

func (es DeleteSlice) Print() {
	iconDelete := "-"

	for _, e := range es {
		color.Red(iconFormat(iconDelete, e))
	}
}

type NoChangeSlice []string

func (es NoChangeSlice) Print() {
	for _, e := range es {
		color.White(e)
	}
}

func iconFormat(icon, msg string) string {
	iconFormat := "%v %v"
	return fmt.Sprintf(iconFormat, icon, msg)
}
