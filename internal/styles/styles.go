package styles

import (
	lg "github.com/charmbracelet/lipgloss"
)

var ns = lg.NewStyle()

var Base = ns
var Question = ns.Foreground(lg.Color("#A855F7")).Bold(true)

var Selected = ns.Foreground(lg.Color("#39BAE6"))
var Correct = ns.Foreground(lg.Color("#13CE66"))
var Incorrect = ns.Foreground(lg.Color("#FF5C57"))
var IncorrectUnselected = ns.Foreground(lg.Color("#FF8F12"))

var Warn = ns.Foreground(lg.Color("#FF5C57"))
var Error = ns.Foreground(lg.Color("#FF5C57"))
var Info = ns.Foreground(lg.Color("#39BAE6"))
