package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func SendError(c *gin.Context, code int, error string) {
	c.JSON(code, Error{
		Code: code,
		Error: error,
	})
}

func SendResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  message,
		Data: data,
	})
}

var Suff map[string]int

func Init() {
	Suff = map[string]int{
		".xbm":   1,
		".tif":   1,
		".pjp":   1,
		".svgz":  1,
		".jpg":   1,
		".jpeg":  1,
		".ico":   1,
		".tiff":  1,
		".gif":   1,
		".svg":   1,
		".jfif":  1,
		".webp":  1,
		".png":   1,
		".bmp":   1,
		".pjpeg": 1,
		".avif":  1,
	}
}
