package cmd

import (
	"errors"
	"fmt"

	"github.com/sahilm/fuzzy"
)

func match(projects []Project, pattern string) ([]Project, error) {
	dict := make(map[string]Project)
	for _, project := range projects {
		dict[project.Name] = project
	}

	list := make([]string, 0, len(projects))
	for _, project := range projects {
		list = append(list, project.Name)
	}

	matches := fuzzy.Find(pattern, list)

	if matches.Len() == 0 {
		return []Project{}, errors.New("No match found")
	}

	filteredProjects := make([]Project, 0, len(matches))
	for _, match := range matches {
		filteredProjects = append(filteredProjects, dict[match.Str])
	}

	return filteredProjects, nil
}

func search(pattern string) (Project, error) {
	matches, err := match(projects(), pattern)

	if err != nil {
		return Project{}, err
	}

	if len(matches) > 1 {
		fmt.Println("The following folders are candidates, please select one of:")
		for idx, match := range matches {
			fmt.Printf("%d. %s (%s)\n", idx+1, match.Name, blue(match.Path))
		}

		fmt.Println("Which item do you want to use ?")
		var choice int
		n, err := fmt.Scanf("%d\n", &choice)

		if err != nil || n != 1 {
			return Project{}, errors.New("Invalid choice selected")
		}

		if choice <= len(matches) && choice > 0 {
			return matches[choice-1], nil
		}

		return Project{}, errors.New("This is not an acceptable choice")
	}

	return matches[0], nil
}
