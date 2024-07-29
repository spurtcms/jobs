package jobs

import (
	"strings"
	"time"

	"github.com/spurtcms/categories"

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
type JobsListReq struct {
	Limit  int
	Offset int
}
type Filter struct {
	Keyword        string
	JobType        string
	Name           string
	EmailId        string
	MobileNo       string
	JobId          int
	Status         string
	ApplicantName  string
	ApplicantEmail string
	Experience     int
	JobTitle       string
	JobLocation    string
	CategoryId     int
	CategorySlug   string
	KeyWord        string
	MinimumYears   int
	MaximumYears   int
	DatePosted     string
	Skill          string
}

type TblJobs struct {
	Id             int                        `gorm:"primaryKey;auto_increment;type:serial"`
	CategoriesId   int                        `gorm:"type:integer"`
	Category       categories.TblCategories   `gorm:"foreignKey:Id;"`
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
	JobList        []TblJobsRegisters         `gorm:"foreignKey:Id;"`
	Jobregstatus   string                     `gorm:"-:migration;<-:false"`
	JobSlug        string                     `gorm:"type:character varying"`
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
	CreatedDate    string              `gorm:"-:migration;<-:false"`
	Status         string              `gorm:"type:character varying"`
}

func (jobsmodel JobsModel) JobsList(limit int, offset int, filter Filter, DB *gorm.DB) (job []TblJobs, Totaljob int64, err error) {

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "0"
	}

	query := DB.Debug().Table("tbl_jobs").Where("is_deleted = 0").Order("id desc")

	if jobsmodel.Dataaccess == 1 {

		query = query.Where("tbl_jobs.created_by=?", jobsmodel.Userid)
	}

	if filter.Keyword != "" {

		if filter.Keyword == "0" || filter.Keyword == "1" {

			query = query.Where(" tbl_jobs.status=?", filter.Keyword)

		} else {

			query = query.Where(" LOWER(TRIM(job_title)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(job_type)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}

	}

	if filter.JobType != "" {

		query = query.Debug().Where("tbl_jobs.job_type=?", filter.JobType)
	}

	if filter.JobTitle != "" {

		query = query.Debug().Where("LOWER(TRIM(job_title)) LIKE LOWER(TRIM(?))", "%"+filter.JobTitle+"%")
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

	if err := DB.Model(TblJobs{}).Where("id = ?", job.Id).UpdateColumns(map[string]interface{}{"job_title": job.JobTitle, "categories_id": job.CategoriesId, "job_type": job.JobType, "job_location": job.JobLocation, "job_description": job.JobDescription, "skill": job.Skill, "salary": job.Salary, "education": job.Education, "minimum_years": job.MinimumYears, "maximum_years": job.MaximumYears, "status": job.Status, "posted_date": job.PostedDate, "valid_through": job.ValidThrough, "keywords": job.Keywords, "modified_on": job.ModifiedOn, "modified_by": job.ModifiedBy, "job_slug": job.JobSlug}).Error; err != nil {
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

	if err := DB.Model(TblJobs{}).Where("id in (?)", jobid).UpdateColumns(map[string]interface{}{"status": status, "modified_by": job.ModifiedBy, "modified_on": job.ModifiedOn}).Error; err != nil {

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
	}).Table("tbl_jobs_registers").Select("tbl_jobs_applicants.id,tbl_jobs_registers.status as Jobregstatus ").Joins("inner join tbl_jobs on tbl_jobs.id = tbl_jobs_registers.job_id").Joins("inner join tbl_jobs_applicants on tbl_jobs_applicants.id =tbl_jobs_registers.applicant_id").Where("tbl_jobs.id = ?", id).Order("tbl_jobs_applicants.id desc").Find(&applicant)

	if filter.Keyword != "" {

		if filter.Keyword == "0" || filter.Keyword == "1" {

			query = query.Where(" tbl_jobs_applicants.status=?", filter.Keyword)

		} else {

			query = query.Where("LOWER(TRIM(name)) LIKE LOWER(TRIM(?))   OR LOWER(TRIM(email_id)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(tbl_jobs_applicants.job_type)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}
	}

	if filter.JobType != "" {

		query = query.Where("tbl_jobs_applicants.job_type=?", filter.JobType)
	}

	if filter.Experience != 0 {

		query = query.Where("tbl_jobs_applicants.experience=?", filter.Experience)
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

func (jobsmodel JobsModel) ChangeApplicantStatus(jobid int, applicantid int, status string, DB *gorm.DB) error {

	if err := DB.Model(TblJobsRegisters{}).Where("job_id=? And applicant_id=?", jobid, applicantid).UpdateColumns(map[string]interface{}{"status": status}).Error; err != nil {

		return err
	}

	return nil
}
func (jobsModel JobsModel) GetJobsList(limit int, offset int, filter Filter, DB *gorm.DB) (jobsList []TblJobs, count int64, err error) {

	listQuery := DB.Debug().Table("tbl_jobs").Select("tbl_jobs.*,tbl_categories.id as CatId,tbl_categories.category_name,tbl_categories.category_slug").Joins("inner join tbl_categories on tbl_jobs.categories_id = tbl_categories.id").Where("tbl_jobs.is_deleted = 0 AND tbl_jobs.status = 1").Preload("Category")

	if filter.JobTitle != "" {

		listQuery = listQuery.Where("job_title = ?", filter.JobTitle)
	}

	if filter.KeyWord != "" {

		listQuery = listQuery.Where("LOWER(TRIM(job_title)) like LOWER(TRIM(?))", "%"+filter.KeyWord+"%")
	}

	if filter.JobLocation != "" {

		listQuery = listQuery.Where("LOWER(TRIM(job_location)) = LOWER(TRIM(?))", filter.JobLocation)
	}

	if filter.CategorySlug != "" {

		listQuery = listQuery.Where("tbl_categories.category_slug = ?", filter.CategorySlug)
	}

	if filter.CategoryId != 0 {

		listQuery = listQuery.Where("categories_id = ?", filter.CategoryId)
	}

	if filter.Skill != "" {
		listQuery = listQuery.Where("skill = ?", filter.Skill)
	}

	if filter.MinimumYears != 0 && filter.MaximumYears != 0 {

		listQuery = listQuery.Where("minimum_years >= ? and maximum_years <= ?", filter.MinimumYears, filter.MaximumYears)

	} else if filter.MinimumYears != 0 {

		listQuery = listQuery.Where("minimum_years = ? or maximum_years = ? or (minimum_years > ? and maximum_years < ?) ", filter.MinimumYears, filter.MinimumYears, filter.MinimumYears, filter.MinimumYears)

	} else if filter.MaximumYears != 0 {

		listQuery = listQuery.Where("maximum_years >= ?", filter.MaximumYears)
	}

	if filter.DatePosted != "" {

		var startDate, endDate time.Time

		var currentDate = time.Now().Local()

		if filter.DatePosted == "This Week" {

			currentDay := time.Now().Local().Weekday().String()

			switch currentDay {

			case "Monday":
				startDate = currentDate
				endDate = currentDate.AddDate(0, 0, 6)

			case "Tuesday":
				startDate = currentDate.AddDate(0, 0, -1)
				endDate = currentDate.AddDate(0, 0, 5)

			case "Wednesday":
				startDate = currentDate.AddDate(0, 0, -2)
				endDate = currentDate.AddDate(0, 0, 4)

			case "Thursday":
				startDate = currentDate.AddDate(0, 0, -3)
				endDate = currentDate.AddDate(0, 0, 3)

			case "Friday":
				startDate = currentDate.AddDate(0, 0, -4)
				endDate = currentDate.AddDate(0, 0, 2)

			case "Saturday":
				startDate = currentDate.AddDate(0, 0, -5)
				endDate = currentDate.AddDate(0, 0, 1)

			case "Sunday":
				startDate = currentDate.AddDate(0, 0, -6)
				endDate = currentDate
			}

		}

		if filter.DatePosted == "This Month" {

			startDate = time.Date(currentDate.Year(), currentDate.Month(), 1, 0, 0, 0, 0, currentDate.Location())
			firstDayOfNxtMnth := startDate.AddDate(0, 1, 0)
			endDate = firstDayOfNxtMnth.Add(-time.Second)
		}

		if filter.DatePosted == "This Year" {

			startDate = time.Date(currentDate.Year(), time.January, 1, 0, 0, 0, 0, currentDate.Location())
			startofNxtYear := startDate.AddDate(1, 0, 0)
			endDate = startofNxtYear.Add(-time.Second)
		}

		if filter.DatePosted == "Today" {

			startDate = time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location())
			nxtDay := startDate.AddDate(0, 0, 1)
			endDate = nxtDay.Add(-time.Second)
		}

		listQuery = listQuery.Where("posted_date between (?) and (?)", startDate, endDate)

	}

	listQuery = listQuery.Limit(limit).Offset(offset).Order("tbl_jobs.id desc").Find(&jobsList)

	if listQuery.Error != nil {
		return []TblJobs{}, -1, listQuery.Error

	}

	if len(jobsList) <= 0 {
		return []TblJobs{}, -1, gorm.ErrRecordNotFound

	}

	countQuery := listQuery.Count(&count)

	if countQuery.Error != nil {

		return []TblJobs{}, -1, countQuery.Error
	}

	return jobsList, count, nil
}

func (jobsModel JobsModel) GetJobDetails(id int, jobSlug string, DB *gorm.DB) (jobDetail TblJobs, err error) {

	query := DB.Debug().Table("tbl_jobs").Select("tbl_jobs.*,tbl_categories.id as CatId,tbl_categories.category_name,tbl_categories.category_slug").Joins("inner join tbl_categories on tbl_jobs.categories_id = tbl_categories.id").Where("tbl_jobs.is_deleted = 0").Preload("Category")

	if id != 0 {

		query = query.Where("tbl_jobs.id = ?", id)
	} else if jobSlug != "" {

		query = query.Where("tbl_jobs.job_slug = ? ", jobSlug)
	}

	query = query.Find(&jobDetail)

	if query.Error != nil {

		return TblJobs{}, err
	}

	return jobDetail, nil
}
