package player

//Config for player
type Config struct {
	DataDir string `json:"data_dir"`
	Name    string `json:"name"`
	Port    int    `json:"port"`
}
