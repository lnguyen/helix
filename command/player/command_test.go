package player

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestCommand_implements(t *testing.T) {
	var _ cli.Command = new(Command)
}
