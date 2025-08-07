package model

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// --- constants / helpers ----------------------------------------------------

const (
	UserTypeDriver = "driver"
	UserTypeRider  = "rider"
)

func ValidUserType(t string) bool { return t == UserTypeDriver || t == UserTypeRider }

// --- model -------------------------------------------------------------------

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	FirstName            string `gorm:"size:100;not null"            json:"first_name"`
	LastName             string `gorm:"size:100;not null"            json:"last_name"`
	Email                string `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PasswordHash         string `gorm:"size:60;not null"             json:"-"`
	Type                 string `gorm:"size:20;index;not null"       json:"type"`
	DriversLicenseNumber string `gorm:"size:100"                     json:"drivers_license_number,omitempty"`
}

// TableName keeps us on "users" (opcional; gorm ya lo hace bien)
func (User) TableName() string { return "users" }

// --- business helpers --------------------------------------------------------

func (u User) IsDriver() bool { return u.Type == UserTypeDriver }
func (u User) IsRider() bool  { return u.Type == UserTypeRider }

func (u User) DisplayName() string {
	if len(u.LastName) == 0 {
		return strings.Title(u.FirstName)
	}
	return strings.Title(u.FirstName) + " " + strings.ToUpper(string(u.LastName[0])) + "."
}

// --- password API ------------------------------------------------------------

// SetPassword genera el hash bcrypt y lo guarda en el struct.
func (u *User) SetPassword(raw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword compara el password plano con el hash.
func (u User) CheckPassword(raw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(raw)) == nil
}

// --- gorm hooks --------------------------------------------------------------

// Simple “validaciones” al estilo Rails.
// Si necesitas algo más completo, puedes meter validator.v10 en la capa servicio.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if !ValidUserType(u.Type) {
		return fmt.Errorf("invalid user type: %s", u.Type)
	}
	if len(u.PasswordHash) == 0 {
		return fmt.Errorf("password hash cannot be blank")
	}
	return nil
}
