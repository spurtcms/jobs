package jobs

import (
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
	CreatedOn      time.Time
	CreatedBy      int
	IsDeleted      int
	DeletedOn      time.Time
	DeletedBy      int
	Keywords       string
	Skill          string
	MinimumYears   int
	MaximumYears   int
	PostedDate     time.Time
	ValidThrough   time.Time
	Status         int
	ModifiedOn     time.Time
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
