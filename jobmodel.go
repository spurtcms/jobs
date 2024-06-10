package jobs

import (
	"github.com/spurtcms/categories"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CreateJobReq struct {
	Id             int
	CategoriesId   int
	JobTitle       string
	JobSlug        string
	JobDescription string
	JobLocation    string
	JobType        string
	Education      string
	Department     string
	Experience     string
	Salary         string
	CreatedBy      int
	DeletedBy      int
	Keywords       string
	Skill          string
	MinimumYears   int
	MaximumYears   int
	PostedDate     time.Time
	ValidThrough   time.Time
	Status         int
	ModifiedBy     int
}

type Filter struct {
	Keyword        string
	JobType        string
	Name           string
	EmailId        string
	MobileNo       string
	JobId          int
	Status         string
	JobTitle       string
	ApplicantName  string
	ApplicantEmail string
}
type TblJobs struct {
	Id             int                        `gorm:"primaryKey;auto_increment;type:serial"`
	CategoriesId   int                        `gorm:"type:integer"`
	JobTitle       string                     `gorm:"type:character varying"`
	JobDescription string                     `gorm:"type:character varying"`
	JobLocation    string                     `gorm:"type:character varying"`
	JobType        string                     `gorm:"type:character varying"`
	Education      string                     `gorm:"type:character varying"`
	Department     string                     `gorm:"type:character varying"`
	Experience     string                     `gorm:"type:character varying"`
	Salary         string                     `gorm:"type:character varying"`
	CreatedOn      time.Time                  `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy      int                        `gorm:"type:integer"`
	CreatedDate    string                     `gorm:"-:migration;<-:false"`
	IsDeleted      int                        `gorm:"type:integer"`
	DeletedOn      time.Time                  `gorm:"DEFAULT:NULL"`
	DeletedBy      int                        `gorm:"DEFAULT:NULL"`
	ApplicantsList []TblJobsApplicants        `gorm:"foreignKey:Id;"`
	Keywords       string                     `gorm:"type:character varying"`
	Skill          string                     `gorm:"type:character varying"`
	MinimumYears   int                        `gorm:"type:integer"`
	MaximumYears   int                        `gorm:"type:integer"`
	PostedDate     time.Time                  `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ValidThrough   time.Time                  `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Status         int                        `gorm:"type:integer"`
	Startdate      string                     `gorm:"-:migration;<-:false"`
	Enddate        string                     `gorm:"-:migration;<-:false"`
	ModifiedOn     time.Time                  `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy     int                        `gorm:"type:integer"`
	CategoryNames  []categories.TblCategories `gorm:"-"`
	JobList        []TblJobs                  `gorm:"foreignKey:Id;"`
}

type TblJobsEducation struct {
	Id        int    `gorm:"primaryKey;auto_increment;type:serial"`
	Education string `gorm:"type:character varying"`
}

type TblJobsDepartment struct {
	Id         int    `gorm:"primaryKey;auto_increment;type:serial"`
	Department string `gorm:"type:character varying"`
}

type TblJobsRegisters struct {
	Id             int `gorm:"primaryKey;auto_increment;type:serial"`
	JobId          int
	ApplicantId    int
	CreatedBy      int                 `gorm:"type:integer"`
	CreatedOn      time.Time           `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ApplicantsList []TblJobsApplicants `gorm:"foreignKey:Id;"`
	JobList        []TblJobs           `gorm:"foreignKey:Id;"`
}

func (jobsmodel JobsModel) JobsList(offset int, limit int, filter Filter, DB *gorm.DB) (job []TblJobs, Totaljob int64, err error) {

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "0"
	}

	query := DB.Debug().Table("tbl_jobs").Where("is_deleted = 0").Order("id desc")

	if filter.Keyword != "" {

		if filter.Keyword == "0" || filter.Keyword == "1" {

			query = query.Where(" tbl_jobs.status=?", filter.Keyword)

		} else {

			query = query.Where("LOWER(TRIM(job_title)) ILIKE LOWER(TRIM(?)) OR LOWER(TRIM(job_type)) ILIKE LOWER(TRIM(?))  ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}
	}

	if filter.JobType != "" {

		query = query.Debug().Where("tbl_jobs.job_type=?", filter.JobType)
	}

	if filter.JobTitle != "" {

		query = query.Debug().Where("LOWER(TRIM(job_title)) ILIKE LOWER(TRIM(?))", "%"+filter.JobTitle+"%")
	}

	if filter.Status == "InActive" {

		filter.Status = "0"

	} else if filter.Status == "Active" {

		filter.Status = "1"

	}

	if filter.Status != "" {

		query = query.Where("tbl_jobs.status=?", filter.Status)

	}
	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&job)

		return job, 0, err

	}

	query.Find(&job).Count(&Totaljob)

	return job, Totaljob, nil

}

//Create Job Function//

func (jobsmodel JobsModel) JobCreate(job *TblJobs, DB *gorm.DB) error {

	if err := DB.Model(TblJobs{}).Create(&job).Error; err != nil {
		return err
	}
	return nil
}

//Get Jobdetails By Id//

func (jobsmodel JobsModel) JobDetailsById(jobs *TblJobs, id int, DB *gorm.DB) error {

	if err := DB.Model(TblJobs{}).Where("id = ?", id).First(&jobs).Error; err != nil {
		return err
	}
	return nil
}

//Update Function of Job//

func (jobsmodel JobsModel) JobUpdate(job *TblJobs, DB *gorm.DB) error {

	if err := DB.Model(TblJobs{}).Where("id = ?", job.Id).UpdateColumns(map[string]interface{}{"job_title": job.JobTitle, "categories_id": job.CategoriesId, "job_type": job.JobType, "job_location": job.JobLocation, "job_description": job.JobDescription, "skill": job.Skill, "salary": job.Salary, "education": job.Education, "minimum_years": job.MinimumYears, "maximum_years": job.MaximumYears, "status": job.Status, "posted_date": job.PostedDate, "valid_through": job.ValidThrough, "keywords": job.Keywords, "modified_on": job.ModifiedOn, "modified_by": job.ModifiedBy}).Error; err != nil {
		return err
	}

	return nil
}

//Delete Function of Job//

func (jobsmodel JobsModel) JobDelete(job *TblJobs, DB *gorm.DB) error {

	if err := DB.Model(TblJobs{}).Where("id = ?", job.Id).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": job.DeletedOn, "deleted_by": job.DeletedBy}).Error; err != nil {

		return err
	}

	return nil
}

//Multiselect Delete Function of Job//

func (jobsmodel JobsModel) MultiSelectedJobDelete(job *TblJobs, id []int, DB *gorm.DB) error {

	if err := DB.Model(TblJobs{}).Where("id in (?)", id).UpdateColumns(map[string]interface{}{"is_deleted": job.IsDeleted, "deleted_on": job.DeletedOn, "deleted_by": job.DeletedBy}).Error; err != nil {

		return err

	}

	return nil
}

//Multiselect  status change function of job//

func (jobsmodel JobsModel) MultiJobsIsActive(job *TblJobs, jobid []int, status int, DB *gorm.DB) error {

	if err := DB.Model(TblJobs{}).Where("id in (?)", jobid).UpdateColumns(map[string]interface{}{"status": job.Status, "modified_by": job.ModifiedBy, "modified_on": job.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

//GetJobApplicant By JobId Function//

func (jobsmodel JobsModel) GetJobApplicantByJobId(id int, limit int, offset int, filter Filter, DB *gorm.DB) (applicant []TblJobs, Totalapplicants int64, err error) {

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "0"
	}

	query := DB.Debug().Preload("ApplicantsList", func(db *gorm.DB) *gorm.DB {
		return db.Order("id asc")
	}).Table("tbl_jobs_registers").Select("tbl_jobs_applicants.id").Joins("inner join tbl_jobs on tbl_jobs.id = tbl_jobs_registers.job_id").Joins("inner join tbl_jobs_applicants on tbl_jobs_applicants.id =tbl_jobs_registers.applicant_id").Where("tbl_jobs.id = ?", id).Find(&applicant)

	if filter.Keyword != "" {

		if filter.Keyword == "0" || filter.Keyword == "1" {

			query = query.Where(" tbl_jobs_applicants.status=?", filter.Keyword)

		} else {

			query = query.Debug().Where("LOWER(TRIM(name)) ILIKE LOWER(TRIM(?))   OR LOWER(TRIM(email_id)) ILIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}
	}

	if filter.ApplicantName != "" {

		query = query.Where("LOWER(TRIM(name)) ILIKE LOWER(TRIM(?))", "%"+filter.ApplicantName+"%")
	}

	if filter.ApplicantEmail != "" {

		query = query.Debug().Where("LOWER(TRIM(email_id)) ILIKE LOWER(TRIM(?))", "%"+filter.ApplicantEmail+"%")
	}

	if filter.Status == "InActive" {

		filter.Status = "0"

	} else if filter.Status == "Active" {

		filter.Status = "1"

	}

	if filter.Status != "" {

		query = query.Where("tbl_jobs_applicants.status=?", filter.Status)

	}
	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&applicant)

		return applicant, 0, err

	}

	query.Find(&applicant).Count(&Totalapplicants)

	return applicant, Totalapplicants, nil
}
