package routes

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "strconv"
	"time"
	"log"

	. "github.com/danmondy/conspire/models"
)



//////Register Routes
func RegisterRoutes(r *mux.Router){
	log.Printf("registering routes...")
	r.HandleFunc("/cards/edit/{id}", editHandler)
	r.HandleFunc("/project/{project_id}", projectDetail)
	r.HandleFunc("/{username}/", profile)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout",logoutHandler)
	r.HandleFunc("/", indexHandler)

	//server static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	log.Printf("Routes registered.")
}

/////MIDDLEWARE
func groupAdminOnly(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc{	
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Let's pretend your a group admin")
		next(w, r)
	}
}



//////ROUTES
func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Index handler reached")
	model := &Page{Title: "Conspire - project management in one place."}
	
	//model.Page.AddWarning("You better check yourself.")
	//model.Page.AddAlert("Something pretty cool you should know.")
	//model.Page.AddMessage("Good work.")
	display(w, "index", model)
}

func projectDetail(w http.ResponseWriter, r *http.Request){
	
}
func profile(w http.ResponseWriter, r *http.Request){
	log.Println("Profile Handler hit")
	Repo.Insert("project", Project{Name:"Sarah Jenkins", CustomerEmail:"sj@gmail.com", Lon:"33.53", Lat:"-86.021"})
	w.Write([]byte("Check the log, Goof."))
}
func editHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	display(w, "product", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("%v -- %v", r.FormValue("email"), r.FormValue("password"))
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	
	//temp without db
	user, err := Repo.GetUserByEmail(email)
	if err != nil {
		page := &Page{}
		page.AddWarning(err.Error())
		http.Redirect(w, r, "/index?msg=UserNotFound", http.StatusFound)
		return
	}
	if err := user.CompareHash([]byte(password)); err != nil{
		page := &Page{}
		page.AddWarning(err.Error())
		indexHandler(w, r)
		return
	}
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "conspire-cookie", Value: email+"123456789", Expires: expiration}
	http.SetCookie(w, &cookie)
	
}
func logoutHandler(w http.ResponseWriter, r *http.Request){
	//not implemented
}

/*func getProjects() []Project { //this will be replaced with a database call
	return []Project{
		{1", "T-Shirt", "http://www.placecage.com/c/200/200", "This is the best tshirt you will ever own"},
		{2", "Socks", "http://www.fillmurray.com/200/200", "The best socks"},
		{3", "Hoodie", "http://www.stevensegallery.com/200/200", "Our favorite hoodie"},
		{4", "T-Shirt", "http://www.placecage.com/c/200/200", "This is the best tshirt you will ever own"},
		{5", "Socks", "http://www.fillmurray.com/200/200", "The best socks"},
		{6", "Hoodie", "http://www.stevensegallery.com/200/200", "Our favorite hoodie"}}
}*/

func handleErr(w http.ResponseWriter, err error, m string){
	if err != nil{
		fmt.Println(m)
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops! We have experienced an internal server error. Sorry for the inconvenience."))
	}
}


