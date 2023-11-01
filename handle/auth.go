package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AuthUser 用户认证 返回类型：HandlerFunc
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jilijili-token")
		c.Next()
		return
		if err != nil {
			log.Println("获取Cookie错误:", err)
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"flag":    401,
				"message": "访问未授权",
				"data":    nil,
			})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
		log.Println("访问令牌: cookie", cookie)
		// 验证通过，会继续访问下一个中间件
		c.Next()
	}

}

// InterceptParams 参数拦截 分页设置
func InterceptParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数值
		page := c.Query("page")
		size := c.Query("size")

		// 判断参数是否为空，如果为空则设置默认值
		if page == "" {
			page = "1"
		}
		if size == "" {
			size = "10"
		}

		// 将处理后的参数值设置回上下文中
		c.Set("page", page)
		c.Set("size", size)

		// 继续处理请求
		c.Next()
	}
}
