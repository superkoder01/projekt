package config

type TemplateConfig struct {
	Name     string
	Template []Template
	//JSON OBJECTS + TEMPLATE PATH
}

type Template struct {
	Path     string
	Prefix   string
	Template []Template
}
