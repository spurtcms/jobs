package postgres

import (
	"log"
	"strconv"
	"time"
	"gorm.io/gorm"
)

type TblJobsApplicants struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId       int       `gorm:"type:integer"`
	Name           string    `gorm:"type:character varying"`
	EmailId        string    `gorm:"type:character varying"`
	MobileNo       string    `gorm:"type:character varying"`
	JobType        string    `gorm:"type:character varying"`
	Password       string    `gorm:"type:character varying"`
	Location       string    `gorm:"type:character varying"`
	Education      string    `gorm:"type:character varying"`
	Graduation     int       `gorm:"type:integer"`
	CompanyName    string    `gorm:"type:character varying"`
	Experience     int       `gorm:"type:integer"`
	Skills         string    `gorm:"type:character varying"`
	ImagePath      string    `gorm:"type:character varying"`
	Image          string    `gorm:"type:character varying"`
	CreatedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy      int       `gorm:"type:integer"`
	ModifiedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"type:integer"`
	IsDeleted      int       `gorm:"type:integer"`
	DeletedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy      int       `gorm:"type:integer"`
	CurrentSalary  int       `gorm:"type:integer"`
	ExpectedSalary int       `gorm:"type:integer"`
	Status         int       `gorm:"type:integer"`
}

type TblJobs struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	CategoriesId   int       `gorm:"type:integer"`
	JobTitle       string    `gorm:"type:character varying"`
	JobSlug        string    `gorm:"type:character varying"`
	JobDescription string    `gorm:"type:character varying"`
	JobLocation    string    `gorm:"type:character varying"`
	JobType        string    `gorm:"type:character varying"`
	Education      string    `gorm:"type:character varying"`
	Department     string    `gorm:"type:character varying"`
	Experience     string    `gorm:"type:character varying"`
	Salary         string    `gorm:"type:character varying"`
	CreatedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy      int       `gorm:"type:integer"`
	IsDeleted      int       `gorm:"type:integer"`
	DeletedOn      time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy      int       `gorm:"DEFAULT:NULL"`
	Keywords       string    `gorm:"type:character varying"`
	Skill          string    `gorm:"type:character varying"`
	MinimumYears   int       `gorm:"type:integer"`
	MaximumYears   int       `gorm:"type:integer"`
	PostedDate     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ValidThrough   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Status         int       `gorm:"type:integer"`
	ModifiedOn     time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"type:integer"`
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
	Id          int `gorm:"primaryKey;auto_increment;type:serial"`
	JobId       int
	ApplicantId int
	CreatedBy   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

func MigrationTables(db *gorm.DB) {

	err := db.AutoMigrate(

		TblJobs{},
		TblJobsApplicants{},
		TblJobsDepartment{},
		TblJobsEducation{},
		TblJobsRegisters{},
	)

	if err != nil {

		panic(err)
	}

	var permissionmaxID int

	if err := db.Debug().Raw("SELECT max(id) FROM tbl_modules").Row().Scan(&permissionmaxID); err != nil {
		// Handle error
		log.Println(err)
	}

	// permissionidstr := strconv.Itoa(permissionmaxID + 1)

	// permissionidstr1 := strconv.Itoa(permissionmaxID + 2)

	// permissionidstr2 := strconv.Itoa(permissionmaxID + 3)

	var permissionvaluemaxID int

	if err := db.Debug().Raw("SELECT max(id) FROM tbl_module_permissions").Row().Scan(&permissionvaluemaxID); err != nil {
		// Handle error
		log.Println(err)
	}

	permissionvaluemaxIDstr := strconv.Itoa(permissionvaluemaxID + 1)

	permissionvaluemaxIDstr1 := strconv.Itoa(permissionvaluemaxID + 2)

	permissionvaluemaxIDstr2 := strconv.Itoa(permissionvaluemaxID + 3)

	db.Exec(`INSERT INTO tbl_modules(id,module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) values (39, 'Jobs',1,1,'2023-03-14 11:09:12',0,0,0,'/public/img/job.svg','',1,'left',1,1 )`)

	db.Exec(`INSERT INTO tbl_modules(id,module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) values (40, 'Jobs',1,1,'2023-03-14 11:09:12',0,39,0,'/public/img/job-tab.svg','',2,'tab',1,0 )`)

	db.Exec(`INSERT INTO tbl_modules(id,module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) values (41, 'Applicants',1,1,'2023-03-14 11:09:12',0,39,0,'/public/img/group.svg','',3,'tab',1,0 )`)

	db.Exec(`INSERT INTO tbl_modules(id,module_name, is_active, created_by, created_on, default_module, parent_id, assign_permission, icon_path, description, order_index, menu_type,full_access_permission,group_flg) values (42, 'Settings',1,1,'2023-03-14 11:09:12',0,39,0,'/public/img/settings-tab.svg','',4,'tab',1,0 )`)

	/*module permission*/

	db.Exec(`INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name)
	SELECT * FROM (
		SELECT 
			` + permissionvaluemaxIDstr + ` AS id,
			'/jobs/' AS route_name, 
			'Jobs' AS display_name, 
			'' AS description, 
			40 as module_id,
			1 AS created_by,
			'2023-03-14 11:09:12'::timestamp AS created_on, 
			1 AS full_access_permission, 
			0 AS parent_id, 
			1 AS assign_permission, 
			1 AS order_index, 
			'Jobs' AS slug_name
	) AS temp 
	WHERE NOT EXISTS (
		SELECT * 
		FROM tbl_module_permissions 
		WHERE route_name = '/jobs/'
	);`)

	db.Exec(`INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name)
	SELECT * FROM (
		SELECT 
			` + permissionvaluemaxIDstr1 + ` AS id,
			'/jobs/applicants/' AS route_name, 
			'Applicants' AS display_name, 
			'' AS description, 
			41 as module_id,
			1 AS created_by,
			'2023-03-14 11:09:12'::timestamp AS created_on, 
			1 AS full_access_permission, 
			0 AS parent_id, 
			1 AS assign_permission, 
			2 AS order_index, 
			'Applicants' AS slug_name
	) AS temp 
	WHERE NOT EXISTS (
		SELECT * 
		FROM tbl_module_permissions 
		WHERE route_name = '/jobs/applicants/'
	);`)

	db.Exec(`INSERT INTO tbl_module_permissions(id,route_name, display_name, description, module_id, created_by, created_on, full_access_permission, parent_id, assign_permission,order_index, slug_name)
	SELECT * FROM (
		SELECT 
			` + permissionvaluemaxIDstr2 + ` AS id,
			'/jobs/settings/' AS route_name, 
			'settings' AS display_name, 
			'' AS description, 
			42 as module_id,
			1 AS created_by,
			'2023-03-14 11:09:12'::timestamp AS created_on, 
			1 AS full_access_permission, 
			0 AS parent_id, 
			1 AS assign_permission, 
			2 AS order_index, 
			'settings' AS slug_name
	) AS temp 
	WHERE NOT EXISTS (
		SELECT * 
		FROM tbl_module_permissions 
		WHERE route_name = '/jobs/settings/'
	);`)

}
