package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/motojouya/geezer_auth/pkg/shelter/jwt"
	"github.com/motojouya/geezer_auth/pkg/shelter/user"
	localRepository "github.com/motojouya/ddd_go/pkg/local/repository"
	userInput "github.com/motojouya/ddd_go/pkg/user/input"
)

func Hand[C any, I any, O any](createControl func() (C, error), handleControl func(C, I, *user.Authentic) (O, error)) echo.HandlerFunc {
	return func(c echo.Context) error {

		header := userInput.RequestHeader{}
		if err := (&echo.DefaultBinder{}).BindHeaders(c, &header); err != nil {
			return err
		}

		var request I
		if err := c.Bind(&request); err != nil {
			return err
		}

		authentic, err := getAuthentic(header)
		if err != nil {
			return err
		}

		control, err := createControl()
		if err != nil {
			return err
		}

		response, err := handleControl(control, request, authentic)
		if err != nil {
			return err
		}

		return c.JSON(200, response)
	}
}

func getAuthentic(header userInput.RequestHeader) (*user.Authentic, error) {
	var jwtParse, err = localRepository.GetEnv[jwt.JwtParse]() // FIXME use cache!
	if err != nil {
		return nil, err
	}

	token, err := header.GetBearerToken()
	if err != nil {
		return nil, err
	}

	if token == "" {
		return nil, nil
	}

	return jwtParse.Parse(token)
}
