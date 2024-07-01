package jobs

import (
	"log"
	"time"

	"github.com/spurtcms/categories"
	"github.com/spurtcms/jobs/migration"
)

// JobsSetup used initialize Jobs configruation
func JobsSetup(config Config) *Jobs {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Jobs{
		AuthEnable:       config.AuthEnable,
		Permissions:      config.Permissions,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		DB:               config.DB,
		DataBaseType:     config.DataBaseType,
	}

}

// Jobs List Function
func (Ap *Jobs) JobsList(limit, offset int, filter Filter) (job []TblJobs, count int64, err error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return []TblJobs{}, 0, AuthErr
	}

	Jobsmodel.Dataaccess = Ap.Dataaccess

	Jobsmodel.Userid = Ap.Userid

	joblist, _, _ := Jobsmodel.JobsList(limit, offset, filter, Ap.DB)

	_, totalcount, _ := Jobsmodel.JobsList(0, 0, filter, Ap.DB)

	var Job []TblJobs

	for _, jobs := range joblist {

		child_page, _ := categories.Categorymodel.GetCategoryById(jobs.CategoriesId, Ap.DB)
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

				child, _ := categories.Categorymodel.GetCategoryById(flg, Ap.DB)
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

func (Ap *Jobs) CreateJob(Jc CreateJobReq) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr
	}

	var job TblJobs

	job.JobTitle = Jc.JobTitle

	job.CategoriesId = Jc.CategoriesId

	job.JobType = Jc.JobType

	job.JobLocation = Jc.JobLocation

	job.JobDescription = Jc.JobDescription

	job.Keywords = Jc.Keywords

	job.Education = Jc.Education

	job.Skill = Jc.Skill

	job.MinimumYears = Jc.MinimumYears

	job.MaximumYears = Jc.MaximumYears

	job.PostedDate = Jc.PostedDate

	job.ValidThrough = Jc.ValidThrough

	job.Status = Jc.Status

	job.Salary = Jc.Salary

	job.CreatedBy = Jc.CreatedBy

	job.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Jobsmodel.JobCreate(&job, Ap.DB)

	if err1 != nil {

		return err1
	}

	return nil
}

//Get JobDetailsById Function//

func (Ap *Jobs) GetJobById(id int) (job TblJobs, err error) {

	var jobdata TblJobs

	err = Jobsmodel.JobDetailsById(&jobdata, id, Ap.DB)

	layout := "2006-01-02T15:04"

	if !jobdata.PostedDate.IsZero() {

		jobdata.Startdate = jobdata.PostedDate.Format(layout)

	}
	if !jobdata.ValidThrough.IsZero() {

		jobdata.Enddate = jobdata.ValidThrough.Format(layout)

	}
	child_page, _ := categories.Categorymodel.GetCategoryById(jobdata.CategoriesId, Ap.DB)
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

			child, _ := categories.Categorymodel.GetCategoryById(flg, Ap.DB)
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

	jobdata.CategoryNames = reverseCategoryOrder

	if err != nil {

		return TblJobs{}, err
	}

	return jobdata, nil

}

//Update Function//

func (Ap *Jobs) UpdateJob(Jc CreateJobReq) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr
	}

	var Updatejob TblJobs

	Updatejob.Id = Jc.Id

	Updatejob.JobTitle = Jc.JobTitle

	Updatejob.CategoriesId = Jc.CategoriesId

	Updatejob.JobType = Jc.JobType

	Updatejob.JobLocation = Jc.JobLocation

	Updatejob.JobDescription = Jc.JobDescription

	Updatejob.Keywords = Jc.Keywords

	Updatejob.Education = Jc.Education

	Updatejob.Skill = Jc.Skill

	Updatejob.MinimumYears = Jc.MinimumYears

	Updatejob.MaximumYears = Jc.MaximumYears

	Updatejob.PostedDate = Jc.PostedDate

	Updatejob.ValidThrough = Jc.ValidThrough

	Updatejob.Status = Jc.Status

	Updatejob.Salary = Jc.Salary

	Updatejob.ModifiedBy = Jc.ModifiedBy

	Updatejob.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Jobsmodel.JobUpdate(&Updatejob, Ap.DB)

	if err1 != nil {

		return err1
	}

	return nil

}

//Job Delete Function//

func (Ap *Jobs) DeleteJob(id int, userid int) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr
	}

	var job TblJobs

	job.DeletedBy = userid

	job.Id = id

	job.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Jobsmodel.JobDelete(&job, Ap.DB)

	if err != nil {

		return err
	}

	return nil
}

//Multiselect Job delete function//

func (Ap *Jobs) MultiSelectedJobDelete(jobids []int, modifiedby int) (bool, error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {
		return false, AuthErr
	}

	var Jobs TblJobs
	Jobs.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	Jobs.DeletedBy = modifiedby
	Jobs.IsDeleted = 1

	err := Jobsmodel.MultiSelectedJobDelete(&Jobs, jobids, Ap.DB)

	if err != nil {

		return false, err
	}

	return true, nil
}

// multiselecte member status change
func (Ap *Jobs) MultiSelectJobsStatus(memberid []int, status int, modifiedby int) (bool, error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return false, AuthErr
	}

	var jobstatus TblJobs

	jobstatus.ModifiedBy = modifiedby

	jobstatus.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Jobsmodel.MultiJobsIsActive(&jobstatus, memberid, status, Ap.DB)

	if err != nil {

		return false, err
	}

	return true, nil

}

func (Ap *Jobs) GetJobApplicant(id int, limit, offset int, filter Filter) (app []TblJobs, Totalapplicants int64, err error) {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return []TblJobs{}, 0, AuthErr
	}

	applicant, _, err1 := Jobsmodel.GetJobApplicantByJobId(id, limit, offset, filter, Ap.DB)

	_, totalcount, _ := Jobsmodel.GetJobApplicantByJobId(id, 0, 0, filter, Ap.DB)

	if err1 != nil {
		log.Println(err)
	}
	return applicant, totalcount, nil
}

func (Ap *Jobs) ChangeApplicantStatus(jobid int, applicantid int, status string) error {

	if AuthErr := AuthandPermission(Ap); AuthErr != nil {

		return AuthErr

	}

	err := Jobsmodel.ChangeApplicantStatus(jobid, applicantid, status, Ap.DB)

	if err != nil {

		log.Println(err)
	}

	return nil

}

func (jobs *Jobs) GetJobsList(limit int, offset int, filter Filter) (jobsList []TblJobs, count int64, err error) {

	if AuthErr := AuthandPermission(jobs); AuthErr != nil {
		return []TblJobs{}, -1, AuthErr
	}

	jobsList, count, err = JobsModel.GetJobsList(JobsModel{}, limit, offset, filter, jobs.DB)
	if err != nil {

		return []TblJobs{}, -1, err
	}

	return jobsList, count, nil
}


func (jobs *Jobs) GetJobDetails(id int, jobSlug string) (jobDetail TblJobs, err error) {

	if AuthErr := AuthandPermission(jobs); AuthErr != nil {
		return TblJobs{}, AuthErr
	}

	jobDetail, err = JobsModel.GetJobDetails(JobsModel{}, id, jobSlug, jobs.DB)
	if err != nil {

		return TblJobs{}, err
	}

	return jobDetail, nil
}