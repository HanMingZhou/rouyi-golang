package rouyiconf

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"os"
	ruoyilog "ruoyi/ruoyiFrame/log"
	ruoyiConv "ruoyi/ruoyiFrame/utils/conv"
	"ruoyi/ruoyiFrame/utils/ruoyiFile"
	"strings"
)

type ConfigDefault struct {
	vipperCfg   *viper.Viper
	proxyMap    map[string]string
	proxyEnable bool
}

func NewConfigDefault() *ConfigDefault {
	return &ConfigDefault{}
}
func (e *ConfigDefault) GetVipperCfg() *viper.Viper {
	return e.vipperCfg
}
func (e *ConfigDefault) GetValueStr(key string) string {
	if e.vipperCfg == nil {
		e.vipperCfg = e.LoadConf()
	}
	val := cast.ToString(e.vipperCfg.Get(key))
	if strings.HasPrefix(val, "$") { //存在动态表达式
		val = strings.TrimSpace(val)               //去空格
		val = ruoyiConv.SubStr(val, 2, len(val)-1) //去掉 ${}
		if strings.HasPrefix(val, "\"") {
			panic("${...} format error !!!")
		}
		index := strings.Index(val, ":") //ssz:按第一个: 分割，前半部分是占位符，后半部分是默认值
		val0 := ruoyiConv.SubStr(val, 0, index)
		val0 = os.Getenv(val0) //从环境变量中取值,替换
		if val0 == "" {        //未设置环境变量,使用默认值
			val = ruoyiConv.SubStr(val, index+1, len(val))
			val = strings.Trim(val, "\"")
		} else {
			val = val0
		}
	}
	return val
}

func (e *ConfigDefault) LoadConf() *viper.Viper {
	e.vipperCfg = viper.New()
	// ruoyiFile 读取配置文件 bootstrap.yml
	if ruoyiFile.IsFileExist("bootstrap.yml") || ruoyiFile.IsFileExist("bootstrap.yaml") {
		e.vipperCfg.SetConfigFile("bootstrap.yml")
		e.vipperCfg.SetConfigType("yaml")
		e.vipperCfg.AddConfigPath("./")
		e.vipperCfg.ReadInConfig()
	} else {
		ruoyilog.Warn("rouyiconf 未找到配置文件bootstrap.yml,将使用默认配置！！！")
	}

	// ruoyiFile 读取配置文件 application.yml
	if ruoyiFile.IsFileExist("application.yml") || ruoyiFile.IsFileExist("application.yaml") {
		e.vipperCfg.SetConfigName("application")
		e.vipperCfg.SetConfigType("yaml")
		e.vipperCfg.AddConfigPath("./")
		e.vipperCfg.MergeInConfig()
	} else {
		ruoyilog.Warn("未找到配置文件 application.yml 将使用默认配置！！！")
	}

	// ruoyiFile 读取配置文件的go.proxy.enable
	if e.vipperCfg.GetBool("go.proxy.enable") == true {
		e.proxyEnable = true
		e.GetProxyMap()
	} else {
		fmt.Println("!!！！！！！！！！！！！！!!! porxy feature is disabled ！！！！！！！！！！！！！！！！！！！！！！！")
		e.proxyEnable = false
	}

	return e.vipperCfg
}
func (e *ConfigDefault) GetProxyMap() *map[string]string {
	if e.proxyEnable && e.proxyMap == nil {
		e.LoadProxyInfo()
	}
	return &e.proxyMap
}
func (e *ConfigDefault) LoadProxyInfo() *map[string]string {
	fmt.Println("######### 加载代理配置信息 start #############")
	if !e.IsProxyEnable() {
		return nil
	}
	list := e.GetVipperCfg().GetStringSlice("go.proxy.prefix")
	e.proxyMap = make(map[string]string)
	for _, v := range list {
		index := strings.Index(v, "=")
		key := ruoyiConv.SubStr(v, 0, index)
		hostPort := ruoyiConv.SubStr(v, index+1, len(v))
		e.proxyMap[key] = hostPort
	}
	e.proxyEnable = e.GetBool("go.proxy.enable")
	fmt.Println("go.proxy:", e.proxyMap)
	fmt.Println("######### 加载代理配置信息 end #############")
	return &e.proxyMap
}
func (e *ConfigDefault) IsProxyEnable() bool {
	return e.proxyEnable
}
func (e *ConfigDefault) GetBool(key string) bool {
	if e.vipperCfg == nil {
		e.vipperCfg = e.LoadConf()
	}
	val := cast.ToString(e.vipperCfg.Get(key))
	val = e.parseVal(val)
	if val == "true" {
		return true
	} else {
		return false
	}
}
func (e *ConfigDefault) parseVal(val string) string {
	if strings.HasPrefix(val, "$") { //存在动态表达式
		val = strings.TrimSpace(val)               //去空格
		val = ruoyiConv.SubStr(val, 2, len(val)-1) //去掉 ${}
		index := strings.Index(val, ":")           //ssz:按第一个: 分割，前半部分是占位符，后半部分是默认值
		val0 := ruoyiConv.SubStr(val, 0, index)
		val0 = os.Getenv(val0) //从环境变量中取值,替换
		if val0 == "" {        //未设置环境变量,使用默认值
			val = ruoyiConv.SubStr(val, index, len(val))
		} else {
			val = val0
		}
	}
	return val
}
