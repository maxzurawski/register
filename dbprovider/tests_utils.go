package dbprovider

import (
	"os"

	"github.com/maxzurawski/utilities/stringutils"
)

func EnvironmentPreparations() {
	_ = os.Setenv("SERVICE_NAME", "register")
	_ = os.Setenv("HTTP_PORT", "8102")
	_ = os.Setenv("EUREKA_SERVICE", "http://xdevicesdev.home:8761")
	userHomeDir := stringutils.UserHomeDir()
	_ = os.Setenv("DB_PATH", userHomeDir+"/.databases/xdevices/test/sensorregister.db")
	_ = os.Setenv("CONNECT_TO_RABBIT", "FALSE")
	InitDbManager()
}
