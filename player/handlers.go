package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//PlayFileHandler to play file
func PlayFileHandler(p Player,
	w http.ResponseWriter, r *http.Request) {
	var file File
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &file)
	err = p.PlayFile(file.Filename)
	if err != nil {
		fmt.Fprint(w, `{ "error": "`+err.Error()+`" }`)
		return
	}
	fmt.Fprint(w, `{ "success": "true" }`)
}

//StopFileHandler is handler to stop playing file
func StopFileHandler(p Player, w http.ResponseWriter) {
	err := p.StopFile()
	if err != nil {
		fmt.Fprint(w, `{ "error": "`+err.Error()+`" }`)
		return
	}
	fmt.Fprint(w, `{ "success": "true" }`)
}

//FilesHandler list the files that can be played
func FilesHandler(c Config, w http.ResponseWriter) {
	var files Files
	osFiles, _ := ioutil.ReadDir(c.DataDir)
	for _, f := range osFiles {
		files = append(files, f.Name())
	}
	output, err := json.Marshal(files)
	if err != nil {
		log.Printf("[ERROR] error unmarshalling files: %v", err)
	}
	fmt.Fprint(w, string(output))
}

//ChannelHandler handle the channel player is listed as
//TODO
func ChannelHandler() string {
	return fmt.Sprintf(`{ "channel": "%s" }`, "todo")
}

//StatusHandler handles status of player
func StatusHandler(p Player, w http.ResponseWriter) {
	var status Status
	status.Playing = p.IsPlaying()
	status.Filename = p.FilePlaying()
	output, err := json.Marshal(status)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(output))
}
