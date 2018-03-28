package main
import (
	//"time"

	"github.com/kataras/iris"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/xorm"
	//"github.com/go-xorm/core"
	"github.com/kataras/iris/mvc"
	"demo/service"
	"demo/controller"
	//"path"
)

////定义mysql中的Articals表,放在了service中
//type Articals struct {
//	Id int64
//	Headline string
//	Content string
//}

//定义封装的全局变量
type GlobalVaribal struct {
	//engine *xorm.Engine //放到了service中
	err error
	app *iris.Application //注意这个Application是iris的，不是mvc的
}
var Globalvaribal GlobalVaribal

//func main(){
//	//app:=iris.New()
//	app=iris.New()
//
//	//app.RegisterView(iris.HTML("./view", ".html"))
//	app.RegisterView(iris.HTML("./view",".html"))
//	mvc.Configure(app.Party("/rout"),myMVC)
//
//	engine,err=xorm.NewEngine("mysql","root:123@/test?charset=utf8")
//	//engine,err:=xorm.NewEngine("mysql","root:123@/test?charset=utf8")
//	engine.ShowSQL(true)
//	engine.SetMapper(core.SameMapper{})
//
//	if err!=nil{
//		app.Logger().Fatal("orm failed to initialized: %v",err)
//	}
//
//	iris.RegisterOnInterrupt(func(){
//		engine.Close()
//	})
//
//	app.Run(iris.Addr(":8080"))
//}

func main(){
	//创建iris.Application
	Globalvaribal.app=iris.New()

	//导入html模版
	Globalvaribal.app.RegisterView(iris.HTML("./view",".html"))

	//*****************************************************************************************
	////配置mvc（创建mvc.Application）
	//mvc.Configure(Globalvaribal.app.Party("/rout"),myMVC)
	///////另一种方式
	artcal:=mvc.New(Globalvaribal.app.Party("/rout"))
	articalService:=service.NewArticalService()
	//注册mvc的service接口
	artcal.Register(articalService)
	//配置mvc的Handle，即controller
	artcal.Handle(new(controller.MyController))
	artcal.Handle(new(controller.UserController))
	//*****************************************************************************************
	//以下的内容放到了service中
	//创建mysql数据库引擎，指定数据库种类，登陆名及密码，数据库名，放到了service struct中
	//Globalvaribal.engine,Globalvaribal.err=xorm.NewEngine("mysql","root:123@/test?charset=utf8")
	//设置数据库引擎
	//Globalvaribal.engine.ShowSQL(true)//打印执行的SQL语句
	//Globalvaribal.engine.SetMapper(core.SameMapper{})//表名映射方式

	//if Globalvaribal.err!=nil{
	//	Globalvaribal.app.Logger().Fatal("orm failed to initialized: %v",Globalvaribal.err)
	//}
	//iris.RegisterOnInterrupt(func(){
	//	Globalvaribal.engine.Close()
	//})
	//以上的内容放到了service中
	//*******************************************************************************************
	//启动
	Globalvaribal.app.Run(iris.Addr(":8080"))
}

////mvc.Configure的形参，在该函数中设置Handle
//func myMVC(app *mvc.Application){
//	app.Handle(new(MyController))
//}

//***************************************************************************************************
////以下放到了controller中
////定义mvc.Application的Handle struct
//type MyController struct {
//	Ctx iris.Context
//	Service service.ArticalService
//}
//
//func (m *MyController) BeforeActivation(b mvc.BeforeActivation){
//	//b.Handle("GET","/artical/{id:long}","GetArticals")
//	b.Handle("GET","/artical/{id:long}","GetArticals")
//}
//
//func (m *MyController) GetArticals() mvc.Result{
//	//获取url中的参数：id
//	id,_:=m.Ctx.Params().GetInt64("id")
//	//设置SQL操作映射的struct,Articals结构是在service包中定义的
//	artical :=service.Articals{Id:id}
//
//
//	//if ok,_:=Globalvaribal.engine.Get(&artical);ok{
//	//	return mvc.View{
//	//		Name:"information.html",
//	//		Data:iris.Map{
//	//			//"Title":"mytitle",
//	//			"Articals":artical,
//	//		},
//	//	}
//	//}else{
//	//	return mvc.Response{Path:"/view/wrong"}
//	//}
//
//	m.Service.GetArticalByID(id,&artical)
//
//	return mvc.View{
//				Name:"information.html",
//				Data:iris.Map{
//					//"Title":"mytitle",
//					"Articals":artical,
//				},
//			}
//}
////以上放到了controller中
//*******************************************************************************************

//func (m *MyController) Get() string { return "Hey" }

//func (m *MyController) GetArticals() mvc.Result{
//	return mvc.Response{ContentType:"text/html", Text:"<h1>W</h1>"}
//
//	//artical:=Articals{"1","aaa","aaa"}
//	//return mvc.View{
//	//	Name:"view/information.html",
//	//	Data:iris.Map{
//	//		"Title":"mytitle",
//	//		"Articals":artical,
//	//	}
//	//}
//}


















