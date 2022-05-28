package config

import ()

// Signature 签名设置
type Signature struct {
	Salts map[string]interface{} `mapstructure:"salts" json:"salts" yaml:"salts"`
}
