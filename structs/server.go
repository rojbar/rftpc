package structs

type Server struct {
	Domain string `mapstructure:"domain"`
	Port   string `mapstructure:"port"`
	Name   string `mapstructure:"name"`
}

type Config struct {
	KnownHosts map[string]Server `mapstructure:"knownhosts"`
}
