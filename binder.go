package demo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/labstack/echo"
)

type binder struct {
	echoBinder echo.Binder
}

func (b *binder) Bind(i interface{}, c echo.Context) error {
	req := c.Request()
	ctype := req.Header().Get(echo.HeaderContentType)
	if strings.HasPrefix(ctype, echo.MIMEApplicationProtobuf) {
		data, _ := ioutil.ReadAll(req.Body())
		if err := proto.Unmarshal(data, i.(proto.Message)); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return nil
	}

	// pass Bind call to default Echo binder
	// process the request without Protobuf content
	return b.echoBinder.Bind(i, c)
}

func Negotiate(c echo.Context, status int, obj interface{}) error {
	accept := strings.ToLower(c.Request().Header().Get("Accept"))
	switch {
	case strings.Contains(accept, echo.MIMEApplicationProtobuf):
		content, _ := proto.Marshal(obj.(proto.Message))
		resp := c.Response()
		resp.WriteHeader(status)
		resp.Header().Set("Content-Type", echo.MIMEApplicationProtobuf)
		resp.Write(content)
	case strings.Contains(accept, echo.MIMEApplicationJSON):
		content, _ := json.Marshal(obj)
		resp := c.Response()
		resp.WriteHeader(status)
		resp.Header().Set("Content-Type", echo.MIMEApplicationJSON)
		resp.Write(content)
	}
	return nil
}
