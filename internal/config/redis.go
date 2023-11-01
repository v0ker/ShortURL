package config

type Redis struct {
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
