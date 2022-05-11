package httpServer

type Endpoint struct {
	Path   string
	Method string
	Handle HandlerFunc
}

type GroupEndpoints struct {
	Prefix   string
}

