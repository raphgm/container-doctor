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
	dockerChecks "github.com/raphgm/container-doctor/providers/docker/checks"
	"github.com/raphgm/container-doctor/providers/compose"
	composeChecks "github.com/raphgm/container-doctor/providers/compose/checks"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run all registered diagnostic checks",
	Long: `Run all diagnostic checks to determine the health of your environment.
This will query all providers (Docker, Kubernetes, etc.) and generate a report.`,
	Run: func(cmd *cobra.Command, args []string) {
		exec := executor.New()
		
		reg := engine.NewRegistry()
		reg.Register(docker.New(
			dockerChecks.NewInstalled(exec),
			dockerChecks.NewDaemon(exec),
			dockerChecks.NewContext(exec),
			dockerChecks.NewVersion(exec),
			dockerChecks.NewSocket(exec),
		))
		
		reg.Register(compose.New(
			composeChecks.NewInstalled(exec),
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
