package controller

import (
	"net/http"
	"reflect"

	echo "github.com/labstack/echo"

	demo "github.com/kaporzhu/echo-demo"
)

const SessionUserKey = "session:user"

type Condition interface {
	Validate(req interface{}, resp interface{}, c echo.Context) (bool, error)
}

type loginRequiredCondition struct{}

var LoginRequired loginRequiredCondition

func (loginRequiredCondition) Validate(req interface{}, resp interface{}, c echo.Context) (bool, error) {
	user := c.Get(SessionUserKey)
	if nil == user {
		return false, echo.NewHTTPError(http.StatusUnauthorized, "Login required")
	}
	return true, nil
}

type Controller func(req interface{}, resp interface{}, c echo.Context) error

func controllerWrap(controller Controller, request interface{}, response interface{}, conditions ...Condition) echo.HandlerFunc {
	return func(c echo.Context) error {
		// retrieve user from grpc
		// user := AccountService.GetUser(id)
		// c.Set(SessionUserKey, user)

		req := reflect.New(reflect.TypeOf(request).Elem()).Interface()
		resp := reflect.New(reflect.TypeOf(response).Elem()).Interface()
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Bind request failed")
		}

		// validate conditions
		for _, condition := range conditions {
			if pass, err := condition.Validate(req, resp, c); pass == false {
				return err
			}
		}

		if err := controller(req, resp, c); err != nil {
			return err
		}
		return demo.Negotiate(c, http.StatusOK, resp)
	}
}
