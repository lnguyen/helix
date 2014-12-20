package player

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/go-martini/martini"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHandlers(t *testing.T) {
	Convey("Test Handlers", t, func() {

		//Setup temp files
		tmpDir, _ := ioutil.TempDir("", "handlers")
		tmpFile, _ := ioutil.TempFile(tmpDir, "handlers")

		c := Config{DataDir: tmpDir}

		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		m := martini.Classic()
		m.Map(NewMockPlayer())
		m.Map(c)
		m.Get("/files", FilesHandler)
		m.Get("/status", StatusHandler)
		m.Post("/playfile", PlayFileHandler)
		m.Post("/stopfile", StopFileHandler)
		mux.Handle("/", m)

		//Test files
		Convey("Temp files is listed", func() {
			var files Files
			resp, err := http.Get(server.URL + "/files")
			So(err, ShouldBeNil)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, &files)
			So(err, ShouldBeNil)
			So(files[0], ShouldEqual, path.Base(tmpFile.Name()))

			Convey("Play file", func() {
				file := File{path.Base(tmpFile.Name())}
				fileMarshal, err := json.Marshal(file)
				So(err, ShouldBeNil)
				resp, err := http.Post(server.URL+"/playfile", "", strings.NewReader(string(fileMarshal)))
				So(err, ShouldBeNil)
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, `{ "success": "true" }`)
				Convey("Get Status", func() {
					var status Status
					resp, err := http.Get(server.URL + "/status")
					So(err, ShouldBeNil)
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					So(err, ShouldBeNil)
					err = json.Unmarshal(body, &status)
					So(err, ShouldBeNil)
					So(status.Playing, ShouldEqual, true)
					So(status.Filename, ShouldEqual, path.Base(tmpFile.Name()))

				})
				Convey("Stop file", func() {
					resp, err := http.Post(server.URL+"/stopfile", "", strings.NewReader(""))
					So(err, ShouldBeNil)
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					So(err, ShouldBeNil)
					So(string(body), ShouldEqual, `{ "success": "true" }`)
				})
			})
		})

		//player := NewMockPlayer()
		//err = player.PlayFile("testfile.mkv")
		//So(err, ShouldBeNil)
		//So(player.FilePlaying(), ShouldEqual, "testfile.mkv")
		//So(player.IsPlaying(), ShouldEqual, true)
		//err = player.StopFile()
		//So(err, ShouldBeNil)
		//So(player.FilePlaying(), ShouldEqual, "")
		//So(player.IsPlaying(), ShouldEqual, false)

	})

}
