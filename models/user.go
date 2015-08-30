package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct{
	Id int
	Email string
	Hashword string
	Rank int
	Since time.Time
}
func NewUser(email string, password string, rank int)User{
	u := User{Email: email, Rank: rank}
	u.SetHashword([]byte(password))
	u.Since = time.Now()
	return u
}

func (u *User) SetHashword(password []byte)error{
	hash, err := bcrypt.GenerateFromPassword(password, 0)//0 -> uses default cost (10 at time of writing, 4 min / 31 max)
	if err != nil {
		return err
	}else{ 
		u.Hashword = string(hash)
		return nil
	}
}
func (u *User) CompareHash(p []byte)error{
	return bcrypt.CompareHashAndPassword([]byte(u.Hashword), p)
}

