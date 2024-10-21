package config

// Config 结构体用于映射 YAML 文件内容，From config.go
type Config struct {
	AppName  string `yaml:"app_name"`
	Version  string `yaml:"version"`
	Debug    bool   `yaml:"debug"`
	Database struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		Db_type     string `yaml:"db_type"`
		Db_name     string `yaml:"db_name"`
		Sqlite_file string `yaml:"sqlite_file"`
	} `yaml:"database"`
}

type AppImage struct {
	ID   int
	Type string
	Path string
	Name string
}
