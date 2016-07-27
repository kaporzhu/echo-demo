package controller

import (
	"net/http"

	"github.com/labstack/echo"

	pb "github.com/kaporzhu/echo-demo/proto"
)

var SampleController = sampleController{}

type sampleController struct{}

func (sampleController) Load(e *echo.Echo) {
	e.POST("/success", controllerWrap(successController, &pb.SampleReq{}, &pb.SampleResp{}))
	e.POST("/unauthed", controllerWrap(unauthorizedController, &pb.SampleReq{}, &pb.SampleResp{},
		LoginRequired))
	e.POST("/error", controllerWrap(errorController, &pb.SampleReq{}, &pb.SampleResp{}))
}

func successController(request interface{}, response interface{}, c echo.Context) error {
	resp := response.(*pb.SampleResp)
	resp.Text = "kaporzhu"
	return nil
}

func unauthorizedController(request interface{}, response interface{}, c echo.Context) error {
	return nil
}

func errorController(request interface{}, response interface{}, c echo.Context) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "Whatever")
}
