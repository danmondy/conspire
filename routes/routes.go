package routes

import(
	"html/template"
	"net/http"
	"fmt"
	"log"

	"github.com/danmondy/conspire/repo"
) 

//Routes Helpers
var (
	templates *template.Template
	Repo *repo.Repo
)

func init(){
	log.Println("Routes initializer called")
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil{	
		fmt.Print(err)
	}
}
