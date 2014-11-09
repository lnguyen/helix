package player

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreate(t *testing.T) {
	Convey("Test Create", t, func() {
		c := &Config{
			DataDir: "/tmp/foo",
			Name:    "TmpName",
			Port:    1234,
		}
		player, err := Create(c)
		So(err, ShouldBeNil)
		So(player, ShouldHaveSameTypeAs, &Player{})
	})
}
