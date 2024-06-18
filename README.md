# Jobs Package

Our Jobs Template is built on NextJS and React, offering a modern and efficient solution for job listing and applicant management. This template is designed to seamlessly integrate with spurtCMS Admin, providing a robust platform for managing job-related content.


## Features

- Integration with spurtCMS Admin: Use the template seamlessly with spurtCMS Admin to manage job listings, applicant details, and other job-related content efficiently.
- Standalone Usage: The Jobs Template can also be used independently, allowing for flexibility in content management according to your specific needs.
- Dynamic Job Listings: Easily create, edit, and delete job listings, complete with detailed information about qualifications, salary, and more.
- Applicant Management: Track and manage applicant statuses, view applications, and streamline the recruitment process.



# Installation

``` bash
go get github.com/spurtcms/jobs
```
# Usage Example
``` bash
func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "Secret123",
		RoleId:     1,
		DB:         &gorm.DB{},
	})
	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, Auth.SecretKey)

	permisison, _ := Auth.IsGranted("Jobs", auth.CRUD)

	Jobs := jobs.JobsSetup(jobs.Config{

		DB:               &gorm.DB{},
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		//List Jobs
		jobslist, totalcount, err := Jobs.JobsList(10, 0, jobs.Filter{})

		fmt.Println(jobslist, totalcount, err)

		//Create new job
		cerr := Jobs.CreateJob(jobs.CreateJobReq{JobTitle: "Developer", JobDescription: "developing, coding, installing, and maintaining software systems.", JobType: "Full-Time", JobLocation: "Chennai"})

		if cerr != nil {

			fmt.Println(cerr)
		}
		//Update Job
		uerr := Jobs.UpdateJob(jobs.CreateJobReq{JobTitle: "Manager", JobDescription: "Develops, coordinates, and enforces systems, policies, procedures, and productivity", JobType: "Hybrid-Work", JobLocation: "Bangalore", Id: 1})

		if uerr != nil {

			fmt.Println(uerr)
		}
		//Delete job
		derr := Jobs.DeleteJob(1, 1)

		if derr != nil {

			fmt.Println(derr)
		}

	} else {

		fmt.Println("unauthroized")
	}

	apermisison, _ := Auth.IsGranted("Applicants", auth.CRUD)

		if apermisison {
	
			//List applicants
			applicantlist, totalcount, err := Jobs.ApplicantsList(10, 0, jobs.Filter{})
	
			fmt.Println(applicantlist, totalcount, err)
	       //Create applicant
			cerr := Jobs.CreateApplicant(jobs.CreateApplicantReq{Name: "Joe",Education: "BE",Experience: 2,JobType: "Full-Time"})
	
			if cerr != nil {
	
				fmt.Println(cerr)
			}
	       //Update applicant
			uerr := Jobs.UpdateApplicant(jobs.CreateApplicantReq{Name: "John",Education: "Bsc",Experience: 1,JobType: "Remote-Work"},1 )
	
			if uerr != nil {
	
				fmt.Println(uerr)
			}
	        //Delete applicant
			derr := Jobs.DeleteApplicant(1, 1)
	
			if derr != nil {
	
				fmt.Println(derr)
			}
	
		} else {
	
			fmt.Println("unauthroized")
		}
}



```




# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/spurtcms)] or you can create a new Issue in this repo[https://github.com/spurtcms/jobs/issues]. 
