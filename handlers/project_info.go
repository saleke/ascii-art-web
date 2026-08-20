package handlers

import (
	"os"
	"strings"
	"time"
)

const (
	defaultPortfolioURL  = "https://saleke.vercel.app"
	defaultRepositoryURL = "https://github.com/saleke/ascii-art-web"
)

type ProjectInfo struct {
	Name          string
	Description   string
	Technologies  string
	DeveloperName string
	DeveloperRole string
	PortfolioURL  string
	RepositoryURL string
	CopyrightYear int
}

func projectInfo() ProjectInfo {
	return ProjectInfo{
		Name:          "ASCII Art Web",
		Description:   "A Go-based web application for generating ASCII art from user-provided text using configurable banner styles.",
		Technologies:  "Go · net/http · HTML · CSS",
		DeveloperName: "Aleke Emmanuel Solomon",
		DeveloperRole: "Software Developer in Training",
		PortfolioURL:  environmentValue("PORTFOLIO_URL", defaultPortfolioURL),
		RepositoryURL: environmentValue("GITHUB_REPOSITORY_URL", defaultRepositoryURL),
		CopyrightYear: time.Now().Year(),
	}
}

func environmentValue(name, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
}
