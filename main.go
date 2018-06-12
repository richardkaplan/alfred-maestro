package main

import (
	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func run() {
	macros, err := getKmMacros()
	if err != nil {
		wf.Fatal(err.Error())
		return
	}

	for _, macro := range macros {
		wf.NewItem(macro.Name).UID(macro.UID).Valid(true).Arg(macro.UID)
	}

	args := wf.Args()
	var searchQuery string
	if len(args) > 0 {
		searchQuery = args[0]
	}

	if searchQuery == "" {
		wf.WarnEmpty("No macros found", "It seems that you haven't created any macros yet.")
	} else {
		wf.Filter(searchQuery)
		wf.WarnEmpty("No macros found", "Try a different query.")
	}

	wf.SendFeedback()
}

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}
