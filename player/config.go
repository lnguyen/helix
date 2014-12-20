package player

//Config for player
type Config struct {
	DataDir string `json:"data_dir"`
	Host    string `json:host`
	Name    string `json:"name"`
	Port    int    `json:"port"`
	Mock    bool   `json:mock`
}
