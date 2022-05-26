package db

import (
	"open_im_sdk/pkg/utils"
)

func (d *DataBase) GetDepartmentMemberListByDepartmentID(departmentID string, args ...int) ([]*LocalDepartmentMember, error) {
	d.mRWMutex.RLock()
	defer d.mRWMutex.RUnlock()
	var departmentMemberList []LocalDepartmentMember
	var err error
	sql := d.conn.Where("department_id = ? ", departmentID).Order("order_member DESC")
	if len(args) == 2 {
		offset := args[0]
		count := args[1]
		err = sql.Offset(offset).Limit(count).Find(&departmentMemberList).Error
	} else {
		err = sql.Find(&departmentMemberList).Error
	}
	var transfer []*LocalDepartmentMember
	for _, v := range departmentMemberList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, utils.Wrap(err, utils.GetSelfFuncName()+" failed")
}

func (d *DataBase) GetAllDepartmentMemberList() ([]*LocalDepartmentMember, error) {
	d.mRWMutex.RLock()
	defer d.mRWMutex.RUnlock()
	var departmentMemberList []LocalDepartmentMember
	err := d.conn.Find(&departmentMemberList).Error

	var transfer []*LocalDepartmentMember
	for _, v := range departmentMemberList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, utils.Wrap(err, utils.GetSelfFuncName()+" failed")
}

func (d *DataBase) InsertDepartmentMember(departmentMember *LocalDepartmentMember) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	return utils.Wrap(d.conn.Create(departmentMember).Error, "InsertDepartmentMember failed")
}

func (d *DataBase) UpdateDepartmentMember(departmentMember *LocalDepartmentMember) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	return utils.Wrap(d.conn.Model(departmentMember).Select("*").Updates(*departmentMember).Error, "UpdateDepartmentMember failed")
}

func (d *DataBase) DeleteDepartmentMember(departmentID string, userID string) error {
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()
	//local := LocalDepartmentMember{DepartmentID: departmentID, UserID: userID}
	return utils.Wrap(d.conn.Where("department_id = ? and user_id = ?", departmentID, userID).Delete(&LocalDepartmentMember{}).Error, "DeleteDepartmentMember failed")
}

func (d *DataBase) GetDepartmentMemberListByUserID(userID string) ([]*LocalDepartmentMember, error) {
	d.mRWMutex.RLock()
	defer d.mRWMutex.RUnlock()
	var departmentMemberList []LocalDepartmentMember
	err := d.conn.Where("user_id = ? ", userID).Find(&departmentMemberList).Error
	var transfer []*LocalDepartmentMember
	for _, v := range departmentMemberList {
		v1 := v
		transfer = append(transfer, &v1)
	}
	return transfer, utils.Wrap(err, utils.GetSelfFuncName()+" failed")
}
