package db

import (
	"open_im_sdk/pkg/utils"
)

func (d *DataBase) GetSubDepartmentList(departmentID string, args ...int) ([]*LocalDepartment, error) {
	d.mRWMutex.RLock()
	defer d.mRWMutex.RUnlock()
	var departmentList []LocalDepartment
	var err error
	sql := d.conn.Where("parent_id = ? ", departmentID).Order("order_department DESC")
	if len(args) == 2 {
		offset := args[0]
		count := args[1]
		err = sql.Offset(offset).Limit(count).Find(&departmentList).Error
	} else {
		err = sql.Find(&departmentList).Error
	}
	var transfer []*LocalDepartment
	for _, v := range departmentList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, utils.Wrap(err, utils.GetSelfFuncName()+" failed")
}

func (d *DataBase) InsertDepartment(department *LocalDepartment) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	return utils.Wrap(d.conn.Create(department).Error, "InsertDepartment failed")
}

func (d *DataBase) UpdateDepartment(department *LocalDepartment) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	return utils.Wrap(d.conn.Model(department).Select("*").Updates(*department).Error, "UpdateDepartment failed")
}

func (d *DataBase) DeleteDepartment(departmentID string) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	local := LocalDepartment{DepartmentID: departmentID}
	return utils.Wrap(d.conn.Delete(&local).Error, "DeleteDepartment failed")
}

func (d *DataBase) GetDepartmentInfo(departmentID string) (*LocalDepartment, error) {
	d.mRWMutex.RLock()
	defer d.mRWMutex.RUnlock()
	var local LocalDepartment
	return &local, utils.Wrap(d.conn.Where("department_id=?", departmentID).First(&local).Error, "GetDepartmentInfo failed")
}

func (d *DataBase) GetAllDepartmentList() ([]*LocalDepartment, error) {
	d.mRWMutex.RLock()
	defer d.mRWMutex.RUnlock()
	var departmentList []LocalDepartment
	d.conn.Debug()
	//	err := d.conn.Order("order DESC").Find(&departmentList).Error
	err := d.conn.Order("order_department DESC").Find(&departmentList).Error
	var transfer []*LocalDepartment
	for _, v := range departmentList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, utils.Wrap(err, utils.GetSelfFuncName()+" failed")
}
