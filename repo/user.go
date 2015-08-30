package repo

import(
	
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"fmt"

	. "github.com/danmondy/conspire/models"
)
func (r *Repo) GetUserById(id int) (*User, error){
	query := "SELECT * FROM USERS WHERE Id = ?"
	row := r.db.QueryRow(query, id)
	return MapUser(row)
}
func (r *Repo) InsertUser(u *User)(error){

	_, err := r.db.Exec(fmt.Sprintf("INSERT into User (email, hashword, rank, since) VALUES (%v, %v, %v, %v)"), u.Email,u.Hashword,u.Rank,u.Since)
	if err != nil{
		return err
	}

	return nil
}
func (r *Repo) GetUserByEmail(email string) (*User, error){
	query := "SELECT * FROM USERS WHERE EMAIL = ?"
	row := r.db.QueryRow(query, email)
	return MapUser(row)
}

func MapUser(r *sql.Row)(*User, error){
	var u User
	var t string
	err := r.Scan(&u.Id, &u.Email, &u.Hashword, &u.Rank, &t)
	if err != nil{
		return nil, err
	}
	u.Since, err = StringToTime(t)
	return &u, err
} 

func MapUsers(r *sql.Rows)([]User, error){		
	var users []User
	for r.Next(){
		var u User
		err := r.Scan(&u.Id, &u.Email, &u.Hashword, &u.Rank, &u.Since)
		if err != nil{
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

