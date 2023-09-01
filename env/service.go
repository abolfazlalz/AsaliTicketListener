package env

import "github.com/caarlos0/env/v9"

type ServiceType struct {
	Interval string `env:"SERVICE_INTERVAL_TIME" envDefault:"10s"`
}

var Service *ServiceType

func loadService() error {
	Service = &ServiceType{}
	if err := env.Parse(Service); err != nil {
		return err
	}
	return nil
}
