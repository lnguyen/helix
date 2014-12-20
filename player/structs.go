package player

//File struct to output filename
type File struct {
	Filename string `json:"filename"`
}

//Status struct to output status
type Status struct {
	Playing  bool   `json:"playing"`
	Filename string `json:"filename"`
}

//Files List of files that can be played
type Files []string
