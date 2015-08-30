package repo

import(
	//"errors"
	"database/sql"

	. "github.com/danmondy/conspire/models"
)

// func (r *Repo) GetUserById(id int) (*User, error){
// 	query := "SELECT * FROM USER WHERE Id = ?"
// 	row, err := r.db.Query(query, id)
// 	defer row.Close()
// 	if err != nil { return nil, err }
	
// 	return MapUser(row)
// }
// func (r *Repo) InsertUser(u *User)(error){

// 	_, err := r.db.Exec(fmt.Sprintf("INSERT into User (email, hashword, rank, since) VALUES (%v, %v, %v, %v)"), u.Email,u.Hashword,u.Rank,u.Since)
// 	if err != nil{
// 		return err
// 	}

// 	return nil
// }
func (r *Repo) GetAllProjects() ([]Project, error){
	/*query := "SELECT id, title, description, long, lat, start_date, end_date FROM EVENTS"
	rows, err := r.db.Query(query)
	if err != nil { return nil, err }
	defer rows.Close()
	events, err := MapEvents(rows)
	if err != nil {
		return nil, err
	}
	if len(events) == 0{
		return nil, errors.New("Query return 0 rows")
	}

	return events, err*/
	return nil, nil
}

func MapProjects(r *sql.Rows)([]Project, error){		
	/*var events []Event
	for r.Next(){
		var e Event
		var t string
		err := r.Scan(&e.Id, &e.Title, &e.Description, &e.Long, &e.Lat, &t, &e.Open)
		if err != nil{
			return nil, err
		}
		e.Date, err = StringToTime(t)
		if err != nil{
			return nil, err
		}
		events = append(events, e)
	}*/
	return nil, nil
}

