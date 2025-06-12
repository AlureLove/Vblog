package user

import (
	"Vblog/utils"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	v = validator.New()
)

type User struct {
	utils.ResourceMeta
	RegistryRequest
}

type RegistryRequest struct {
	Username string `json:"username" gorm:"column:username;unique;index" validate:"required"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)" validate:"required"`
	Profile
	Status
}

type Profile struct {
	Avatar   string `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(100)"`
	Email    string `json:"email" gorm:"column:email;type:varchar(100)"`
}

type Status struct {
	BlockReason string     `json:"block_reason" gorm:"column:block_reason;type:text"`
	BlockAt     *time.Time `json:"block_at" gorm:"column:block_at"`
}

func NewUser(req *RegistryRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("parameter validation failed: %s", err)
	}
	return &User{
		ResourceMeta:    *utils.NewResourceMeta(),
		RegistryRequest: *req,
	}, nil
}

func NewRegistryRequest() *RegistryRequest {
	return &RegistryRequest{}
}

func (s *Status) IsBlocked() bool {
	return s.BlockAt != nil
}

func (u *User) TableName() string {
	return "users"
}

func (r *RegistryRequest) Validate() error {
	return v.Struct(r)
}

func (u *User) String() string {
	return pretty.ToJSON(u)
}

func (r *RegistryRequest) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(password))
}
