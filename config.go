/**
 * @Desc 读取相关配置文件
 * @author zjhfyq 
 * @data 2018/3/20 11:09.
 */
package goGenerater

import (
	"github.com/Unknwon/goconfig"
	"strings"
	"os"
	"log"
)




func GetConfigFilePointer() (cfg *goconfig.ConfigFile) {
	cfgp , err := goconfig.LoadConfigFile("./resources/goGenerater.ini")
	if err != nil {
		panic(err)
	}else {
		cfg = cfgp
	}
	return
}


//加载数据库连接配置信息
func LoadDatasourceConfig(cfg *goconfig.ConfigFile) (ds DataSource){
	dataSourceName , err := cfg.GetValue("DataSource","DataSourceName")
	if err != nil {
		panic(err)
	}else {
		driverName , err := cfg.GetValue("DataSource","DriverName")
		if err != nil {
			panic(err)
		}else {
			dbName ,err := cfg.GetValue("DataSource","DbName")
			if err != nil{
				panic(err)
			}else {
				ds.DataSourceName = dataSourceName
				ds.DriverName = driverName
				ds.DbName = dbName
			}
		}
	}
	return
}

//获取表名和结构体名的映射关系
func LoadModelsMapping(cfg *goconfig.ConfigFile)(modelMappings map[string]string)  {
	modelMappings , err :=cfg.GetSection("ModelsMapping")
	if err != nil {
		log.Println("not found config [ModelMapping], the struct will be named tablename")
	}
	return
}

//获取文件生成策略
func LoadGenerationStrategyConfig(cfg *goconfig.ConfigFile) (gs GenerationStrategy) {
	gs = GenerationStrategy{}
	//设置默认值
	gs.FileName = "models.go"
	gs.FilePath="./model/"
	gs.OneFile = true
	gs.PackageName = "models"

	//获取配置的值
	fileName ,_ := cfg.GetValue("GenerationStrategy","FileName")
	filePath ,_ :=cfg.GetValue("GenerationStrategy","FilePath")
	oneFile := true
	oneFileString ,_ :=cfg.GetValue("GenerationStrategy","OneFile")
	if strings.ToUpper(oneFileString) != "TRUE" {
		oneFile = false
	}
	packageName,_ :=cfg.GetValue("GenerationStrategy","PackageName")

	//验证文件路径的正确性
	err := os.MkdirAll(filePath,0777)
	if err != nil {
		panic(err)
	}else {
		if oneFile {
			log.Println("create file:",filePath+string(os.PathSeparator)+fileName)
			file , err :=os.Create(filePath+string(os.PathSeparator)+fileName)
			if err != nil {
				panic(err)
			}else {
				file.Close()
				gs.FilePath = filePath
				gs.FileName = fileName
				gs.OneFile = oneFile
				gs.PackageName= packageName
			}
		}else {
			gs.FilePath = filePath
			gs.FileName = fileName
			gs.OneFile = oneFile
			gs.PackageName= packageName
		}
	}
	return
}

