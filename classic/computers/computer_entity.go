// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package computers

import "encoding/xml"

type ComputerNameId struct {
	Id   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Computers represents a list of computers enrolled in Jamf
type Computers struct {
	List []BasicComputerInfo `json:"computers"`
}

// ComputerGroup represents a group a device is a member of in Jamf
type ComputerGroup struct {
	ID      int    `json:"id,omitempty" xml:"id,omitempty"`
	Name    string `json:"name" xml:"name"`
	IsSmart bool   `json:"is_smart" xml:"is_smart,omitempty"`
}

// BasicComputerInfo represents the information returned in a list of all computers from Jamf
type BasicComputerInfo struct {
	GeneralInformation
}

// Computer represents an individual computer enrolled in Jamf with all its associated information

type Computer struct {
	General             GeneralInformation       `json:"general"`
	UserLocation        LocationInformation      `json:"location"`
	Hardware            HardwareInformation      `json:"hardware"`
	Certificates        []CertificateInformation `json:"certificates"`
	Software            SoftwareInformation      `json:"software"`
	ExtensionAttributes []ExtensionAttributes    `json:"extension_attributes"`
	Groups              GroupInformation         `json:"groups_accounts"`
	ConfigProfiles      []ConfigProfile          `json:"configuration_profiles"`
	Security            SecurityInformation      `json:"security,omitempty"`
}

// GeneralInformation holds basic information associated with Jamf device
type GeneralInformation struct {
	XMLName      xml.Name `json:"-" xml:"computer,omitempty"`
	Id           int      `json:"id,omitempty" xml:"id,omitempty"`
	Name         string   `json:"name" xml:"name,omitempty"`
	MACAddress   string   `json:"mac_address" xml:"mac_address,omitempty"`
	SerialNumber string   `json:"serial_number" xml:"serial_number,omitempty"`
	UDID         string   `json:"udid" xml:"udid,omitempty"`
	JamfVersion  string   `json:"jamf_version" xml:"jamf_version,omitempty"`
	Platform     string   `json:"platform" xml:"platform,omitempty"`
	MDMCapable   bool     `json:"mdm_capable" xml:"mdm_capable,omitempty"`
	ReportDate   string   `json:"report_date" xml:"report_date,omitempty"`
}

// LocationInformation holds the information in the User & Locations section
type LocationInformation struct {
	Username     string `json:"username"`
	RealName     string `json:"realname"`
	EmailAddress string `json:"email_address"`
	Position     string `json:"position"`
	Department   string `json:"department"`
	Building     string `json:"building"`
}

// HardwareInformation holds the hardware specific device information
type HardwareInformation struct {
	Make                        string    `json:"make"`
	Model                       string    `json:"model"`
	ModelIdentifier             string    `json:"model_identifier"`
	OSName                      string    `json:"os_name"`
	OSVersion                   string    `json:"os_version"`
	OSBuild                     string    `json:"os_build"`
	SoftwareUpdateDeviceID      string    `json:"software_update_device_id"`
	ActiveDirectoryStatus       string    `json:"active_directory_status"`
	ServicePack                 string    `json:"service_pack"`
	ProcessorType               string    `json:"processor_type"`
	IsAppleSilicon              bool      `json:"is_apple_silicon"`
	ProcessorArchitecture       string    `json:"processor_architecture"`
	ProcessorSpeed              int       `json:"processor_speed"`
	ProcessorSpeedMhz           int       `json:"processor_speed_mhz"`
	NumberProcessors            int       `json:"number_processors"`
	NumberCores                 int       `json:"number_cores"`
	TotalRAM                    int64     `json:"total_ram"`
	TotalRAMMb                  int64     `json:"total_ram_mb"`
	BootRom                     string    `json:"boot_rom"`
	BusSpeed                    int       `json:"bus_speed"`
	BusSpeedMhz                 int       `json:"bus_speed_mhz"`
	BatteryCapacity             int       `json:"battery_capacity"`
	CacheSize                   int       `json:"cache_size"`
	CacheSizeKb                 int       `json:"cache_size_kb"`
	AvailableRAMSlots           int       `json:"available_ram_slots"`
	OpticalDrive                string    `json:"optical_drive"`
	NicSpeed                    string    `json:"nic_speed"`
	SmcVersion                  string    `json:"smc_version"`
	BleCapable                  bool      `json:"ble_capable"`
	SupportsIosAppInstalls      bool      `json:"supports_ios_app_installs"`
	SipStatus                   string    `json:"sip_status"`
	GatekeeperStatus            string    `json:"gatekeeper_status"`
	XProtectVersion             string    `json:"xprotect_version"`
	InstitutionalRecoveryKey    string    `json:"institutional_recovery_key"`
	DiskEncryptionConfiguration string    `json:"disk_encryption_configuration"`
	FilevaultUsers              []string  `json:"filevault2_users"`
	Storage                     []Storage `json:"storage"`
}

//type Storage struct {
//	Device Device `json:"device"`
//}

type Storage struct {
	Disk            string      `json:"disk"`
	Model           string      `json:"model"`
	Revision        string      `json:"revision"`
	SerialNumber    string      `json:"serial_number"`
	Size            int64       `json:"size"`
	DriveCapacityMB int64       `json:"drive_capacity_mb"`
	ConnectionType  string      `json:"connection_type"`
	SmartStatus     string      `json:"smart_status"`
	Partition       []Partition `json:"partitions"`
}
type Partition struct {
	Name                 string `json:"name"`
	Size                 int64  `json:"size"`
	PartitionType        string `json:"type"`
	PartitionCapacityMB  int64  `json:"partition_capacity_mb"`
	PercentageFull       int    `json:"percentage_full"`
	FilevaultStatus      string `json:"filevault_status"`
	FilevaultPercent     int    `json:"filevault_percent"`
	Filevault2Status     string `json:"filevault2_status"`
	Filevaul2tPercent    int    `json:"filevault2_percent"`
	BootDriveAvailableMB int64  `json:"boot_drive_available_mb"`
	LvgUUID              string `json:"lvg_uuid"`
	LvUUID               string `json:"lv_uuid"`
	PvUUID               string `json:"pv_uuid"`
}

// CertificateInformation holds information about certs intalled on the device
type CertificateInformation struct {
	CommonName string `json:"common_name"`
	Identity   bool   `json:"identity"`
	ExpiresUTC string `json:"expires_utc"`
	Name       string `json:"name"`
}

// SoftwareInformation holds information about the software installed on a device
type SoftwareInformation struct {
	UnixExecutables          []string                 `json:"unix_executables"`
	InstalledByCasper        []string                 `json:"installed_by_casper"`
	InstalledByInstaller     []string                 `json:"installed_by_installer_swu"`
	AvailableSoftwareUpdates []string                 `json:"available_software_updates"`
	RunningServices          []string                 `json:"running_services"`
	Applications             []ApplicationInformation `json:"applications"`
}

// ApplicationInformation holds information about the applications on a device
type ApplicationInformation struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

// ExtensionAttributes holds extension attribute information for a device
type ExtensionAttributes struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GroupInformation holds the groups the device is a member of
type GroupInformation struct {
	Memberships   []string `json:"computer_group_memberships"`
	LocalAccounts []struct {
		Name             string `json:"name"`
		RealName         string `json:"realname"`
		UID              string `json:"uid"`
		Administrator    bool   `json:"administrator"`
		FilevalutEnabled bool   `json:"filevault_enabled"`
	} `json:"local_accounts"`
}

// ConfigProfile represents an active configuration profile in Jamf
type ConfigProfile struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	Removable bool   `json:"is_removable"`
}
type SecurityInformation struct {
	ActivationLock      bool   `json:"activation_lock,omitempty"`
	RecoveryLockEnabled bool   `json:"recovery_lock_enabled,omitempty"`
	SecureBootLevel     string `json:"secure_boot_level,omitempty"`
	ExternalBootLevel   string `json:"external_boot_level,omitempty"`
	FirewallEnabled     bool   `json:"firewall_enabled,omitempty"`
}
