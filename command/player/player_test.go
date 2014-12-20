package player

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	Convey("Test Create", t, func() {
		c := &Config{
			DataDir: "/tmp/foo",
			Host:    "127.0.0.1",
			Name:    "TmpName",
			Port:    1234,
		}
		player, err := Create(c)
		So(err, ShouldBeNil)
		So(player, ShouldHaveSameTypeAs, &Player{})
	})
}
