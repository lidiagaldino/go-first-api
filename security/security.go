package security

import "github.com/lidiagaldino/go-first-api/config"

var (
	logger *config.Logger
)

func Init() {
	logger = config.GetLogger("security")
}
