package app

import (
	"encoding/json"

	boshinf "github.com/cloudfoundry/bosh-agent/infrastructure"
	bosherr "github.com/cloudfoundry/bosh-agent/internal/github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-agent/internal/github.com/cloudfoundry/bosh-utils/system"
	boshplatform "github.com/cloudfoundry/bosh-agent/platform"
)

type Config struct {
	Platform       boshplatform.Options
	Infrastructure boshinf.Options
}

func LoadConfigFromPath(fs boshsys.FileSystem, path string) (Config, error) {
	var config Config

	if path == "" {
		return config, nil
	}

	bytes, err := fs.ReadFile(path)
	if err != nil {
		return config, bosherr.WrapError(err, "Reading file")
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, bosherr.WrapError(err, "Loading file")
	}

	return config, nil
}
