package covid

/*



 */

import (
	"github.com/junkd0g/applogger"
)

var (
	logger applogger.AppLogger
)

func init() {
	logger.Initialise()
}

func Logger() applogger.AppLogger {
	return logger
}
