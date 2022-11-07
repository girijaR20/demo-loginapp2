package userapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"demo-loginapp2/domain_mydomain/usecase/loginapp2"
	"demo-loginapp2/shared/infrastructure/logger"
	"demo-loginapp2/shared/model/payload"
	"demo-loginapp2/shared/util"
)

// loginapp2Handler ...
func (r *Controller) loginapp2Handler() gin.HandlerFunc {

	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		//Username string `json:"username"`
		//Password string `json:"password"`
		Status string `json:"status"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req loginapp2.InportRequest
		req.Username = jsonReq.Username
		req.Password = jsonReq.Password

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.Loginapp2Inport.Execute(ctx, req)
		fmt.Println("handler res ", res)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		//jsonRes.Username = req.Username
		//jsonRes.Password = req.Password
		jsonRes.Status = res.Status

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
