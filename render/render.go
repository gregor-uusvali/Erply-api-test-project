package render

import (
	"Erply-api-test-project/models"
	"fmt"
	"net/http"
	"text/template"
)

// render.Template renders templates using html/template
func Template(w http.ResponseWriter, r *http.Request, templateName string, tmplData *models.TemplateData) error {
	parsedTemplate, err := template.ParseFiles("./templates/"+templateName, "./templates/base.html")
	if err != nil {
		fmt.Println("Error parsing template: ", err)
		return err
	}

	err = parsedTemplate.Execute(w, *tmplData)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		return err
	}
	return nil
}

