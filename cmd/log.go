package cmd

import (
	"encoding/json"
	"errors"
	"github.com/somenave/eventsCalendar/storage"
	"sync"
	"time"
)

var mu sync.Mutex

type LogType string

const (
	LogInput        LogType = "input"
	LogOutput       LogType = "output"
	LogError        LogType = "error"
	LogNotification LogType = "notification"
)

func validateLogType(logType LogType) error {
	switch logType {
	case LogInput, LogOutput, LogError, LogNotification:
		return nil
	default:
		return errors.New("invalid log type")
	}
}

type Logger struct {
	logs    []Log
	storage storage.Store
}

func NewLogger(s storage.Store) *Logger {
	return &Logger{
		logs:    []Log{},
		storage: s,
	}
}

type Log struct {
	Type    LogType   `json:"type"`
	Message string    `json:"message"`
	At      time.Time `json:"at"`
}

func (l *Logger) Add(logType LogType, message string) {
	mu.Lock()
	defer mu.Unlock()
	log := Log{
		Type:    logType,
		Message: message,
		At:      time.Now(),
	}
	l.logs = append(l.logs, log)
}

func (l *Logger) GetLogs() []Log {
	return l.logs
}

func (l *Logger) Save() error {
	data, err := json.Marshal(l.logs)
	if err != nil {
		return err
	}
	return l.storage.Save(data)
}

func (l *Logger) Load() error {
	data, err := l.storage.Load()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &l.logs)
}
