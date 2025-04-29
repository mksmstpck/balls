package config

type Config struct {
	Radius int
}

func NewConfig(radius int) *Config {
	return &Config{
		Radius: radius,
	}
}
