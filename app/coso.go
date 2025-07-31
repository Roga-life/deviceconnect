package main

import (
	"net"
	"net/http"
)

type Server struct{
	port int
	listener net.Listener
	sm *http.ServeMux
}


func NewServer(port int, sm *http.ServeMux) *Server {
	sm.HandleFunc("/",mifunc) 
	return &Server{
		port: port,
		sm: sm,
	}
}


func (s *Server) Start() error {
listener, err:= net.Listen("tcp", ":8080")
s.listener = listener
return err
}
func mifunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}	

func miMiddleware(user string)http.HandlerFunc{

	if user == "" {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))}
}


func main() {

	myMuxer:= http.NewServeMux()
	server:= NewServer(8080, myMuxer)
	server.Start()
}




