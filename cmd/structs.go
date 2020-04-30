package cmd

// Project Any folder that contains a `docker-compose.yml` file
type Project struct {
	Path string
	Name string
}

// Config CLI config
type Config struct {
	Root      string
	Blacklist []string
	Depth     int
}