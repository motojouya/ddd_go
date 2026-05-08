package controller

import (
	"fmt"
	"os/user"

	geezerUser "github.com/motojouya/geezer_auth/pkg/shelter/user"
	databaseRepository "github.com/motojouya/ddd_go/pkg/database/repository"
)

func HandleCmd[C any, I any, O any](createControl func() (C, error), handleControl func(C, I, *geezerUser.Authentic) (O, error), entry I) (O, error) {
	var zeroResult O

	user, err := user.Current()
	if err != nil {
		return zeroResult, err
	}
	fmt.Println("User Name: " + user.Username) // TODO 本当は、Authenticに変換してアプリケーションに渡したい

	// db connectionを取得してserver停止時にclose
	dbGetter := databaseRepository.NewDatabaseGet()
	db, err := dbGetter.GetDatabase()
	defer db.Close()
	if err != nil {
		return zeroResult, err
	}

	control, err := createControl()
	if err != nil {
		return zeroResult, err
	}

	return handleControl(control, entry, nil)
}
