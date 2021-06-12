package log

import (
    "Backend/tools/color"
    "log"
)

func Log(tag, message, tagc, msgc string) {
    log.Printf("%s%s%s: %s%s%s\n", tagc, tag, color.Reset, msgc, message, color.Reset)
}

func Warning(tag, message string) {
    Log(tag, message, color.Yellow, color.White)
}

func Error(tag, message string) {
    Log(tag, message, color.Red, color.White)
}

func Info(tag, message string) {
    Log(tag, message, color.Green, color.White)
}