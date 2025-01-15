package ports

type Listener interface {
	Listen() error
}
