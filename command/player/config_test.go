package player

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecodeConfig(t *testing.T) {
	Convey("Test Decoding Config", t, func() {
		input := `{"data_dir": "/tmp/", "name": "TestName", "host": "127.0.0.1", "port": 8312}`
		config, err := DecodeConfig(bytes.NewReader([]byte(input)))
		So(err, ShouldBeNil)
		So(config.DataDir, ShouldEqual, "/tmp/")
		So(config.Host, ShouldEqual, "127.0.0.1")
		So(config.Name, ShouldEqual, "TestName")
		So(config.Port, ShouldEqual, 8312)
	})

}

func TestDefaultConfig(t *testing.T) {
	Convey("Test Default Config", t, func() {
		config := DefaultConfig()
		So(config.DataDir, ShouldEqual, ".")
		So(config.Host, ShouldEqual, "0.0.0.0")
		So(config.Name, ShouldEqual, "Player-1")
		So(config.Port, ShouldEqual, 8821)
	})
}

func TestMergeConfig(t *testing.T) {
	Convey("Test bad path for reading config", t, func() {
		a := &Config{
			DataDir: "/tmp/foo",
			Name:    "TmpName",
			Port:    1234,
		}

		b := &Config{
			DataDir: "/real/dir",
			Host:    "0.0.0.0",
			Name:    "RealName",
			Port:    4567,
		}

		c := MergeConfig(a, b)
		So(reflect.DeepEqual(c, b), ShouldBeTrue)
	})
}

func TestReadConfigPaths(t *testing.T) {
	Convey("Test bad path for reading config", t, func() {
		_, err := ReadConfig("/this/is/great/fake/path")
		So(err, ShouldNotBeNil)
	})

	Convey("Test reading config", t, func() {
		tf, err := ioutil.TempFile("", "player")
		So(err, ShouldBeNil)
		tf.Write([]byte(`{"data_dir": "/tmp/", "name": "TestName", "host": "127.0.0.1", "port": 8312}`))
		tf.Close()
		defer os.Remove(tf.Name())

		config, err := ReadConfig(tf.Name())
		So(err, ShouldBeNil)
		So(config.DataDir, ShouldEqual, "/tmp/")
		So(config.Host, ShouldEqual, "127.0.0.1")
		So(config.Name, ShouldEqual, "TestName")
		So(config.Port, ShouldEqual, 8312)
	})
}
