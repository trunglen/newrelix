package apii

import (
	"g/x/web"
	"io"
	"net/http"
)

type ApiServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewApiServer() *ApiServer {
	var s = &ApiServer{
		ServeMux: http.NewServeMux(),
	}

	s.HandleFunc("/setting", s.HandleSetting)
	s.HandleFunc("/config", s.HandleConfig)
	s.HandleFunc("/flag", s.HandleFlag)
	return s
}
func (s *ApiServer) HandleSetting(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}
func (s *ApiServer) HandleConfig(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}
func (s *ApiServer) HandleFlag(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}
