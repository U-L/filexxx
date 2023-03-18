package config

type MyConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`

}
type Address struct {
	Addr []string `mapstructure:"addr" json:"addr"`
}