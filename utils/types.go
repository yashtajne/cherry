package utils

const (
	O_EXT       = ".o"
	C_EXT       = ".c"
	CPP_EXT     = ".cpp"
	TIME_FORMAT = "2006-01-02 15:04:05"
)

type ProjectConfig struct {
	Project  Project
	Build    Build
	Packages []Pkg
}

type Project struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	Version     string `toml:"version"`
}

type Build struct {
	OS         string `toml:"os"`
	Shell      string `toml:"shell"`
	Compiler   string `toml:"compiler"`
	IncludeDir string `toml:"includedir"`
	LibDir     string `toml:"libdir"`
}

type Pkg struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	URL         string `toml:"url"`
	Version     string `toml:"version"`
	Libs        string `toml:"libs"`
	Cflags      string `toml:"cflags"`
}
