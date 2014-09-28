package logger

import (
    "io/ioutil"
    "os"
    "log"
    "time"
)

var (
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

var responseCode		int

func Init(log_level string) {

    traceHandle := ioutil.Discard
    infoHandle := ioutil.Discard
    warningHandle := ioutil.Discard
    errorHandle := os.Stderr

    if (log_level == "trace") {
	traceHandle = os.Stdout
	infoHandle = os.Stdout
	warningHandle = os.Stdout
    } else if (log_level == "info") {
	infoHandle = os.Stdout
	warningHandle = os.Stdout
    } else if (log_level == "warn") {
	warningHandle = os.Stdout
    }

    Trace = log.New(traceHandle,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)

	responseCode = 200
}

func  SetResponseCode(code int) {
	responseCode = code
}

func Elapsed(start time.Time, line string) {
	elapsed := time.Since(start)
	Info.Println(line, " dur:", int64(elapsed/time.Millisecond), "ms", " response: ", responseCode)
}
