/**
 * @Desc 逆向工程
 * @author zjhfyq
 * @data 2018/3/20 15:14.
 */
package goGenerater

import (
	"database/sql"
	"os"
	"strings"
	"log"
)

func Run() {
	//获取配置文件指针
	cfg := GetConfigFilePointer()
	//加载数据库配置文件信息
	ds := LoadDatasourceConfig(cfg)
	//加载逆向工程文件生成策略
	gs := LoadGenerationStrategyConfig(cfg)
	//获取数据操作指针
	db := GetDB(ds)
	//获取数据库表名
	tableNames := GetTableNames(db, ds.DbName)
	//获取model和table名称的映射配置
	modelMappings :=LoadModelsMapping(cfg)



	if gs.OneFile {
		//生成在一个文件中
		file, err := os.OpenFile(gs.FilePath + string(os.PathSeparator) + gs.FileName,os.O_RDWR,0777)
		if err != nil {
			panic(err)
		}else {
			//延迟关闭文件
			defer file.Close()
			log.Println("output file===>",string(gs.FilePath + string(os.PathSeparator) + gs.FileName))
			GeneraterOneFile(file, gs, tableNames, ds, db,modelMappings)
		}
	} else {
		GeneraterMulFile(gs,tableNames,ds,db,modelMappings)
	}
}

func GeneraterMulFile( gs GenerationStrategy, tableNames []string, ds DataSource, db *sql.DB,modelMappings  map[string]string)  {
	//写入包名
	for _, tableName := range tableNames {
		log.Printf("output model===>%s",tableName)
		fields := GetFieldByTableName(db, ds.DbName, tableName)


		file ,err := os.Create(gs.FilePath+string(os.PathSeparator)+tableName+".go")
		if err != nil{
			panic(err)
		}else {
			defer file.Close()
			file.WriteString("package " + gs.PackageName + "\n\n\n")
			//导包
			WriteImport(fields,file)
			WriteFile(file, tableName,fields,modelMappings)
		}
	}
}

func GeneraterOneFile(file *os.File, gs GenerationStrategy, tableNames []string, ds DataSource, db *sql.DB,modelMappings  map[string]string) {
	//写入包名
	file.WriteString("package " + gs.PackageName + "\n\n\n")
	for _, tableName := range tableNames {
		log.Printf("output model===>%s",tableName)
		fields := GetFieldByTableName(db, ds.DbName, tableName)

		//导包
		WriteImport(fields,file)

		WriteFile(file, tableName,fields,modelMappings)
	}
}

//写入导包
func WriteImport(fields [] Field,file *os.File)  {
	for _,field :=range  fields {
		if strings.Contains(field.Type,"time") {
			file.WriteString("import \"time\" \n\n\n")
		}
	}
}

func WriteFile(file *os.File, tableName string,fields []Field,modelMappings map[string]string) {

	//如果存在映射关系
	if tName ,ok := modelMappings[tableName];ok{
		tableName = tName
	}

	file.WriteString("type "+ tableName+ " struct {\n")
	for _,field := range fields {
		log.Printf("output model===>%s's field ===>%s ",tableName,field.FieldName)
		_ , err :=file.WriteString("	"+field.FieldName+"	" + SelectTypeMysql(field.Type)+"\n")
		if err != nil {
			panic(err)
		}
	}
	file.WriteString("}\n")
}



func SelectTypeMysql(mysqlType string)(goType string){
	if strings.Contains(mysqlType,"varchar") {
		goType = "string"
	}else if strings.Contains(mysqlType,"int") {
		goType = "int"
	}else if strings.Contains(mysqlType,"bool") {
		goType = "bool"
	} else if strings.Contains(mysqlType,"datetime") {
		goType = "time.Time"
	} else if strings.Contains(mysqlType,"timestamp") {
		goType = "int64"
	}else {
		goType = "string"
	}
	return
}
