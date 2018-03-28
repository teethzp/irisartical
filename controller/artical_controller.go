package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"demo/service"
)

//定义mvc.Application的Handle struct
type MyController struct {
	Ctx iris.Context
	Service service.ArticalService
}

type UserController struct{
	Ctx iris.Context
	Service service.UserService
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation){
	b.Handle("GET","/user/{name:string}","GetUser")
}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation){
	b.Handle("GET","/artical/{id:long}","GetArticals")
	b.Handle("GET","/artical/all","GetAllArticals")
	b.Handle("POST","/artical/insert","InsertArticals")
	b.Handle("DELETE","/artical/delete/{tital:string}","DeleteArticals")
}


func (m *MyController) GetArticals() mvc.Result{
	//获取url中的参数：id
	id,_:=m.Ctx.Params().GetInt64("id")
	//设置SQL操作映射的struct,Articals结构是在service包中定义的
	artical :=service.Articals{Id:id}


	//if ok,_:=Globalvaribal.engine.Get(&artical);ok{
	//	return mvc.View{
	//		Name:"information.html",
	//		Data:iris.Map{
	//			//"Title":"mytitle",
	//			"Articals":artical,
	//		},
	//	}
	//}else{
	//	return mvc.Response{Path:"/view/wrong"}
	//}

	//调用service的方法与数据库操作
	err:=m.Service.GetArticalByID(id,&artical)

	if(err!=nil){
		return mvc.Response{
			Path:"/view/wrong.html",
		}
	}else{
		return mvc.View{
			Name:"information.html",
			Data:iris.Map{
				//"Title":"mytitle",
				"Articals":artical,
			},
		}
	}
}

func (m *MyController) GetAllArticals() []service.Articals{
	m.Ctx.Application().Logger().Infof("excute GetAllArticals")
	articalsice:=make([]service.Articals,0)
	err,length:=m.Service.GetAllArticals(&articalsice)
	length2:=len(articalsice)
	m.Ctx.Application().Logger().Infof("GetALlArticals length:%d,%d",length,length2)
	if(err !=nil){
		m.Ctx.Application().Logger().Infof("GetALlArticals %s",err.Error())
		return articalsice
	}else{
		return articalsice
	}
}

//curl -d '{"Headline":"ppppp","Content":"zzzzzzz"}' "http://127.0.0.1:8080/rout/artical/insert"
func (m *MyController) InsertArticals() string{
	var artical service.Articals
	m.Ctx.ReadJSON(&artical)
	err:=m.Service.InsertArtical(&artical)

	if(err != nil){
		m.Ctx.Application().Logger().Infof("insertArticals %s",err.Error())
		return err.Error()
	}else{
		return "OK"
	}

}

//curl -X DELETE "http://127.0.0.1:8080/rout/artical/delete/ppppp"
func (m *MyController) DeleteArticals() string{
	var artical service.Articals
	tital:=m.Ctx.Params().GetTrim("tital")
	artical=service.Articals{Headline:tital}
	err:=m.Service.DeleteArtical(&artical)

	if(err != nil){
		m.Ctx.Application().Logger().Infof("DeleteArticals %s",err.Error())
		return err.Error()
	}else{
		return "OK"
	}
}

//*************************************************************************
//以下是UserController控制器的方法
func (m *UserController) GetUser() service.Userinfo{
	name:=m.Ctx.Params().GetTrim("name")
	userinfo:=service.Userinfo{Name:name}
	err:=m.Service.GetUser(name,&userinfo)
	if(err!=nil){
		return userinfo
	}else{
		return userinfo
	}
}

















