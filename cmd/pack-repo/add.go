package main

import (
	"fmt"
	"io"

	"github.com/Azure/draft/pkg/draft/draftpath"
	"github.com/spf13/cobra"

	"github.com/Azure/draft-pack-repo/repo"
	"github.com/Azure/draft-pack-repo/repo/installer"
)

type addCmd struct {
	out    io.Writer
	err    io.Writer
	source string
	home   draftpath.Home
}

func newAddCmd(out io.Writer) *cobra.Command {
	add := &addCmd{out: out}

	cmd := &cobra.Command{
		Use:   "add [flags] <path|url>",
		Short: "add a pack repository",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return add.complete(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return add.run()
		},
	}
	return cmd
}

func (a *addCmd) complete(args []string) error {
	if err := validateArgs(args, []string{"path|url"}); err != nil {
		return err
	}
	a.source = args[0]
	a.home = draftpath.Home(homePath())
	return nil
}

func (a *addCmd) run() error {
	ins, err := installer.New(a.source, "", a.home)
	if err != nil {
		return err
	}

	debug("installing pack repo from %s", a.source)
	if err := installer.Install(ins); err != nil {
		return err
	}

	var installedRepo repo.Repository
	repos := repo.FindRepositories(a.home.Packs())
	for i := range repos {
		if repos[i].Dir == ins.Path() {
			installedRepo = repos[i]
		}
	}

	fmt.Fprintf(a.out, "Installed pack repository %s\n", installedRepo.Name)
	return nil
}
