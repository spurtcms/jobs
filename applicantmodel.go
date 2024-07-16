package jobs

import (
	"strings"

	"gorm.io/gorm"
)

func (jobsmodel JobsModel) ApplicantsList(limit int, offset int, filter Filter, DB *gorm.DB) (applicant []TblJobsApplicants, Totalapplicants int64, err error) {

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "0"
	}

	query := DB.Debug().Table("tbl_jobs_applicants").Where("is_deleted = 0").Order("id desc")

	if jobsmodel.Dataaccess == 1 {

		query = query.Where("tbl_jobs_applicants.created_by=?", jobsmodel.Userid)
	}

	if filter.Keyword != "" {

		if filter.Keyword == "0" || filter.Keyword == "1" {

			query = query.Where(" tbl_jobs_applicants.status=?", filter.Keyword)

		} else {

			query = query.Where("LOWER(TRIM(name)) ILIKE LOWER(TRIM(?))   OR LOWER(TRIM(email_id)) ILIKE LOWER(TRIM(?)) OR LOWER(TRIM(job_type)) ILIKE LOWER(TRIM(?)) or tbl_jobs_applicants.experience=? ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", filter.Keyword)

		}
	}

	if filter.Experience != 0 {

		query = query.Debug().Where("tbl_jobs_applicants.experience=?", filter.Experience)
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

//Getapplicant Details by Id//

func (jobsmodel JobsModel) GetApplicantById(id int, DB *gorm.DB) (applicant TblJobsApplicants, err error) {

	if err := DB.Model(TblJobsApplicants{}).Where("id=?", id).First(&applicant).Error; err != nil {

		return TblJobsApplicants{}, err
	}
	return applicant, nil
}

//Update Applicant//

func (jobsmodel JobsModel) ApplicantUpdate(applicant *TblJobsApplicants, DB *gorm.DB) error {

	query := DB.Model(TblJobsApplicants{}).Where("member_id=?", applicant.MemberId)

	if applicant.Image == "" && applicant.ImagePath == "" && applicant.Password == "" {

		query.Omit("password, image ,image_path").UpdateColumns(map[string]interface{}{"name": applicant.Name, "email_id": applicant.EmailId, "mobile_no": applicant.MobileNo, "job_type": applicant.JobType, "member_id": applicant.MemberId, "location": applicant.Location, "company_name": applicant.CompanyName, "experience": applicant.Experience, "education": applicant.Education, "graduation": applicant.Graduation, "skills": applicant.Skills, "current_salary": applicant.CurrentSalary, "expected_salary": applicant.ExpectedSalary, "status": applicant.Status, "modified_on": applicant.ModifiedOn})

	} else {

		query.UpdateColumns(map[string]interface{}{"name": applicant.Name, "email_id": applicant.EmailId, "mobile_no": applicant.MobileNo, "job_type": applicant.JobType, "member_id": applicant.MemberId, "password": applicant.Password, "location": applicant.Location, "company_name": applicant.CompanyName, "experience": applicant.Experience, "education": applicant.Education, "graduation": applicant.Graduation, "skills": applicant.Skills, "current_salary": applicant.CurrentSalary, "expected_salary": applicant.ExpectedSalary, "status": applicant.Status, "modified_on": applicant.ModifiedOn, "image": applicant.Image, "image_path": applicant.ImagePath, "storage_type": applicant.StorageType})
	}

	return nil
}

//Applicant Delete Function//

func (jobmodel JobsModel) ApplicantDelete(applicant TblJobsApplicants, DB *gorm.DB) error {

	if err := DB.Table("tbl_jobs_applicants").Where("member_id = ?", applicant.MemberId).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": applicant.DeletedOn, "deleted_by": applicant.DeletedBy}).Error; err != nil {
		return err
	}
	return nil

}

//Applicant jobs get by applicantId//

func (jobmodel JobsModel) GetApplicantJobs(applicantid int, limit int, offset int, DB *gorm.DB) (applicantjobs []TblJobsRegisters, Totaljobs int64, err error) {

	query := DB.Debug().Preload("JobList", func(db *gorm.DB) *gorm.DB {
		return db.Order("id asc")
	}).Table("tbl_jobs_registers").Select("tbl_jobs.id,tbl_jobs_registers.created_on,tbl_jobs_registers.status").Joins("inner join tbl_jobs on tbl_jobs.id = tbl_jobs_registers.job_id").Joins("inner join tbl_jobs_applicants on tbl_jobs_applicants.id =tbl_jobs_registers.applicant_id").Where("tbl_jobs_registers.applicant_id = ?", applicantid).Find(&applicantjobs)

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&applicantjobs)

		return applicantjobs, 0, err

	}

	query.Find(&applicantjobs).Count(&Totaljobs)

	return applicantjobs, Totaljobs, nil
}

//Multiselect Delete Function of Job//

func (jobsmodel JobsModel) MultiSelectedApplicantDelete(applicant *TblJobsApplicants, id []int, DB *gorm.DB) error {

	if err := DB.Model(TblJobsApplicants{}).Where("member_id in (?)", id).UpdateColumns(map[string]interface{}{"is_deleted": applicant.IsDeleted, "deleted_on": applicant.DeletedOn, "deleted_by": applicant.DeletedBy}).Error; err != nil {

		return err

	}

	return nil
}

//Multiselect  status change function of job//

func (jobsmodel JobsModel) MultiApplicantIsActive(applicant *TblJobsApplicants, jobid []int, status int, DB *gorm.DB) error {

	if err := DB.Model(TblJobsApplicants{}).Where("member_id in (?)", jobid).UpdateColumns(map[string]interface{}{"status": status, "modified_by": applicant.ModifiedBy, "modified_on": applicant.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}
