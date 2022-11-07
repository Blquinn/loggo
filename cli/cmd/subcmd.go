package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	cmdPrefix    = "loggo-"
	cmdGcpStream = "gcp-stream"
)

func getExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func getExecutableSubCommands() []string {

	files, err := ioutil.ReadDir(getExecutablePath())
	if err != nil {
		panic(err)
	}

	subcommands := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), cmdPrefix) {
			subcommands = append(subcommands, strings.TrimPrefix(file.Name(), cmdPrefix))
		}
	}

	return subcommands
}

func runExecutableSubcommand(cmdName string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executing subcommand %s\n", cmdName)
		p := filepath.Join(getExecutablePath(), cmdPrefix+cmdName)
		execCmd := exec.CommandContext(cmd.Context(), p, args...)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Stdin = os.Stdin
		if err := execCmd.Run(); err != nil {
			fmt.Printf("Failed to run subcommand: %s\n", err)
			os.Exit(3)
		}

		os.Exit(execCmd.ProcessState.ExitCode())
	}
}

var executableCommands = map[string]*cobra.Command{
	cmdGcpStream: {
		Use:                cmdGcpStream,
		Short:              "Continuously stream GCP stack driver logs",
		DisableFlagParsing: true,
		Run:                runExecutableSubcommand(cmdGcpStream),
	},
}

func init() {
	// Add any found exectuable subcommands to root cmd.
	for _, c := range getExecutableSubCommands() {
		if cmd, f := executableCommands[c]; f {
			rootCmd.AddCommand(cmd)
		} else {
			panic(fmt.Sprintf("No implementation found for command %s", c))
		}
	}
}
