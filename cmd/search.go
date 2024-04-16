package cmd

import (
	"errors"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
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
		return []Project{}, errors.New("no match found")
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
		cyan := color.New(color.FgCyan).SprintFunc()
		items := make([]string, len(matches))
		for idx, match := range matches {
			items[idx] = match.Name + " " + cyan(match.Path)
		}
		prompt := promptui.Select{
			Label: "The following projects are candidates",
			Items: items,
		}

		choice, _, err := prompt.Run()

		if err != nil {
			return Project{}, errors.New("invalid choice selected")
		}

		return matches[choice], nil
	}

	return matches[0], nil
}
