package mysql

import (
	"log"
	"strconv"
	"time"
	"gorm.io/gorm"
)

type TblJobsApplicants struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:int"`
	MemberId    int       `gorm:"type:int"`
	Name        string    `gorm:"type:varchar (255)"`
	EmailId     string    `gorm:"type:varchar (255)"`
	MobileNo    string    `gorm:"type:varchar (255)"`
	JobType     string    `gorm:"type:varchar (255)"`
	Password    string    `gorm:"type:varchar(255)"`
	Location    string    `gorm:"type:varchar (255)"`
	Education   string    `gorm:"type:varchar (255)"`
	Graduation  int       `gorm:"type:int"`
	CompanyName string    `gorm:"type:varchar (255)"`
	Experience  string    `gorm:"type:varchar (255)"`
	Skills      string    `gorm:"type:varchar (255)"`
	ImagePath   string    `gorm:"type:varchar (255)"`
	Image       string    `gorm:"type:varchar (255)"`
	CreatedOn   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:int"`
	IsDeleted   int       `gorm:"type:int"`
	DeletedOn   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:int"`
	StorageType string     `gorm:"type:varchar (255)"`
}

type TblJobs struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:int"`
	CategoriesId   int       `gorm:"type:int"`
	JobTitle       string    `gorm:"type:varchar (255)"`
	JobSlug        string    `gorm:"type:varchar (255)"`
	JobDescription string    `gorm:"type:varchar (255)"`
	JobLocation    string    `gorm:"type:varchar (255)"`
	JobType        string    `gorm:"type:varchar (255)"`
	Education      string    `gorm:"type:varchar (255)"`
	Department     string    `gorm:"type:varchar (255)"`
	Experience     string    `gorm:"type:varchar (255)"`
	Salary         string    `gorm:"type:varchar (255)"`
	CreatedOn      time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy      int       `gorm:"type:int"`
	IsDeleted      int       `gorm:"type:int"`
	DeletedOn      time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy      int       `gorm:"DEFAULT:NULL"`
	Keywords       string    `gorm:"type:varchar (255)"`
	Skill          string    `gorm:"type:varchar (255)"`
	MinimumYears   int       `gorm:"type:int"`
	MaximumYears   int       `gorm:"type:int"`
	PostedDate     time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	ValidThrough   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	Status         int       `gorm:"type:int"`
	ModifiedOn     time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"type:int"`
}

type TblJobsEducation struct {
	Id        int    `gorm:"primaryKey;auto_increment;type:int"`
	Education string `gorm:"type:varchar (255)"`
}

type TblJobsDepartment struct {
	Id         int    `gorm:"primaryKey;auto_increment;type:int"`
	Department string `gorm:"type:varchar (255)"`
}

type TblJobsRegisters struct {
	Id          int               `gorm:"primaryKey;auto_increment;type:int"`
	JobId       TblJobsApplicants `gorm:"type:int;foreignKey:Id"`
	ApplicantId int               `gorm:"type:int"`
	CreatedBy   int               `gorm:"type:int"`
	CreatedOn   time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
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
