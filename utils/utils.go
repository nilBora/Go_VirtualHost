package utils
import (
    "fmt"
)

const(
    TEXT_COLOR_RED = "\033[91m"
    TEXT_COLOR_GREEN = "\033[92m"
    TEXT_COLOR_YELLOW = "\033[93m"
    TEXT_COLOR_BLUE = "\033[94m"
    TEXT_COLOR_MAGENTA = "\033[95m"
    TEXT_COLOR_CYAN = "\033[96m"
    TEXT_COLOT_WHITE = "\033[97m"
    TEXT_COLOR_GREY = "\033[90m"
    TEXT_COLOR_BLACK = "\033[30m"
    TEXT_STYLE_BOLD = "\033[1m"
    TEXT_STYLE_ITALIC = "\033[3m"
    TEXT_STYLE_UNDERLINE = "\033[4m"
    TEXT_END = "\033[0m"
)
func Info(msg string) {
    fmt.Println(ColorMessage(msg + " [ SUCCESS ]", TEXT_COLOR_GREEN))
    //return ColorMessage(msg + " [ SUCCESS ]", TEXT_COLOR_GREEN)
}

func Debug(msg string) {
    fmt.Println(ColorMessage(msg + " [ DEBUG ]", TEXT_COLOR_BLUE))
}

func Error(msg string) {
    fmt.Println(ColorMessage(msg + " [ ERROR ]", TEXT_COLOR_RED))
}

func FetchDebug(msg string) (string) {
    return ColorMessage(msg + " [ DEBUG ]", TEXT_COLOR_BLUE)
}

func FetchInfo(msg string) string {
    return ColorMessage(msg + " [ SUCCESS ]", TEXT_COLOR_GREEN)
}

func FetchError(msg string) (string) {
    return ColorMessage(msg + " [ ERROR ]", TEXT_COLOR_RED)
}

func ColorMessage(msg string, color string) string {
    return color + msg + TEXT_END;
}