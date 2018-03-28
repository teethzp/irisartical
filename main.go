package main
import (
	//"time"

	"github.com/kataras/iris"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	//_ "github.com/mattn/go-sqlite3"
)

//type Articals struct {
//	Id int64
//	Headline string
//	Content string
//}

func main(){
	app:=iris.New()
	app.RegisterView(iris.HTML("./view",".html"))
	engine,err:=xorm.NewEngine("mysql","root:123@/test?charset=utf8")
	engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})

	if err!=nil{
		app.Logger().Fatal("orm failed to initialized: %v",err)
	}

	iris.RegisterOnInterrupt(func(){
		engine.Close()
	})

	app.Get("/insert",func(ctx iris.Context){
		artical:=&Articals{Headline:"区块链",Content:"区块链"}
		//artical :=&Articals{Headline:"aaa",Content:"bbb"}
		engine.Insert(artical)
		ctx.Writef("%s",artical.Content)
	})

	app.Get("/artical",func(ctx iris.Context){
		artical :=Articals{Id:5}
		if ok,_:=engine.Get(&artical);ok{
			//ctx.Writef("artical: %s",artical.Content)
			ctx.ViewData("message",artical.Headline)
			ctx.ViewData("information",artical.Content)
			ctx.View("information1.html")
		} else{
			ctx.Writef("wrong")
		}

	})

	app.Run(iris.Addr(":8080"))
}












