package repairModel

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Configuration 配置文件
type Configuration struct {
	ServerPort     string `json:"serverPort"`
	DatabaseName   string `json:"databaseName"`
	DataSourceName string `json:"dataSourceName"`
	Repair         struct {
		EmailAddr      string   `json:"emailAddr"`
		EmailPort      string   `json:"emailPort"`
		EmailPassword  string   `json:"emailPassword"`
		BsAuthUsername string   `json:"bsAuthUsername"`
		BsAuthPassword string   `json:"bsAuthPassword"`
		InfoEmailForm  string   `json:"infoEmailForm"`
		InfoEmailTitle []string `json:"infoEmailTitle"`
	} `json:"repair"`
}

// conf 配置文件单例对象
var conf Configuration

// GetConf 获取配置文件单例对象
func GetConf() Configuration {
	var once sync.Once

	// 初始化配置文件单例
	readConf := func() {
		// 读取配置文件
		fileData, err := os.ReadFile("conf.json")
		if err != nil {
			fmt.Println("无法打开配置文件:", err)
			return
		}

		// 解析配置文件
		if err := json.Unmarshal(fileData, &conf); err != nil {
			fmt.Println("无法解析配置文件:", err)
			return
		}
	}

	// 初始化过程仅运行一次
	once.Do(readConf)

	return conf
}
