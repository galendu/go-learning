package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	authMap sync.Map
)

// cookie name 需要符合规则,否则该cookie会被Gin框架默默地丢弃掉
func genCookieName(ctx *gin.Context) string {
	return base64.StdEncoding.EncodeToString([]byte(ctx.Request.RemoteAddr))
}

// 登录
func login(engine *gin.Engine) {
	engine.POST("/login", func(ctx *gin.Context) {
		//为客户端生成cookie
		cookie_key := genCookieName(ctx)
		cookie_value := strconv.Itoa(time.Now().Nanosecond())
		//服务端维护所有客户端的cookie,用于对客户端进行认证
		authMap.Store(cookie_key, cookie_value)
		//把cookie发给客户端
		ctx.SetCookie(cookie_key, cookie_value,
			3000,
			"/",
			"localhost",
			false,
			true,
		)
		fmt.Printf("set cookie %s = %s to client\n", cookie_key, cookie_value)
		ctx.String(http.StatusOK, "登录成功")
	})
}

// 用户中心
func userCenter(engine *gin.Engine) {
	engine.POST("/center", authMiddleWare(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "您已通过身份认证,这里是你的私人空间")
	})
}

func authMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie_key := genCookieName(ctx)
		var cookie_value string
		//读取客户端的cookie
		for _, cookie := range ctx.Request.Cookies() {
			if cookie.Name == cookie_key {
				cookie_value = cookie.Value
				break
			}
		}

		//验证Cookie Value是否正确
		if v, ok := authMap.Load(cookie_key); !ok {
			fmt.Printf("INVALID auth cookie %s = %s\n", cookie_key, cookie_value)
			ctx.Abort() //验证不通过,调用Abort
		} else {
			if v.(string) == cookie_value {
				ctx.Next() //本中间件顺利通过
			} else {
				fmt.Printf("INVALID auth cookie %s = %s\n", cookie_key, cookie_value)
				ctx.JSON(http.StatusForbidden, gin.H{cookie_key: cookie_value})
				ctx.Abort()
			}
		}
	}
}
func main() {
	engine := gin.Default()

	//路由
	login(engine)
	userCenter(engine)

	gin.SetMode(gin.ReleaseMode)
	engine.Run(":5656")

}
