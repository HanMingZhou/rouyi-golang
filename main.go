package main

import rouyiconf "ruoyi/ruoyiFrame/conf"

func main() {
	configDefault := rouyiconf.NewConfigDefault()
	configDefault.LoadConf()
}
