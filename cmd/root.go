package cmd

import (
	"fmt"
	"os"
	"os/user"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var config Config
var composeSuffix string
var list bool
var stop bool
var find string

var rootCmd = &cobra.Command{
	Use: "cfa [flags] [project] [compose command]",
	Example: `cfa my_project up -d
cfa -u=dev my_project exec my_container sh
cfa -f=my_pro
cfa -s`,
	Short: "docker compose from anywhere",
	Long: `Manage your compose projects from anywhere

cfa allows you to use the same compose CLI you already know
but without the need to cd into your directories.
Just pass your project name as the first argument
and run your compose command on it.`,
	// Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if stop {
			err := stopContainers()
			if err != nil {
				fmt.Printf("Cannot stop containers. %s\n", err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}

		if find != "" || list {
			printList(find)
			os.Exit(0)
		}

		if len(args) < 2 {
			cmd.Help()
			os.Exit(0)
		}

		project, err := search(args[0])

		if err != nil {
			fmt.Printf("Error ! %s\n", err.Error())
			os.Exit(1)
		}

		cyan := color.New(color.FgCyan).SprintFunc()
		fmt.Fprintf(color.Output, "Using project %s...\n", cyan(project.Name))

		composeCommand(project, composeSuffix, args[1:])
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&list, "list", "l", false, "List all available projects")
	rootCmd.Flags().BoolVarP(&stop, "stop", "s", false, "Stop all running containers")
	rootCmd.Flags().StringVarP(&composeSuffix, "suffix", "u", "", "Use a suffixed compose file (ex: -u=dev will use the docker-compose.dev.yml file)")
	rootCmd.Flags().StringVarP(&find, "find", "f", "", "List projects corresponding to search")
	rootCmd.Flags().SetInterspersed(false)
}

func initConfig() {
	usr, _ := user.Current()

	config = Config{
		Blacklist: []string{
			usr.HomeDir + "/Library",
			usr.HomeDir + "/Applications",
			"/node_modules",
			"/vendor",
			"/target",
			".cache",
			".git",
			".github",
			".idea",
			".vscode",
			".terraform",
		},
		Root:  usr.HomeDir,
		Depth: 7,
	}

	depth, ok := os.LookupEnv("CFA_DEPTH")
	if ok {
		depth, err := strconv.Atoi(depth)
		if err == nil {
			config.Depth = depth
		}
	}

	root, ok := os.LookupEnv("CFA_ROOT")
	if ok {
		config.Root = root
	}
}

// Execute - Launch root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
