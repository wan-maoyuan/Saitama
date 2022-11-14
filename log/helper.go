package log

import (
	"fmt"
	"os"
)

type Helper struct {
	logger Logger
}

func NewHelper(logger Logger) *Helper {
	h := &Helper{
		logger: logger,
	}

	return h
}

func (h *Helper) Debug(msg ...any) {
	_ = h.logger.Log(LevelDebug, fmt.Sprint(msg...))
}

func (h *Helper) Debugf(format string, msg ...any) {
	_ = h.logger.Log(LevelDebug, fmt.Sprintf(format, msg...))
}

func (h *Helper) Info(msg ...any) {
	_ = h.logger.Log(LevelInfo, fmt.Sprint(msg...))
}

func (h *Helper) Infof(format string, msg ...any) {
	_ = h.logger.Log(LevelInfo, fmt.Sprintf(format, msg...))
}

func (h *Helper) Warn(msg ...any) {
	_ = h.logger.Log(LevelWarn, fmt.Sprint(msg...))
}

func (h *Helper) Warnf(format string, msg ...any) {
	_ = h.logger.Log(LevelWarn, fmt.Sprintf(format, msg...))
}

func (h *Helper) Error(msg ...any) {
	_ = h.logger.Log(LevelError, fmt.Sprint(msg...))
}

func (h *Helper) Errorf(format string, msg ...any) {
	_ = h.logger.Log(LevelError, fmt.Sprintf(format, msg...))
}

func (h *Helper) Fatal(msg ...any) {
	_ = h.logger.Log(LevelFatal, fmt.Sprint(msg...))
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, msg ...any) {
	_ = h.logger.Log(LevelFatal, fmt.Sprintf(format, msg...))
	os.Exit(1)
}
