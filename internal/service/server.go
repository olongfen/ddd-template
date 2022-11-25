package service

import "ddd-template/internal/ports/controller"

type Server struct {
	Http *controller.HttpServer
}

func NewServer(h *controller.HttpServer) (*Server, func()) {
	return &Server{
			Http: h,
		}, func() {
			h.Close()
		}
}
