package config

import ()

// Case 全局Setup config
var (
	Case Config
)

// Config Config
type Config struct {
	Signature Signature `mapstructure:"signature" json:"signature" yaml:"signature"`
}
