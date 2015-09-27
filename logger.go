package logger

import "time"

const (
	configure string = "logger.conf.json"
)

//Confiuger is logger configure struc
type Confiuger struct {
	LogFile       string
	AccessLogFile string
	ErrlrLogFile  string
}

//Access is access log
type Access struct {
	IP   string
	Time string
}

// Error is error struct
type Error struct {
	Err     error
	Message string
	Time    time.Time
}

func readConf() {

}
