package cmd

func blue(str string) string {
	return "\033[1;36m" + str + "\033[0m"
}
