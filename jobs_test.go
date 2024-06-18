package jobs

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db connection
func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "admin123",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurtcms_may27",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})

	if err != nil {

		fmt.Println("Failed to connect to database:", err)

	}
	if err != nil {

		return nil, err

	}

	return db, nil
}

// Test function of Jobslist//
func TestJobsList(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	//list jobs
	if permisison {

		joblist, count, _ := Jobs.JobsList(10, 0, Filter{})

		fmt.Println(joblist, count)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

//Test function of Create JOb//

func TestCreateJob(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	//Create Job
	if permisison {

		err := Jobs.CreateJob(CreateJobReq{JobTitle: "Developer", JobDescription: "developing, coding, installing, and maintaining software systems.", JobType: "Full-Time", JobLocation: "Chennai"})

		if err != nil {

			fmt.Println(err)
		}
	} else {

		log.Println("permissions enabled not initialised")

	}
}

func TestUpdateJob(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		err := Jobs.UpdateJob(CreateJobReq{JobTitle: "Manager", JobDescription: "Develops, coordinates, and enforces systems, policies, procedures, and productivity", JobType: "Hybrid-Work", JobLocation: "Bangalore", Id: 1})

		if err != nil {

			fmt.Println(err)
		}

	} else {

		log.Println("permissions enabled not initialised")
	}

}

func TestDeleteJob(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		err := Jobs.DeleteJob(1, 1)

		if err != nil {

			fmt.Println(err)
		}
	} else {

		log.Println("permissions enabled not initialised")
	}

}

//Multiselect job delete function//

func TestMultiselectedJobDelete(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	var id = []int{1, 2, 3}

	deletedby := 1

	if permisison {

		_, err := Jobs.MultiSelectedJobDelete(id, deletedby)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}
}

// Test Function of Multiselect job status change
func TestMultiselectedJobStatus(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	var id = []int{1, 2, 3}

	modifiedby := 1

	status := 1

	if permisison {

		_, err := Jobs.MultiSelectJobsStatus(id, status, modifiedby)

		if err != nil {

			fmt.Println(err)
		}
	} else {

		log.Println("permissions enabled not initialised")
	}

}

//Test function of GetJobapplicant//

func TestGetJobApplicant(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := JobsSetup(Config{

		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		jobapplicants, count, err := Jobs.GetJobApplicant(1, 10, 0, Filter{})

		fmt.Println(jobapplicants, count, err)

	} else {

		log.Println("permissions enabled not initialised")
	}
}
