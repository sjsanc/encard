package styles

import (
	lg "github.com/charmbracelet/lipgloss"
)

var Question = lg.NewStyle().Foreground(lg.Color("#FF8F40")).Bold(true)

var CurrentChoice = lg.NewStyle().Foreground(lg.Color("#39BAE6"))
var Choice = lg.NewStyle()
var CorrectChoice = lg.NewStyle().Foreground(lg.Color("#13CE66"))
var IncorrectChoice = lg.NewStyle().Foreground(lg.Color("#FF5C57"))
var UnselectedChoice = lg.NewStyle().Foreground(lg.Color("#FF8F12"))

var Answer = lg.NewStyle().Foreground(lg.Color("#FF8F12"))
