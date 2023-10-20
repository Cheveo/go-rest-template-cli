package rest_cli

import (
	"os"
	"path/filepath"

	"github.com/Cheveo/go-rest-cli/types"
	"github.com/Cheveo/go-rest-cli/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type standardProject struct{}

func (s *standardProject) Create(domain, modName, directory string) error {
	var resolvedUserPath string
	if directory != "" {
		userPath, _ := os.UserHomeDir()
		resolvedUserPath = userPath
	} else {
		resolvedUserPath = ""
	}

	d := types.DomainTmpl{
		Directory:         filepath.Join(resolvedUserPath, directory),
		Domain:            domain,
		CapitalizedDomain: cases.Title(language.English, cases.Compact).String(domain),
		GoMod:             modName,
	}

	err := util.CreateFileSkeleton(util.CreateMain(&d))
	err = util.CreateFileSkeleton(util.CreateGoMod(&d))
	err = util.CreateFileSkeleton(util.CreateHandler(&d))
	err = util.CreateFileSkeleton(util.CreateUtil(&d))
	err = util.CreateFileSkeleton(util.CreateService(&d))
	err = util.CreateFileSkeleton(util.CreateStorage(&d))
	err = util.CreateFileSkeleton(util.CreateModel(&d))
	err = util.CreateFileSkeleton(util.CreateServer(&d))

	return err
}

func NewStandardProject() *standardProject {
	return &standardProject{}
}