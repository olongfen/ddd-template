package service

import "ddd-template/internal/service/controller"

type Server struct {
	Http *controller.HTTPServer
}

func NewServer(h *controller.HTTPServer) *Server {
	return &Server{
		Http: h,
	}
}
