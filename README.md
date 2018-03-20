# goGenerater
go的逆向工程

## 初衷
从java转go，习惯于mybatis和hibernate的逆向工程  
开发这个generater主要目的在于提高编码效率

## 使用详解：
### 1.配置./resources/goGenerater.ini配置文件
      ;数据库配置信息
      [DataSource]  
      DataSourceName=zct:123456@tcp(10.0.0.101:3306)/go_generater?charset=utf8  
      DriverName=mysql  
      DbName=go_generater  

      [ModelsMapping]  



      ;生成model的策略  
      [GenerationStrategy]  
      ;生成的模型的目录  默认为 ./models/  
      FilePath=./models/  

      ;是否生成在一个文件当中,true则生成在一个文件中，false则以表名为文件名生成，默认为true  
      OneFile=false   

      ;文件名 仅在OneFile=true的情况下有效,默认为models.go   
      FileName=models.go   

      ;go文件包名，默认为models   
      PackageName=models   

### 2.编写main方法
        package main

        import "goGenerater"

        func main() {
          goGenerater.Run()
        }
## 目前还存在的问题
### 1. 没有实现import
### 2. 目前仅实现了mysql数据库的支持
### 3. 数据类型映射不够完整
### 4. 目前没有实现字段名的高度定制化，仅仅实现了根据数据库的字段名生成struct的字段名，两者完全一致，  
###    对于习惯于自己的命名方式的农友可能不太友好，后期改善
    
