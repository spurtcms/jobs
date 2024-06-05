package jobs

import (
	"time"

	"github.com/spurtcms/categories"
)

type TblJobs struct {
	Id             int                      `gorm:"primaryKey;auto_increment;type:serial"`
	CategoriesId   int                      `gorm:"type:integer"`
	JobTitle       string                   `gorm:"type:character varying"`
	JobDescription string                   `gorm:"type:character varying"`
	JobLocation    string                   `gorm:"type:character varying"`
	JobType        string                   `gorm:"type:character varying"`
	Education      string                   `gorm:"type:character varying"`
	Department     string                   `gorm:"type:character varying"`
	Experience     string                   `gorm:"type:character varying"`
	Salary         string                   `gorm:"type:character varying"`
	CreatedOn      time.Time                `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy      int                      `gorm:"type:integer"`
	CreatedDate    string                   `gorm:"-:migration;<-:false"`
	IsDeleted      int                      `gorm:"type:integer"`
	DeletedOn      time.Time                `gorm:"DEFAULT:NULL"`
	DeletedBy      int                      `gorm:"DEFAULT:NULL"`
	ApplicantsList []TblJobsApplicants      `gorm:"foreignKey:Id;"`
	Keywords       string                   `gorm:"type:character varying"`
	Skill          string                   `gorm:"type:character varying"`
	MinimumYears   int                      `gorm:"type:integer"`
	MaximumYears   int                      `gorm:"type:integer"`
	PostedDate     time.Time                `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ValidThrough   time.Time                `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Status         int                      `gorm:"type:integer"`
	Startdate      string                   `gorm:"-:migration;<-:false"`
	Enddate        string                   `gorm:"-:migration;<-:false"`
	ModifiedOn     time.Time                `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy     int                      `gorm:"type:integer"`
	CategoryNames  []categories.TblCategories `gorm:"-"`
	JobList        []TblJobs                `gorm:"foreignKey:Id;"`
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

// Jobs List Function
func (jb *Jobs) JobsList(limit, offset int, filter Filter) (job []TblJobs, count int64, err error) {

	if AuthErr := AuthandPermission(jb); AuthErr != nil {

		return []TblJobs{}, 0, AuthErr
	}

	joblist, _, _ := Jobsmodel.JobsList(limit, offset, filter, jb.DB)

	_, totalcount, _ := Jobsmodel.JobsList(0, 0, filter, jb.DB)

	var Job []TblJobs

	for _, jobs := range joblist {

		jobs.CreatedDate = jobs.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		var child_page_Category categories.TblCategories

		_, child_page := categories.GetChildPageCategoriess(&child_page_Category, jobs.CategoriesId, jb.DB)

		var categorynames []categories.TblCategories

		var flg int

		categorynames = append(categorynames, child_page)

		flg = child_page.ParentId

		var count int

		if flg != 0 {

		CLOOP:

			for {

				count++

				if count >= 50 { // for safe

					break //for safe
				}

				var newchildcategory categories.TblCategories

				_, child := categories.GetChildPageCategoriess(&newchildcategory, flg, jb.DB)

				flg = child.ParentId

				if flg != 0 {

					categorynames = append(categorynames, child)

					goto CLOOP

				} else {

					categorynames = append(categorynames, child)

					break
				}

			}

		}

		var reverseCategoryOrder []categories.TblCategories

		for i := len(categorynames) - 1; i >= 0; i-- {

			reverseCategoryOrder = append(reverseCategoryOrder, categorynames[i])

		}

		jobs.CategoryNames = reverseCategoryOrder

		Job = append(Job, jobs)
	}
	return Job, totalcount, nil
}


//Job Create Function//


func (jb *Jobs)CreateJob(CreateJobReq)error{


	if AuthErr := AuthandPermission(jb); AuthErr != nil {

		return AuthErr
	}

	
}