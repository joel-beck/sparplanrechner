package logging

import (
	"html/template"
	"os"

	"github.com/joel-beck/sparplanrechner/pkg/types"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger(level zerolog.Level) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05", NoColor: false}).With().Caller().Logger()
	zerolog.SetGlobalLevel(level)
}

func LogParsedTemplates(parsedTemplates *template.Template) {
	// log the parsed template file names
	for _, t := range parsedTemplates.Templates() {
		log.Debug().Str("template_name", t.Name()).Msg("Parsed template")
	}
}

func LogUserInputs(c echo.Context, inputs types.UserInputs) {
	log.Debug().Interface("user_inputs", inputs).Msg("User Inputs")
}

func LogResponseData(c echo.Context, responseData types.ResponseData) {
	log.Debug().Interface("response_data", responseData).Msg("Response Data")
}
