package infrastructure

import (
	"encoding/json"
	"path/filepath"

	bosherr "bosh/errors"
	boshdpresolv "bosh/infrastructure/devicepathresolver"
	boshplatform "bosh/platform"
	boshsettings "bosh/settings"
	boshdir "bosh/settings/directories"
	boshsys "bosh/system"
)

type wardenInfrastructure struct {
	fs                 boshsys.FileSystem
	dirProvider        boshdir.DirectoriesProvider
	platform           boshplatform.Platform
	devicePathResolver boshdpresolv.DevicePathResolver
}

func NewWardenInfrastructure(
	fs boshsys.FileSystem,
	dirProvider boshdir.DirectoriesProvider,
	platform boshplatform.Platform,
	devicePathResolver boshdpresolv.DevicePathResolver,
) (inf wardenInfrastructure) {
	inf.fs = fs
	inf.dirProvider = dirProvider
	inf.platform = platform
	inf.devicePathResolver = devicePathResolver
	return
}

func (inf wardenInfrastructure) GetDevicePathResolver() boshdpresolv.DevicePathResolver {
	return inf.devicePathResolver
}

func (inf wardenInfrastructure) SetupSsh(username string) error {
	return nil
}

func (inf wardenInfrastructure) GetSettings() (boshsettings.Settings, error) {
	var settings boshsettings.Settings

	// warden-cpi-agent-env.json is written out by warden CPI.
	settingsPath := filepath.Join(inf.dirProvider.BoshDir(), "warden-cpi-agent-env.json")
	contents, err := inf.fs.ReadFile(settingsPath)
	if err != nil {
		return settings, bosherr.WrapError(err, "Read settings file")
	}

	err = json.Unmarshal([]byte(contents), &settings)
	if err != nil {
		return settings, bosherr.WrapError(err, "Unmarshal json settings")
	}

	return settings, nil
}

func (inf wardenInfrastructure) SetupNetworking(networks boshsettings.Networks) error {
	return nil
}

func (inf wardenInfrastructure) GetEphemeralDiskPath(devicePath string) (string, bool) {
	return inf.platform.NormalizeDiskPath(devicePath)
}

func (inf wardenInfrastructure) MountPersistentDisk(volumeID string, mountPoint string) error {
	return nil
}