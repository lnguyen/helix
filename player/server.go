package player

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/longnguyen11288/helix/player/omxplayer"
	"github.com/longnguyen11288/martini"
)

type Server struct {
	config        *Config
	martiniServer *martini.ClassicMartini
}

func New(c *Config) (*Server, error) {

	var player Player
	if c.Mock {
		player = NewMockPlayer()
	} else {
		player = omxplayer.NewOmxPlayer()
	}

	m := martini.Classic()
	m.Map(log.New(ioutil.Discard, "", 0))
	m.Map(c)
	m.Map(player)
	m.Get("/files", FilesHandler)
	m.Get("/status", StatusHandler)
	m.Post("/playfile", PlayFileHandler)
	m.Post("/stopfile", StopFileHandler)

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
