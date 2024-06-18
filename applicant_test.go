package jobs

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
)

//Test Function of ApplicantList//

func TestApplicantList(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicants := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		applicants, count, err := Applicants.ApplicantsList(10, 0, Filter{})

		fmt.Println(applicants, count, err)

	} else {

		log.Println("permissions enabled not initialised")
	}
}

//Function of test Create Applicant//

func TestCreateApplicant(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicants := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		err := Applicants.CreateApplicant(CreateApplicantReq{Name: "Joe", Location: "Chennai", JobType: "Full-Time"})

		if err != nil {

			fmt.Println(err)
		}
	} else {

		log.Println("permissions enabled not initialised")
	}
}

//Test Function of UpdateApplicant//

func TestUpdateApplicant(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicants := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		uerr := Applicants.UpdateApplicant(CreateApplicantReq{Name: "John", Education: "Bsc", Experience: 1, JobType: "Remote-Work"}, 1)

		if uerr != nil {

			fmt.Println(uerr)
		}
	} else {

		log.Println("permissions enabled not initialised")
	}
}

//Test Function of DeleteApplicant//

func TestDeleteApplicant(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicants := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	id := 1

	deletedby := 1

	if permisison {

		derr := Applicants.DeleteApplicant(id, deletedby)

		if derr != nil {

			fmt.Println(derr)
		}
	} else {

		log.Println("permissions enabled not initialised")
	}

}

//Test Function of Multiselected Applicant Delete//

func TestMultiSelectedApplicantDelete(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicant := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	var id = []int{1, 2, 3}

	deletedby := 1

	if permisison {

		_, err := Applicant.MultiSelectedApplicantDelete(id, deletedby)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

//Test Function of MultiselectApplicant Status change//

func TestMultiSelectedApplicantStatus(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicant := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	var id = []int{1, 2, 3}

	modifiedby := 1

	status := 1

	if permisison {

		_, err := Applicant.MultiSelectApplicantStatus(id, status, modifiedby)

		if err != nil {

			fmt.Println(err)
		}
	} else {

		log.Println("permissions enabled not initialised")
	}
}

//Test function of GetApplicantJobs//

func TestGetApplicantJobs(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicant := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		jobs, count, err := Applicant.GetApplicantJobs(1, 10, 0)

		fmt.Println(jobs, count, err)
	} else {

		log.Println("permissions enabled not initialised")
	}
}

//Test Function of GetApplicantById//

func TestGetApplicantById(t *testing.T) {

	db, _ := DBSetup()

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         db,
	})
	Auth.RoleId = 1
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

	Applicant := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	id := 1

	if permisison {

		applicant, err := Applicant.GetApplicantById(id)

		fmt.Println(applicant, err)
	} else {

		log.Println("permissions enabled not initialised")
	}

}
