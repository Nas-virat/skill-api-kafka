package response

import (
	"gokafka/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success(ctx *gin.Context, StatusCode int, data any) {
	ctx.JSON(StatusCode, Response{
		Status:"success",
		Data:data,
	})
}

func SuccessMsg(ctx *gin.Context, StatusCode int, msg string) {
	ctx.JSON(StatusCode, Response{
		Status:"success",
		Message:msg,
	})
}


func Error(ctx *gin.Context, err error) {
	switch e := err.(type){
		case errs.Err:
			ctx.JSON(e.StatusCode,Response{
				Status:"error",
				Message: e.Message,
			})
		case error:
			ctx.JSON(http.StatusInternalServerError,Response{
				Status: "error",
				Message: err.Error(),
			})
	}
}

