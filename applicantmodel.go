package jobs

import (
	"gorm.io/gorm"
	"strings"
)

func (jobsmodel JobsModel) ApplicantsList(limit int, offset int, filter Filter, DB *gorm.DB) (applicant []TblJobsApplicants, Totalapplicants int64, err error) {

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "0"
	}

	query := DB.Debug().Table("tbl_jobs_applicants").Where("is_deleted = 0").Order("id desc")

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

//Create Applicant Function//

func (jobsmodel JobsModel) ApplicantCreate(applicant TblJobsApplicants, DB *gorm.DB) error {

	if err := DB.Table("tbl_jobs_applicants").Create(&applicant).Error; err != nil {

		return err
	}

	return nil
}
