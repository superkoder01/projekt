package rbac

type RBAC interface {
	IsAllowed(role string, endpoint string, httpMethod string) bool
	GetRules(tokenRole string) []Rule
	Omit(path string) bool
	GetApiPrefix() string
}

type Rbac struct {
	Roles         []Role   `yaml:"roles"`
	OmitEndpoints []string `yaml:"omitEndpoints"`
	ApiPrefix     string   `yaml:"apiPrefix"`
}

type Role struct {
	Name  string `yaml:"name"`
	Rules []Rule `yaml:"rules,omitempty"`
}

type Rule struct {
	Resource string   `yaml:"resource"`
	Verbs    []string `yaml:"verbs"`
}
