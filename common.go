package jobs

import (
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	ErrorEmpty      = errors.New("given some values is empty")
	TZONE, _        = time.LoadLocation(os.Getenv("TIME_ZONE"))
)

type JobsModel struct{}

var Jobsmodel JobsModel

func AuthandPermission(jobs *Jobs) error {

	//check auth enable if enabled, use auth pkg otherwise it will return error
	if jobs.AuthEnable && !jobs.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled, use team-role pkg otherwise it will return error
	if jobs.PermissionEnable && !jobs.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}

func HashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}
