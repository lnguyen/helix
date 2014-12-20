package player

import (
	"strings"
	"testing"

	"github.com/mitchellh/cli"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCommand_implements(t *testing.T) {
	var _ cli.Command = new(Command)
}

func TestPlayerRun(t *testing.T) {
	Convey("Test Player Run", t, func() {

		ui := new(cli.MockUi)
		c := &Command{Ui: ui}
		args := []string{"--help"}

		c.Run(args)

		if !strings.Contains(ui.OutputWriter.String(), "Usage: helix player") {
			t.Fatalf("bad: %#v", ui.OutputWriter.String())
		}
	})
}
