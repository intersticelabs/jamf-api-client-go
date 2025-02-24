// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package policies

// Account represents an account set up in Jamf
type Account struct {
	Size    int            `json:"size"`
	Details AccountDetails `json:"account"`
}

// AccountDetails holds the specific account details
type AccountDetails struct {
	Action             string `json:"action"`
	Username           string `json:"username"`
	Realname           string `json:"realname"`
	Password           string `json:"password"`
	ArchiveHomDir      bool   `json:"archive_home_directory"`
	ArchiveHomeDirPath string `json:"archive_home_directory_to"`
	Home               string `json:"home"`
	Picture            string `json:"picture"`
	Admin              bool   `json:"admin"`
	FileVaultEnabled   bool   `json:"filevault_enabled"`
}

// ManagementAccount represents a management account type
type ManagementAccount struct {
	Action                string `json:"action"`
	ManagedPassword       string `json:"managed_password"`
	ManagedPasswordLength string `json:"managed_password_length"`
}
