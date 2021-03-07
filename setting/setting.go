package setting

//YamlSettings Yaml Setting
type YamlSettings struct {
	Database struct {
		Type    []string `yaml:"type"`
		Path    string   `yaml:"path"`
		Name    string   `yaml:"name"`
		Apply   bool     `yaml:"apply"`
		Setting struct {
			Path          string   `yaml:"path"`
			Name          string   `yaml:"name"`
			SingularTable bool     `yaml:"singularTable"`
			TablePrefix   string   `yaml:"tablePrefix"`
			LogLevel      []string `yaml:"logLevel"`
			SlowThreshold int      `yaml:"slowThreshold"`
			Apply         bool     `yaml:"apply"`
		} `yaml:"setting"`
	} `yaml:"database"`
	Service struct {
		From struct {
			Path   string   `yaml:"path"`
			Name   string   `yaml:"name"`
			Ignore []string `yaml:"ignore"`
		} `yaml:"from"`
		To struct {
			Path    string `yaml:"path"`
			Postfix string `yaml:"postfix"`
		} `yaml:"to"`
		Apply bool `yaml:"apply"`
	} `yaml:"service"`
	QueryTools struct {
		Path  string `yaml:"path"`
		Name  string `yaml:"name"`
		Apply bool   `yaml:"apply"`
	} `yaml:"queryTools"`
}
