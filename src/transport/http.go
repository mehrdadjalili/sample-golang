package transport

type Echo interface {
	Start(port int) error
	Shutdown() error
}
