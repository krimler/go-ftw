// Package go-ftw is a Framework for Testing Web Application Firewalls
// It is derived from the work made with the pytest plugin `ftw`
package main

import (
	"fmt"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/fzipi/go-ftw/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {
	const file_name string = "ftw.log"
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Critical: Invalid Log filei: ", file_name)
	}
	defer file.Close()

	log.Logger = log.Output(file)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Load timezone location based on TZ
	tzone, _ := time.Now().Zone()
	loc, err := time.LoadLocation(tzone)
	if err != nil {
		log.Info().Msgf("ftw/main: cannot load timezone")
	} else {
		time.Local = loc // -> set the global timezone
	}

	cmd.Execute(
		buildVersion(version, commit, date, builtBy),
	)

}

func buildVersion(version, commit, date, builtBy string) string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
