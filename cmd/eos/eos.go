package eos

import (
	"fmt"
	"github.com/go-helios/eos/internal/creator"
	"github.com/go-helios/eos/internal/project"
	"github.com/go-helios/eos/internal/upgrade"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

var cmdRoot = &cobra.Command{
	Use:     "eos",
	Example: "eos new demo",
	Version: fmt.Sprintf("\n %s \n", version),
}

func init() {
	cmdRoot.AddCommand(project.CmdNew)
	cmdRoot.AddCommand(creator.CmdCreator)
	cmdRoot.AddCommand(upgrade.CmdUpgrade)

	creator.CmdCreator.AddCommand(creator.CmdCreatorHandler)
	creator.CmdCreator.AddCommand(creator.CmdCreatorService)
	creator.CmdCreator.AddCommand(creator.CmdCreatorRepository)
	creator.CmdCreator.AddCommand(creator.CmdCreatorModel)
	creator.CmdCreator.AddCommand(creator.CmdCreatorRequest)
	creator.CmdCreator.AddCommand(creator.CmdCreatorResponse)
	creator.CmdCreator.AddCommand(creator.CmdCreatorAll)

}

func Execute() error {
	return cmdRoot.Execute()
}
