package role

type Role int

const (
	SUPER_ADMIN Role = iota + 1
	ADMINISTRATOR_FULL
	ADMINISTRATOR_BASIC
	TRADER
	SUPER_AGENT
	AGENT
	PROSUMER
	NONE
)

func RoleName(name string) Role {
	switch name {
	case "SUPER_ADMIN":
		return SUPER_ADMIN
	case "ADMINISTRATOR_FULL":
		return ADMINISTRATOR_FULL
	case "ADMINISTRATOR_BASIC":
		return ADMINISTRATOR_BASIC
	case "TRADER":
		return TRADER
	case "SUPER_AGENT":
		return SUPER_AGENT
	case "AGENT":
		return AGENT
	case "PROSUMER":
		return PROSUMER
	default:
		return NONE
	}

}

func (r Role) Name() string {
	switch r {
	case SUPER_ADMIN:
		return "SUPER_ADMIN"
	case ADMINISTRATOR_FULL:
		return "ADMINISTRATOR_FULL"
	case ADMINISTRATOR_BASIC:
		return "ADMINISTRATOR_BASIC"
	case TRADER:
		return "TRADER"
	case SUPER_AGENT:
		return "SUPER_AGENT"
	case AGENT:
		return "AGENT"
	case PROSUMER:
		return "PROSUMER"
	default:
		return ""
	}

}
