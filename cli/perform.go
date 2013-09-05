package cli

import (
	"flag"
	"fmt"
	"github.com/bruth/playbill/data"
	"os"
)

var performUsage = `
usage: playbill perform [<play>] [-file <path>] [-crew <path>]
`

func runPerform(c *Cmd, args []string) {
	var play *data.Play
	var crew *data.Crew
	var err error

	playFile := *c.Flags.Lookup("file").Value
	crewFile := *c.Flags.Lookup("crew").Value

	if playFile != "" {
		play, err = data.ImportPlay(playFile)
	} else if len(args) > 0 {
		// lookup play
	} else {
		c.Usage()
	}

	if crewFile != "" {
		crew, err = data.ImportCrew(crewFile)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	run, err := play.Perform(crew)
	fmt.Println(run)
}

var PerformCmd = &Cmd{
	Name:        "perform",
	Aliases:     []string{"run"},
	Short:       "Performs/executes a run",
	UsageString: performUsage,
	Flags:       flag.NewFlagSet("perform", flag.ExitOnError),
	Run:         runPerform,
}

func init() {
	PerformCmd.Flags.String("file", "", "Performs a one-off run for a Playbill file.")
	PerformCmd.Flags.String("crew", "", "Crew file")
}
