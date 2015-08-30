/* 
Repo is a parent class which implements generic CRUD operations for any resource. 
Any class that embed a repo will have a default working CRUD model with the following methods:

 *GetOne
 *GetAll
 *Delete
 *Update
 *Create

These methods can be overridden if need be and can be made to throw an error if no implementation is needed
*/

package repo

import(
	"fmt"
	"database/sql"
	"os"
	"time"
	"errors"
	
	_ "github.com/go-sql-driver/mysql"
	. "github.com/danmondy/conspire/models"
)
const(
	SQL_INSERT = "INSERT INTO %s (%s) VALUES (%s)"
	SQL_UPDATE = "UPDATE %s SET %s WHERE %s"
	SQL_DELETE = "DELETE FROM %s WHERE %s"
)

type Repo struct{
	db *sql.DB
}


func (*Repo) Insert(table string, obj interface{}) error{
	return errors.New("not implemented")
}

func (*Repo) GetOne(table string, id int) (interface{}, error){
	
	return nil, errors.New("not implemented")
}


func (*Repo) GetAll(where ...string) ([]interface{}, error){
	return nil, errors.New("not implemented")
}


func (*Repo) Delete(id int) (interface{}, error){
	return nil, errors.New("not implemented")
}


func NewRepo(db *sql.DB) *Repo{
	r := &Repo{db}
	
	err := r.RebuildDB() //comment this out when you don't want to wipe the db
	if err != nil{
		fmt.Println(err)
	}
	return r
}




func (r *Repo)RebuildDB()error{
	os.Remove("./sql/flog.db")
	
	//sqlite doesn't require autoincrement on id's (integer primary key auto assigns an unused random rowid if empty)	
	sqlStmt := "create table users (id integer not null primary key, email text, hashword text, rank integer, since text);"+
		"create table events (id integer not null primary key, title text, description text, long text, lat text, date text, communityId integer, creatorId integer, open integer);"+
		"create table eventUsers (eventId integer not null primary key, userId integer, status integer); "
		
	
	_, err := r.db.Exec(sqlStmt)
	if err != nil{
		return err
	}

	users := []User{NewUser("danmondy@gmail.com", "integrity", 5), NewUser("badguy@gmail.com", "dishonesty", 5)}

	tx, err := r.db.Begin()
	if err != nil{
		return err
	}

	stmt, err := tx.Prepare("insert into users(email, hashword, role, since) values(?,?,?,?)")
	if err != nil{
		return err
	}
	defer stmt.Close()
	for _, u := range users {
		_, err := stmt.Exec(u.Email, u.Hashword, u.Rank, TimeToString(u.Since))
		if err != nil{
			panic(err)
		}
	}

	stmt3, err := tx.Prepare("insert into users(email, hashword, role, since) values(?,?,?,?)")
	if err != nil{
		return err
	}
	defer stmt3.Close()
	for _, u := range users {
		_, err := stmt.Exec(u.Email, u.Hashword, u.Rank, TimeToString(u.Since))
		if err != nil{
			panic(err)
		}
	}


	stmt2, err := tx.Prepare("insert into projects(title, description, long, lat, date, creatorId, open) values(?,?,?,?,?,?,?,?)")
	if err != nil{
		return err
	}
	defer stmt2.Close()

	/*events := []Event{
		Event{Title:"Go to the city",Description:"Seriously go there", Long:"33.53", Lat:"-86.021", Date:time.Now(),Open:false},
		Event{Title:"Games at the tavern",Description:"Seriously go there", Long:"33.93", Lat:"-85.821", Date:time.Now(),Open:false},
		Event{Title:"Go to the desert",Description:"Seriously go there", Long:"33.23", Lat:"-85.921", Date:time.Now(),Open:false}}
	for _, e := range events {
		_, err := stmt2.Exec(e.Title, e.Description, e.Long, e.Lat, TimeToString(e.Date), 1234, 22, e.Open)
		if err != nil{
			panic(err)
		}
	}*/


	tx.Commit()

	return nil
}

func TimeToString(t time.Time) string{
	return t.Format(time.RFC3339)
}
func StringToTime(s string)(time.Time, error){
	return time.Parse(time.RFC3339, s)//returns both a time and an error so it can be returned directly.
}



//Generic query for temp use
//Query for any object, ...???then use the model's "MapRows()" method to get a slice of structs-maybe?
// func (r *Repo) Select(table string, whereClause string) (*sql.Rows, error){
// 	var rmaps []map[string]interface{}
// 	rows, err := r.db.Query(fmt.Sprintf("SELECT * FROM %v WHERE %v", table, whereClause))
// 	defer rows.Close()
// 	if err != nil{
// 		return nil, err
// 	}
	
// 	for r.Next(){
// 		vals := Make(map[string]interface{})
// 		err = r.MapScan(vals)
// 		if err != null{
// 			return nil, err
// 		}
// 		rmaps = append(rmaps, vals)
// 	}

	
// 	return rmaps, err//returning rows is unsafe, it keeps the connection open
// }
