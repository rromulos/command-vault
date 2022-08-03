package command

import "fmt"

const (
	ColorDefault = "\x1b[39m"

	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[92m"
	ColorBlue    = "\x1b[94m"
	ColorGray    = "\x1b[90m"
	ColorYellow  = "\x1b[93m"
	ColorMagenta = "\x1b[95m"
	ColorCyan    = "\x1b[96m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}

func yellow(s string) string {
	return fmt.Sprintf("%s%s%s", ColorYellow, s, ColorDefault)
}

func magenta(s string) string {
	return fmt.Sprintf("%s%s%s", ColorMagenta, s, ColorDefault)
}

func cyan(s string) string {
	return fmt.Sprintf("%s%s%s", ColorCyan, s, ColorDefault)
}
