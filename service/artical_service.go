package service

import (
	_ "github.com/kataras/iris/mvc"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"github.com/kataras/iris"

)

//定义mysql中的Articals表
type Articals struct {
	Id int64
	Headline string
	Content string
}

type Userinfo struct {
	Id int64
	Name string
	Password string
}

//提供的service
type ArticalService interface {
	GetArticalByID(id int64,artical *Articals)(error)//查询服务
	GetAllArticals(articalslice *[]Articals)(error,int)//查询所有记录
	InsertArtical(articals *Articals)(error)//插入服务
	DeleteArtical(articals *Articals)(error)//删除服务
}

type UserService interface {
	GetUser(name string,userinfo *Userinfo)(error)
}

////实现ArticalService接口和UserinfoService接口
//type articalService struct {
//	engine *xorm.Engine //全局
//}

//实现ArticalService接口和UserinfoService接口
type articalService struct {
	engine *xorm.Engine //全局
}

//生成一个articalService，返回ArticalService接口
func NewArticalService() ArticalService{
	//创建mysql数据库引擎，指定数据库种类，登陆名及密码，数据库名，
	mysqlengine,_:=xorm.NewEngine("mysql","root:123@/test?charset=utf8")
	//设置数据库引擎
	mysqlengine.ShowSQL(true)//打印执行的SQL语句
	mysqlengine.SetMapper(core.SameMapper{})//表名映射方式

	iris.RegisterOnInterrupt(func(){
		mysqlengine.Close()
	})

	return &articalService{mysqlengine}
}

//service通过xorm操作mysql数据库
func (s *articalService) GetArticalByID(id int64,artical *Articals)(error) {
	//if ok,_:=Globalvaribal.engine.Get(&artical);ok{
		//return mvc.View{
		//	Name:"information.html",
		//	Data:iris.Map{
		//		//"Title":"mytitle",
		//		"Articals":artical,
		//	},
		//}
	//}
	_,err:=s.engine.Get(artical)
	return err

}

//查询所有记录
func (s *articalService) GetAllArticals(articalslice *[]Articals)(error,int){
	err:=s.engine.Find(articalslice)
	length:=len(*articalslice)
	return err,length
}

func (s *articalService) InsertArtical(articals *Articals)(error){
	_,err:=s.engine.Insert(articals)
	return err
}

func (s *articalService) DeleteArtical(articals *Articals)(error){
	_,err:=s.engine.Delete(articals)
	return err
}

//以下实现Userinfo接口
func (s *articalService) GetUser(name string,userinfo *Userinfo)(error){
	_,err:=s.engine.Get(userinfo)
	return err
}







