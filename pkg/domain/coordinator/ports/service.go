package ports

type Coordinator interface {
	Serve() error
}
