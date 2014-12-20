package player

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMockPlayerInterface(t *testing.T) {
	var _ Player = new(MockPlayer)
}

func TestMockPlayer(t *testing.T) {
	Convey("Test Decoding Config", t, func() {
		player := NewMockPlayer()
		err := player.PlayFile("testfile.mkv")
		So(err, ShouldBeNil)
		So(player.FilePlaying(), ShouldEqual, "testfile.mkv")
		So(player.IsPlaying(), ShouldEqual, true)
		err = player.StopFile()
		So(err, ShouldBeNil)
		So(player.FilePlaying(), ShouldEqual, "")
		So(player.IsPlaying(), ShouldEqual, false)

	})

}
