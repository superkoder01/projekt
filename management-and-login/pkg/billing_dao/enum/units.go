package enum

type Units string

const (
	kWh Units = "kWh"
	mWh Units = "mWh"
)

func (s Units) Name() string {
	switch s {
	case kWh:
		return "kWh"
	case mWh:
		return "mWh"
	default:
		return ""
	}
}
