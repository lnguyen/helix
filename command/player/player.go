package player

import (
	"fmt"

	"github.com/longnguyen11288/helix/player"
)

type Player struct {
	config *Config
	server *player.Server
}

func Create(config *Config) (*Player, error) {
	player := &Player{
		config: config,
	}
	err := player.setupServer()
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (p *Player) setupServer() error {
	server, err := player.New(p.playerConfig())
	if err != nil {
		return fmt.Errorf("Failed to start Player server: %v", err)
	}
	p.server = server
	return nil
}

func (p *Player) playerConfig() *player.Config {
	var base player.Config

	if p.config.DataDir != "" {
		base.DataDir = p.config.DataDir
	}

	if p.config.Host != "" {
		base.Host = p.config.Host
	}

	if p.config.Name != "" {
		base.Name = p.config.Name
	}

	if p.config.Port != 0 {
		base.Port = p.config.Port
	}

	if p.config.Mock != false {
		base.Mock = p.config.Mock
	}

	return &base
}
