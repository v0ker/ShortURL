package config

type UrlConfig struct {
	Domain    string `mapstructure:"domain" json:"domain" yaml:"domain"`
	MinLength int32  `mapstructure:"min_length" json:"min_length" yaml:"min_length"`
}
