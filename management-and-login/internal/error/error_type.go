package error

type ErrorType int

const (
	DATABASE ErrorType = iota + 1
	SERVICE
	API
	BLOCKCHAIN
	CAST_TYPE
	UNKNOWN
)

func (e ErrorType) Name() string {
	switch e {
	case DATABASE:
		return "DATABASE"
	case SERVICE:
		return "SERVICE"
	case API:
		return "API"
	case BLOCKCHAIN:
		return "BLOCKCHAIN"
	case CAST_TYPE:
		return "CAST_TYPE"
	case UNKNOWN:
		return "UNKNOWN"
	default:
		return "UNKNOWN"
	}
}
