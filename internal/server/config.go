package server

type Config struct {
    Server struct {
        Port            string `yaml:"port"`
        MaxConnections  int    `yaml:"max_connections"`
        ShutdownTimeout string `yaml:"shutdown_timeout"`
    } `yaml:"server"`
    Logging struct {
        Level  string `yaml:"level"`
        Format string `yaml:"format"`
    } `yaml:"logging"`
}

func LoadConfig(path string) (*Config, error) {
    // ... загрузка конфигурации
}