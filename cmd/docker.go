package cmd

import (
	"os"
	"os/exec"
)

func composeCommand(project Project, devEnv string, arg []string) error {
	if devEnv != "" {
		arg = append([]string{"-f docker-compose." + devEnv + ".yml"}, arg...)
	}

	return dc(project.Path, arg...)
}

func dc(path string, arg ...string) error {
	return run(path, "docker-compose", arg...)
}

func stopContainers() error {
	return run(".", "docker stop $(docker ps -aq)")
}

func run(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
