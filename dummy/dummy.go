package dummy

import (
	"github.com/gocaio/goca/dork"
	"github.com/gocaio/goca/plugin"
	"github.com/gocaio/goca/rsrc"
)

func init() {
	plg := plugin.NewPlugin("application/pdf")
	plg.Name = "Dummy plugin 1"
	plg.Description = "This plugin is as dumb as a stone, enjoy :)"
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
	output.MainType = "application/pdf"
	return
}

func myDorks() (dorks []*dork.Dork) {
	dorks = append(dorks, dork.NewDork(dork.Google, "filetype:pdf pdf"))
	dorks = append(dorks, dork.NewDork(dork.Bing, "filetype:pdf pdf"))
	return
}
