package cmd

import (
	"os"
	"os/exec"
)

func composeCommand(project Project, composeSuffix string, args []string) error {
	if composeSuffix != "" {
		args = append([]string{"-f", "docker-compose." + composeSuffix + ".yml"}, args...)
	}

	args = append([]string{"compose"}, args...)
	return run(project.Path, "docker", args...)
}

func stopContainers() error {
	return run(".", "sh", "-c", "docker stop $(docker ps -aq)")
}

func run(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
