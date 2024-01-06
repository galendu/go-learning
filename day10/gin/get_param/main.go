package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 从GET请求的URL中获取参数
func url(engine *gin.Engine) {
	engine.GET("/student", func(ctx *gin.Context) {
		name := ctx.Query("name")
		addr := ctx.DefaultQuery("addr", "China") //如果没传addr参数,则默认为China
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

// Restful风格的url中获取参数
func restful(engine *gin.Engine) {
	engine.GET("/student/:name/*addr", func(ctx *gin.Context) {
		name := ctx.Param("name")
		addr := ctx.Param("addr")
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

// 从post表单中获取参数
func post(engine *gin.Engine) {
	engine.POST("/student", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		addr := ctx.DefaultPostForm("addr", "China") //如果没有传addr参数,则默认为China
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}

// 上传单个文件
func upload_file(engine *gin.Engine) {
	//限制表单上传大小为8M,默认上线是32M
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			fmt.Printf("get file error %v\n", err)
			ctx.String(http.StatusInternalServerError, "upload file failed")
		} else {
			ctx.SaveUploadedFile(file, "./data/"+file.Filename) //把用户上次的文件存到data目录下
			ctx.String(http.StatusOK, file.Filename)
		}
	})
}

// 上传多个文件
func upload_multi_file(engine *gin.Engine) {
	engine.POST("/upload_files", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
		} else {
			//从MultipartForm中获取上传的文件
			files := form.File["files"]
			for _, file := range files {
				ctx.SaveUploadedFile(file, "./data/"+file.Filename)
			}
			ctx.String(http.StatusOK, "upload "+strconv.Itoa(len(files))+" files")
		}
	})
}

// type Student struct {
// 	Name string `form:"username" json:"name" uri:"user" xml:"user"  yaml:"user" binding:"required"`
// 	Addr string `form:"addr" json:"addr" uri:"addr" xml:"addr" yaml:"addr" binding:"required"`
// }

type Student struct {
	Name string `form:"username" json:"name" uri:"user" xml:"user" yaml:"user" binding:"required"`
	Addr string `form:"addr" json:"addr" uri:"addr" xml:"addr" yaml:"addr" binding:"required"`
}

func formBind(engine *gin.Engine) {
	engine.POST("/stu/form", func(ctx *gin.Context) {
		var stu Student
		//跟ShouldBind对应的时MustBind. MustBind内部会调用ShouldBind.如果ShouldBind发生error会直接c.AbortWithError(http.StatusBadRequest,err)
		if err := ctx.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		}
	})
}

func jsonBind(engine *gin.Engine) {
	engine.POST("/stu/json", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.URL)
		var stu Student
		if err := ctx.ShouldBindJSON(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" line in "+stu.Addr)
		}
	})
}

func uriBind(engine *gin.Engine) {
	engine.GET("/stu/uri/:user/:addr", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.URL)
		var stu Student
		if err := ctx.ShouldBindUri(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" line in "+stu.Addr)
		}
	})
}

func xmlBind(engine *gin.Engine) {
	engine.POST("/stu/xml", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBindXML(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
		}
	})
}

func yamlBind(engine *gin.Engine) {
	engine.POST("/stu/yaml", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.URL)
		var stu Student
		if err := ctx.ShouldBindYAML(&stu); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			ctx.String(http.StatusOK, stu.Name+" line in "+stu.Addr)
		}
	})
}

func main() {
	engine := gin.Default()

	url(engine)
	restful(engine)
	post(engine)
	upload_file(engine)
	upload_multi_file(engine)

	formBind(engine)
	jsonBind(engine)
	uriBind(engine)
	xmlBind(engine)
	yamlBind(engine)

	engine.Run(":5656")
}
