package jobs

import (
	"log"
	"strings"
	"time"
)

type TblJobsApplicants struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId       int       `gorm:"type:integer"`
	Name           string    `gorm:"type:character varying"`
	EmailId        string    `gorm:"type:character varying"`
	MobileNo       string    `gorm:"type:character varying"`
	JobType        string    `gorm:"type:character varying"`
	Location       string    `gorm:"type:character varying"`
	Education      string    `gorm:"type:character varying"`
	Graduation     int       `gorm:"type:integer"`
	CompanyName    string    `gorm:"type:character varying"`
	Experience     int       `gorm:"type:integer"`
	Skills         string    `gorm:"type:character varying"`
	ImagePath      string    `gorm:"type:character varying"`
	Image          string    `gorm:"type:character varying"`
	CreatedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy      int
	ModifiedOn     time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy     int
	IsDeleted      int
	DeletedOn      time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy      int
	CreatedDate    string `gorm:"-:migration;<-:false"`
	CurrentSalary  int    `gorm:"type:integer"`
	ExpectedSalary int    `gorm:"type:integer"`
	Status         int    `gorm:"type:integer"`
	NameString     string `gorm:"-:migration;<-:false"`
	Password       string
	ResumePath     string `gorm:"type:character varying"`
	ResumeName     string `gorm:"type:character varying"`
	Jobregstatus   string `gorm:"-:migration;<-:false"`
	StorageType    string `gorm:"type:character varying"`
}

type CreateApplicantReq struct {
	Id             int
	MemberId       int
	Name           string
	EmailId        string
	MobileNo       string
	JobType        string
	Location       string
	Education      string
	Graduation     int
	CompanyName    string
	Experience     int
	Skills         string
	ImagePath      string
	Image          string
	CreatedBy      int
	ModifiedBy     int
	DeletedBy      int
	CurrentSalary  int
	ExpectedSalary int
	Status         int
	Password       string
	StorageType    string
}

type ApplicantDetails struct {
	ID             int
	JobID          int
	ApplicantID    int
	Name           string
	EmailID        string
	MobileNo       string
	JobType        string
	Gender         string
	Location       string
	Education      string
	Graduation     int
	CompanyName    string
	Experience     int
	Skills         string
	ImagePath      string
	Image          string
	CreatedOn      time.Time
	CreatedBy      int
	ModifiedOn     time.Time
	ModifiedBy     int
	IsDeleted      int
	DeletedOn      time.Time
	DeletedBy      int
	CurrentSalary  int
	ExpectedSalary int
	Status         int
	ResumePath     string
	ResumeName     string
	StorageType    string
}

// Applicant List Function//
func (Ap *Jobs) ApplicantsList(limit, offset int, filter Filter) (applicants []TblJobsApplicants, count int64, err error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return []TblJobsApplicants{}, 0, AuthErr
	}

	Jobsmodel.Dataaccess = Ap.Dataaccess

	Jobsmodel.Userid = Ap.Userid

	Applicantlist, _, _ := Jobsmodel.ApplicantsList(limit, offset, filter, Ap.DB)

	_, totalcount, _ := Jobsmodel.ApplicantsList(0, 0, filter, Ap.DB)

	var applicant []TblJobsApplicants

	for _, applicants := range Applicantlist {

		var firstn = strings.ToUpper(applicants.Name[:1])

		applicants.NameString = firstn

		applicant = append(applicant, applicants)

	}

	return applicant, totalcount, nil
}

//CreateApplicant Function//

func (Ap *Jobs) CreateApplicant(ap CreateApplicantReq) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr
	}

	var applicant TblJobsApplicants

	applicant.MemberId = ap.MemberId

	applicant.Name = ap.Name

	applicant.EmailId = ap.EmailId

	applicant.MobileNo = ap.MobileNo

	applicant.JobType = ap.JobType

	applicant.Location = ap.Location

	applicant.Education = ap.Education

	applicant.Graduation = ap.Graduation

	applicant.Experience = ap.Experience

	applicant.Skills = ap.Skills

	applicant.CurrentSalary = ap.CurrentSalary

	applicant.ExpectedSalary = ap.ExpectedSalary

	applicant.ImagePath = ap.ImagePath

	applicant.Image = ap.Image

	applicant.Status = ap.Status

	applicant.CreatedBy = ap.CreatedBy

	applicant.StorageType = ap.StorageType

	applicant.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if ap.Password != "" {
		hash_pass := HashingPassword(ap.Password)
		applicant.Password = hash_pass
	}

	err1 := Jobsmodel.ApplicantCreate(applicant, Ap.DB)

	if err1 != nil {

		return err1
	}

	return nil
}

//Get Applicant by Id function//

