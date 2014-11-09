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
	m.Use(martini.Recovery())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	s := &Server{
		config:        c,
		martiniServer: m,
	}
	return s, nil
}

func (s *Server) StartServer() {
	//TODO allow setting of host
	addr := "127.0.0.1:" + strconv.Itoa(s.config.Port)
	log.Printf("[INFO] Starting server on %v\n", addr)
	s.martiniServer.RunOnAddr(addr)
}
