package base

import (
	"github.com/ofree8/service-base/config"
)

func init() {
	if err := config.Init("service-base"); err != nil {

	}
}
