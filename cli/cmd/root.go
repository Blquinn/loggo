/*
Copyright © 2022 Aurelio Calegari, et al.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aurc/loggo/cli/loggo"
	"github.com/spf13/cobra"
)

const (
	cmdPrefix    = "loggo-"
	cmdGcpStream = "gcp-stream"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loggo",
	Short: "Stream json logs as rich TUI",
	Long: `l'oGGo provides a rich Terminal User Interface for streaming json based
logs and a toolset to assist you tailoring the display format.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Initiate adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Initiate() {
	loggo.BuildVersion = BuildVersion
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

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
		Use:   cmdGcpStream,
		Short: "Continuously stream GCP stack driver logs",
		Long: `Continuously stream Google Cloud Platform log entries
from a given selected project and GCP logging filters:

	loggo gcp-stream \ 
            --project myGCPProject123 \
            --from 1m \
            --filter 'resource.labels.namespace_name="awesome-sit" AND resource.labels.container_name="some"' 
`,
		DisableFlagParsing: true,
		Run:                runExecutableSubcommand(cmdGcpStream),
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.loggo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add any found exectuable subcommands to root cmd.
	for _, c := range getExecutableSubCommands() {
		if cmd, f := executableCommands[c]; f {
			rootCmd.AddCommand(cmd)
		} else {
			panic(fmt.Sprintf("No implementation found for command %s", c))
		}
	}
}
