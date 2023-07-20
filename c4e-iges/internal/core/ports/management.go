package ports

type Status interface {
	IsAlive() (bool, error)
}

type StatusFactory interface {
	MakeService() Status
}
