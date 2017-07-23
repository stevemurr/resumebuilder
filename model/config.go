package model

import (
	"resumebuilder/section/education"
	"resumebuilder/section/experience"
	"resumebuilder/section/extracarricular"
	"resumebuilder/section/personal"
	"resumebuilder/section/settings"
	"resumebuilder/section/skills"
)

// Config represents the resume toml file.
type Config struct {
	Settings        settings.Settings
	Personal        personal.Personal
	Skills          skills.Skills
	Experience      map[string]experience.Experience
	Education       map[string]education.Education
	Extracarricular map[string]extracarricular.Extracarricular
}
