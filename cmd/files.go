package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"
)

func projects() []Project {
	channel := make(chan Project)
	var wg sync.WaitGroup

	wg.Add(1)
	go scan(&wg, config.Root, config.Depth, channel)

	// Turn channel into slice.
	projects := []Project{}
	go func() {
		for project := range channel {
			projects = append(projects, project)
		}
	}()

	wg.Wait()

	return projects
}

func scan(wg *sync.WaitGroup, folder string, depth int, results chan Project) {
	defer wg.Done()

	// Get all files and subdirectories in this directory.
	files, _ := ioutil.ReadDir(folder)
	var directories []string

	for _, file := range files {
		path := folder + "/" + file.Name()

		// Add subdirectories to list of yet to be scanned directories.
		if file.IsDir() {
			// Check if folder is in blacklist.
			for _, blacklist := range config.Blacklist {
				if blacklist == path {
					continue
				}
			}

			directories = append(directories, path)
			continue
		}

		// Search for docker-compose.yml file.
		if !file.IsDir() && file.Name() == "docker-compose.yml" {
			results <- Project{
				Path: filepath.Dir(path),
				Name: strings.Trim(strings.Replace(filepath.Dir(path), config.Root, "", 1), "/"),
			}

			// No need to continue scan other subdirectories
			return
		}
	}

	// If no docker-compose.yml file was found, scan all subdirectories that we found.
	if depth > 1 {
		for _, folder := range directories {
			wg.Add(1)
			go scan(wg, folder, depth-1, results)
		}
	}

	return
}

func printList(pattern string) {
	ps := projects()
	if pattern != "" {
		matches, err := match(ps, pattern)

		if err != nil {
			fmt.Printf("Error ! %s\n", err.Error())
			return
		}

		ps = matches
	}

	fmt.Printf("%d project(s) found:\n", len(ps))
	cyan := color.New(color.FgCyan).SprintFunc()
	for _, project := range ps {
		fmt.Fprintf(color.Output,"* %s (%s)\n", project.Name, cyan(project.Path+"\\docker-compose.yml"))
	}
}
