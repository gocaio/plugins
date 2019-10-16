package docx

import (
	"github.com/gocaio/goca/dork"
	"github.com/gocaio/goca/plugin"
	"github.com/gocaio/goca/rsrc"
)

func init() {
	plg := plugin.NewPlugin("application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	plg.Name = "Microsoft Office Word"
	plg.Description = "This plugin analyzes Word 2007+ metadata"
	plg.Dorks = myDorks()
	plg.Check = Check
	plg.Run = Run
}

// Check is the mimetype filter function, it returns true if it suports the
// given file.
func Check(data []byte) bool { return true }

// Run starts the analyzer and returns an goca output
func Run(in []byte) (output *rsrc.Output) {
	output = new(rsrc.Output)
	output.MainType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	return
}

func myDorks() (dorks []*dork.Dork) {
	dorks = append(dorks, dork.NewDork(dork.Google, "filetype:docx docx"))
	//dorks = append(dorks, dork.NewDork(dork.Bing, "filetype:pdf pdf"))
	return
}
