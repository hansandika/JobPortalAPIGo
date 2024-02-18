package infra

import (
	"fmt"

	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/constant"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var logger *logrus.Logger

func NewLogger(conf *general.SectionService) *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(&easy.Formatter{
			TimestampFormat: constant.FullTimeFormat,
			LogFormat:       fmt.Sprintf("%s\n", `[%lvl%]: "%time%" | "%msg%"`),
		})
		return logger
	}
	return logger
}
