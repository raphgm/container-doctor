/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/raphgm/container-doctor/internal/engine"
	"github.com/raphgm/container-doctor/internal/executor"
	"github.com/raphgm/container-doctor/internal/renderer"
	"github.com/raphgm/container-doctor/pkg/report"
	"github.com/raphgm/container-doctor/providers/docker"
	"github.com/raphgm/container-doctor/providers/docker/checks"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exec := executor.New()
		
		reg := engine.NewRegistry()
		reg.Register(docker.New(
			checks.NewInstalled(exec),
			checks.NewDaemon(exec),
			checks.NewContext(exec),
			checks.NewVersion(exec),
			checks.NewSocket(exec),
		))
		
		eng := engine.New(reg)
		rep := eng.Run(cmd.Context())
		
		term := renderer.NewTerminal()
		if err := term.Render(rep); err != nil {
			fmt.Printf("Error rendering report: %v\n", err)
			os.Exit(1)
		}
		
		for _, p := range rep.Providers {
			for _, r := range p.Results {
				if r.Status == report.Fail {
					os.Exit(1)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
