package jobs

import "C"
import (
	"log"
	"strings"
	"time"

	"github.com/spurtcms/member"
	"gorm.io/gorm"
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
}

// Member package connection

func (Ap *Jobs) DBconf() *member.Member {
	var memberconfig = member.MemberSetup(member.Config{DB: Ap.DB, AuthEnable: Ap.AuthEnable, PermissionEnable: Ap.PermissionEnable})
	return memberconfig
}

// Applicant List Function//
func (Ap *Jobs) ApplicantsList(limit, offset int, filter Filter) (applicants []TblJobsApplicants, count int64, err error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return []TblJobsApplicants{}, 0, AuthErr
	}

	Applicantlist, _, _ := Jobsmodel.ApplicantsList(limit, offset, filter, Ap.DB)

	_, totalcount, _ := Jobsmodel.ApplicantsList(0, 0, filter, Ap.DB)

	var applicant []TblJobsApplicants

	for _, applicants := range Applicantlist {

		var firstn = strings.ToUpper(applicants.Name[:1])

		applicants.CreatedDate = applicants.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		if !applicants.ModifiedOn.IsZero() {

			applicants.CreatedDate = applicants.ModifiedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		} else {
			applicants.CreatedDate = applicants.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		}

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

	db := Ap.DBconf()

	amember, err := db.CreateMember(member.MemberCreationUpdation{

		FirstName: ap.Name,

		Email: ap.EmailId,

		MobileNo: ap.MobileNo,

		ProfileImage: ap.Image,

		ProfileImagePath: ap.ImagePath,

		IsActive: 1,

		Password: ap.Password,

		CreatedBy: ap.CreatedBy,
	})

	if err != nil {
		log.Println(err)
	}

	var applicant TblJobsApplicants

	applicant.MemberId = amember.Id

	applicant.Name = ap.Name

	applicant.EmailId = ap.EmailId

	applicant.MobileNo = ap.MobileNo

	applicant.JobType = ap.JobType

	applicant.Password = ap.Password

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

	applicant.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

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

	db := Ap.DBconf()

	err := db.UpdateMember(member.MemberCreationUpdation{

		FirstName: ap.Name,

		Email: ap.EmailId,

		MobileNo: ap.MobileNo,

		ProfileImage: ap.Image,

		ProfileImagePath: ap.ImagePath,

		IsActive: 1,

		Password: ap.Password,

		ModifiedBy: ap.ModifiedBy,
	}, memberid)

	if err != nil {
		log.Println(err)
	}

	var updateapplicant TblJobsApplicants

	updateapplicant.MemberId = memberid

	updateapplicant.Name = ap.Name

	updateapplicant.EmailId = ap.EmailId

	updateapplicant.MobileNo = ap.MobileNo

	updateapplicant.JobType = ap.JobType

	updateapplicant.Password = ap.Password

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

	updateapplicant.ModifiedBy = ap.ModifiedBy

	updateapplicant.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Jobsmodel.ApplicantUpdate(&updateapplicant, Ap.DB)

	if err1 != nil {

		return err1
	}

	return nil
}

//Function of Delete Applicant//

func (Ap *Jobs) DeleteApplicant(memberid int, userid int, DB *gorm.DB) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr
	}

	db := Ap.DBconf()

	err := db.DeleteMember(memberid, userid)

	if err != nil {

		return err
	}
	var applicant TblJobsApplicants

	applicant.DeletedBy = userid

	applicant.MemberId = memberid

	applicant.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Jobsmodel.ApplicantDelete(applicant, Ap.DB)

	if err1 != nil {

		return err
	}

	return nil
}

//Function of Getting Applicant Jobs//

func (Ap *Jobs) GetApplicantJobs(ApplicantId int, limit int, offset int) (applicantjobs []TblJobs, Totaljobs int64, err error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return []TblJobs{}, 0, AuthErr
	}

	var jobslist []TblJobs

	jobs, _, err1 := Jobsmodel.GetApplicantJobs(ApplicantId, limit, offset, Ap.DB)

	if err1 != nil {

		log.Println(err1)
	}

	for _, Job := range jobs {

		for _, val := range Job.JobList {

			val.CreatedDate = val.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

			if !val.ModifiedOn.IsZero() {

				val.CreatedDate = val.ModifiedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

			} else {
				val.CreatedDate = val.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

			}
			jobslist = append(jobslist, val)

		}

	}

	_, totalcount, _ := Jobsmodel.GetApplicantJobs(ApplicantId, 0, 0, Ap.DB)

	return jobslist, totalcount, nil
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

	err := Jobsmodel.MultiSelectedApplicantDelete(&applicants, applicantids, Ap.DB)

	if err != nil {

		return false, err
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

	err := Jobsmodel.MultiApplicantIsActive(&applicantstatus, memberid, status, Ap.DB)

	if err != nil {

		return false, err
	}

	return true, nil

}
