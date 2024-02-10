package transport

type (
	ClientGrpc interface {
		StartClient(port int) error
	}
	ManagerGrpc interface {
		StartManager(port int) error
	}
	PageGrpc interface {
		StartPage(port int) error
	}
	DeveloperGrpc interface {
		StartDeveloper(port int) error
	}
)
