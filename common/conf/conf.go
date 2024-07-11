package conf

import rouyiconf "ruoyi/ruoyiFrame/conf"

var cfg *MyConfig

type MyConfig struct {
	rouyiconf.ConfigDefault
}

func int() {
	GetConfigInstance()
}
func GetConfigInstance() {

}
