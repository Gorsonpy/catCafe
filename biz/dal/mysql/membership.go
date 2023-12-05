package mysql

import (
	"time"

	"github.com/Gorsonpy/catCafe/biz/model/membership"
	"gorm.io/gorm"
)

type Membership struct {
	CustomerID       int64 `gorm:"primary_key"`
	Username         string
	ContactInfo      string
	Points           int64
	Level            string
	RegistrationDate time.Time
	Passwd           string
	IsAdmin          bool
}

func ExistUsername(username string) bool {
	mem := &Membership{}
	return DB.Table("memberships").Where("username = ?", username).First(mem).Error != gorm.ErrRecordNotFound
}
func AddMembership(m *membership.MembershipModel) error {
	mem := &Membership{}
	mem.Username = m.Username
	mem.Passwd = m.Passwd
	mem.RegistrationDate = time.Now()
	return DB.Table("memberships").Create(&mem).Error
}

func GetMembershipByUsername(username string) *Membership {
	mem := &Membership{}
	DB.Table("memberships").Where("username = ?", username).First(&mem)
	return mem
}

func GetMembershipById(userId int64) (*Membership, error) {
	mem := &Membership{}
	err := DB.Table("memberships").Where("customer_id = ?", userId).First(&mem).Error
	return mem, err
}

func UpdateMemPoint(memId int64, points int64) error {
	return DB.Table("memberships").Where("customer_id = ?", memId).Update("points", points).Error
}

func IsAdmin(userId int64) bool {
	mem := &Membership{}
	DB.Table("memberships").Where("customer_id = ?", userId).First(&mem)
	return mem.IsAdmin
}

func ListMem() ([]*Membership, error) {
	list := make([]*Membership, 0)
	err := DB.Table("memberships").Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list, err
}

func MemToModel(mem *Membership) *membership.MembershipModel {
	model := membership.NewMembershipModel()
	model.CustomerID = mem.CustomerID
	model.Username = mem.Username
	model.ContactInfo = mem.ContactInfo
	model.Points = mem.Points
	model.Level = mem.Level
	model.RegistrationDate = mem.RegistrationDate.Format(time.DateTime)
	model.IsAdmin = mem.IsAdmin
	return model
}
