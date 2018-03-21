/**
 * @Desc  逆向工程需要的model
 * @author zjhfyq
 * @data 2018/3/20 11:11.
 */
package goGenerater



//数据库连接字段
type DataSource struct {
	DataSourceName string
	DriverName     string
	DbName         string
}



type Field struct {
	FieldName  string      //字段名
	Type       string      //字段类型
	Collation  interface{} //字符集（mysql 5.0以上有）
	Null       string      //是否可以为NULL
	Key        interface{} //索引（PRI,unique,index)
	Default    interface{} //缺省值
	Extra      interface{} //额外（是否 auto_increment)
	Privileges string      //权限
	Comment    interface{} //备注（mysql 5.0以上有)
}

type GenerationStrategy struct {
	FilePath string
	OneFile  bool
	FileName string
	PackageName  string
}
