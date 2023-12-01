package mysql

import (
	"time"

	"github.com/Gorsonpy/catCafe/biz/model/membership"
	"gorm.io/gorm"
)

type Membership struct {
	CustomerID       int64
	Username         string
	ContactInfo      string
	Points           int64
	Level            string
	RegistrationDate time.Time
	Passwd           string
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
	return DB.Table("memberships").Create(mem).Error
}

func GetMembershipByUsername(username string) *Membership {
	mem := &Membership{}
	DB.Table("memberships").Where("username = ?", username).First(mem)
	return mem
}
