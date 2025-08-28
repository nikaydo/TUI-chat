package ui

import "github.com/charmbracelet/lipgloss"

func splitText(name, text string, l int) string {
	lenName := len(name)
	var msg string = name + ": "
	var limit int = lenName + 2
	for _, i := range text {
		if limit == l-2 {
			limit = 0
			msg += " \n"
		}
		limit += 1
		msg += string(i)
	}
	return msg
}
func wrapMessage(msg string, width int) string {
	style := lipgloss.NewStyle().Width(width)
	return style.Render(msg)
}
