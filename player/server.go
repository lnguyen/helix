package player

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/longnguyen11288/martini"
)

type Server struct {
	config        *Config
	martiniServer *martini.Martini
}

func New(c *Config) (*Server, error) {

	m := martini.New(log.New(ioutil.Discard, "", 0))
	r := martini.NewRouter()
	m.Map(c)
	m.Use(martini.Recovery())
	m.MapTo(r, (*martini.Routes)(nil))
	r.Get("/files", FilesHandler)
	r.Get("/status", StatusHandler)
	r.Post("/playfile", PlayFileHandler)
	r.Post("/stopfile", StopFileHandler)
	m.Action(r.Handle)

	s := &Server{
		config:        c,
		martiniServer: m,
	}
	return s, nil
}

func (s *Server) StartServer() {
	//TODO allow setting of host
	addr := s.config.Host + ":" + strconv.Itoa(s.config.Port)
	log.Printf("[INFO] Starting server on %v\n", addr)
	s.martiniServer.RunOnAddr(addr)
}
