package styles

import (
	lg "github.com/charmbracelet/lipgloss"
)

var Input lg.Style = lg.NewStyle().Width(25).Height(14).Align(lg.Center).BorderStyle(lg.RoundedBorder()).BorderForeground(lg.Color("3"))
var Output lg.Style = lg.NewStyle().Width(70).Inherit(Input)
var Heading lg.Style = lg.NewStyle().Background(lg.Color("9"))
