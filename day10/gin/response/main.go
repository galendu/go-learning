package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func text(engine *gin.Engine) {
	engine.GET("/user/text", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hi boy")
	})
}
func json1(engine *gin.Engine) {
	engine.GET("/user/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "zcy", "addr": "bj"})
	})
}

func json2(engine *gin.Engine) {
	var stu struct {
		Name string
		Addr string
	}
	stu.Name = "zcy"
	stu.Addr = "BJ"
	engine.GET("/user/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, stu)
	})
}

func jsonp(engine *gin.Engine) {
	var stu struct {
		Name string
		Addr string
	}
	stu.Name = "zcy"
	stu.Addr = "BJ"
	engine.GET("user/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, stu)
	})
}
func xml(engine *gin.Engine) {
	var stu struct {
		Name string
		Addr string
	}
	stu.Name = "zcy"
	stu.Addr = "BJ"
	engine.GET("/user/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, stu)
	})
}

func yaml(engine *gin.Engine) {
	var stu struct {
		Name string
		Addr string
	}
	stu.Name = "zcy"
	stu.Addr = "BJ"
	engine.GET("/user/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, stu)
	})
}

func html(engine *gin.Engine) {
	engine.LoadHTMLFiles("static/template.html")
	engine.GET("/user/html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "template.html", gin.H{"title": "用户信息", "name": "zcy", "addr": "bj"})
	})
}

func redirect(engine *gin.Engine) {
	engine.GET("/not_exists", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://localhost:5656/user/html")
	})
}
func main() {
	engine := gin.Default()
	text(engine)
	json1(engine)
	json2(engine)
	jsonp(engine)
	xml(engine)
	yaml(engine)
	html(engine)
	redirect(engine)
	engine.Run(":5656")

}
