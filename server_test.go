package tbot

import "testing"

const (
	TestToken    = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"
	InvalidToken = "invalid"
)

func TestNewServerSuccess(t *testing.T) {
	server, err := NewServer(TestToken)
	if err != nil {
		t.Error(err)
	}
	if server == nil {
		t.Error("Server is nil")
	}
	if server.mux == nil {
		t.Error("Server mux is nil")
	}
}

func TestNewServerFail(t *testing.T) {
	server, err := NewServer(InvalidToken)
	if err == nil {
		t.Error("Invalid token should return error")
	}
	if server != nil {
		t.Error("Invalid token should return nil server")
	}
}

func TestAddMiddleware(t *testing.T) {
	server := &Server{}
	if len(server.middlewares) > 0 {
		t.Error("Middleware list should be empty by default")
	}
	server.AddMiddleware(func(HandlerFunction) HandlerFunction { return nil })
	if len(server.middlewares) != 1 {
		t.Error("AddMiddleware should add new middleware")
	}
}
