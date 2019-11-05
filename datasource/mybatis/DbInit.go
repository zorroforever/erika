package mybatis

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhuxiujia/GoMybatis"
	"io/ioutil"
	"iris/datasource/mybatis/mapper"
)

var userDbEngine GoMybatis.GoMybatisEngine
func init() {
	if MysqlUri == "*" {
		println("GoMybatisEngine not init! because MysqlUri is * or MysqlUri is ''")
		return
	}
	userDbEngine = GoMybatis.GoMybatisEngine{}.New()

	//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:123456@(***.mysql.rds.aliyuncs.com:3306)/test
	_, err := userDbEngine.Open("mysql", MysqlUri) //此处请按格式填写你的mysql链接，这里用*号代替
	if err != nil {
		panic(err.Error())
	}

	//动态数据源路由(可选)
	/**
	GoMybatis.Open("mysql", MysqlUri)//添加第二个mysql数据库,请把MysqlUri改成你的第二个数据源链接
	var router = GoMybatis.GoMybatisDataSourceRouter{}.New(func(mapperName string) *string {
		//根据包名路由指向数据源
		if strings.Contains(mapperName, "example.") {
			var url = MysqlUri//第二个mysql数据库,请把MysqlUri改成你的第二个数据源链接
			fmt.Println(url)
			return &url
		}
		return nil
	})
	engine.SetDataSourceRouter(&router)
	**/

	//自定义日志实现(可选)
	/**
		engine.SetLogEnable(true)
		engine.SetLog(&GoMybatis.LogStandard{
			PrintlnFunc: func(messages []byte) {
			},
		})
	    **/
	//读取mapper xml文件
	utes, err := ioutil.ReadFile("/Users/haruka/dev/go/src/iris/mybatis/user/UserMapper.xml")
	if err != nil {
		panic(err.Error())
	}
	//设置对应的mapper xml文件
	userDbEngine.WriteMapperPtr(&mybatis.UserMapperImpl, utes)

}



