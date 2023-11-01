package config

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	LogDir     string `mapstructure:"log_dir" json:"log_dir" yaml:"log_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
