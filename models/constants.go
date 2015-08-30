package models

const (
	IN_TEXTAREA = "textarea"
	IN_TEXT     = "text"
	IN_COMBO    = "combo"
	IN_HIDDEN   = "hidden"
	IN_DATE     = "date"
	IN_EMAIL    = "email"
	IN_CHECKBOX = "checkbox"
	IN_PASSWORD = "password"
)

type Scaffolder interface {
	GetScaffoldMap()
}
