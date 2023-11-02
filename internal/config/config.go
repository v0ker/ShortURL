package config

type Configuration struct {
	App      App       `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log       `mapstructure:"log" json:"log" yaml:"log"`
	Redis    Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
	Database Database  `mapstructure:"database" json:"database" yaml:"database"`
	Url      UrlConfig `mapstructure:"url" json:"url" yaml:"url"`
}
