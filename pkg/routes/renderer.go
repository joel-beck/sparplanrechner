package routes

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, templateName string, data interface{}, c echo.Context) error {
	err := t.Templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Error().Err(err).Msg("Error rendering template")
	}
	return err
}