func (Ap *Jobs) GetApplicantById(id int) (ap TblJobsApplicants, err error) {

	//check if auth or permission enabled
	if autherr := AuthandPermission(Ap); autherr != nil {

		return TblJobsApplicants{}, autherr
	}

	applicant, err := Jobsmodel.GetApplicantById(id, Ap.DB)

	applicant.NameString = strings.ToUpper(applicant.Name[:1])

	if err != nil {

		return TblJobsApplicants{}, err
	}

	return applicant, nil

}

//Update Applicant Function//

func (Ap *Jobs) UpdateApplicant(ap CreateApplicantReq, memberid int) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr
	}

	var updateapplicant TblJobsApplicants

	updateapplicant.MemberId = memberid

	updateapplicant.Name = ap.Name

	updateapplicant.EmailId = ap.EmailId

	updateapplicant.MobileNo = ap.MobileNo

	updateapplicant.JobType = ap.JobType

	updateapplicant.Location = ap.Location

	updateapplicant.Education = ap.Education

	updateapplicant.Graduation = ap.Graduation

	updateapplicant.Experience = ap.Experience

	updateapplicant.Skills = ap.Skills

	updateapplicant.CurrentSalary = ap.CurrentSalary

	updateapplicant.ExpectedSalary = ap.ExpectedSalary

	updateapplicant.ImagePath = ap.ImagePath

	updateapplicant.Image = ap.Image

	updateapplicant.Status = ap.Status

	updateapplicant.StorageType = ap.StorageType

	updateapplicant.ModifiedBy = ap.ModifiedBy

	updateapplicant.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if ap.Password != "" {
		hash_pass := HashingPassword(ap.Password)
		updateapplicant.Password = hash_pass
	}

	err1 := Jobsmodel.ApplicantUpdate(&updateapplicant, Ap.DB)

	if err1 != nil {

		return err1
	}

	return nil
}

//Function of Delete Applicant//

func (Ap *Jobs) DeleteApplicant(memberid int, userid int) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr
	}

	var applicant TblJobsApplicants

	applicant.DeletedBy = userid

	applicant.MemberId = memberid

	applicant.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Jobsmodel.ApplicantDelete(applicant, Ap.DB)

	if err1 != nil {

		return err1
	}

	return nil
}

//Function of Getting Applicant Jobs//

func (Ap *Jobs) GetApplicantJobs(ApplicantId int, limit int, offset int) (applicantjobs []TblJobsRegisters, Totaljobs int64, err error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return []TblJobsRegisters{}, 0, AuthErr
	}

	jobs, _, err1 := Jobsmodel.GetApplicantJobs(ApplicantId, limit, offset, Ap.DB)

	if err1 != nil {

		log.Println(err1)
	}

	_, totalcount, _ := Jobsmodel.GetApplicantJobs(ApplicantId, 0, 0, Ap.DB)

	return jobs, totalcount, nil
}

//Multiselect Job delete function//

func (Ap *Jobs) MultiSelectedApplicantDelete(applicantids []int, modifiedby int) (bool, error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {
		return false, AuthErr
	}

	var applicants TblJobsApplicants
	applicants.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	applicants.DeletedBy = modifiedby
	applicants.IsDeleted = 1

	err1 := Jobsmodel.MultiSelectedApplicantDelete(&applicants, applicantids, Ap.DB)

	if err1 != nil {

		return false, err1
	}

	return true, nil
}

// multiselecte member status change
func (Ap *Jobs) MultiSelectApplicantStatus(memberid []int, status int, modifiedby int) (bool, error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return false, AuthErr
	}

	var applicantstatus TblJobsApplicants

	applicantstatus.ModifiedBy = modifiedby

	applicantstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Jobsmodel.MultiApplicantIsActive(&applicantstatus, memberid, status, Ap.DB)

	if err1 != nil {

		return false, err1
	}

	return true, nil

}

func (jobs *Jobs) GetApplicantDetails(jobId, memberId int, emailId string) (applicantDetails ApplicantDetails, err error) {

	if AuthErr := AuthandPermission(jobs); AuthErr != nil {
		return ApplicantDetails{}, AuthErr
	}

	applicantDetails, err = Jobsmodel.GetApplicantDetails(jobId, memberId, emailId, jobs.DB)
	if err != nil {

		return ApplicantDetails{}, err
	}

	return applicantDetails, nil
}

func (jobs *Jobs) CheckAlreadyRegistered(jobId int, emailId string) (count int64, err error) {

	if AuthErr := AuthandPermission(jobs); AuthErr != nil {
		return -1, AuthErr
	}

	count, err = Jobsmodel.CheckAlreadyRegistered(jobId, emailId, jobs.DB)
	if err != nil {

		return -1, err
	}

	return count, nil
}

func (jobs *Jobs) CreateJobApplication(applicationData ApplicantDetails) error {

	if AuthErr := AuthandPermission(jobs); AuthErr != nil {
		return AuthErr
	}

	err := Jobsmodel.CreateJobApplication(applicationData, jobs.DB)
	if err != nil {

		return err
	}

	return nil

}
