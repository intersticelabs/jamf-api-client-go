// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

package accounts

// Account represents an account set up in Jamf
type JamfAccountsResp struct {
	AccountsId JamfAccountsId `json:"accounts,omitempty"`
}
type JamfAccountsId struct {
	UsersIds  []JamfUserId  `json:"users,omitempty"`
	GroupsIds []JamfGroupId `json:"groups,omitempty"`
}

type JamfUserId struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type JamfGroupId struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type JamfUserResp struct {
	User JamfUser `json:"account,omitempty"`
}

type JamfUser struct {
	Id                  int    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	DirectoryUser       bool   `json:"directory_user,omitempty"`
	FullName            string `json:"full_name,omitempty"`
	Email               string `json:"email,omitempty"`
	EmailAddress        string `json:"email_address,omitempty"`
	PasswordSha256      string `json:"password_sha256,omitempty"`
	Enabled             string `json:"enabled,omitempty"`
	ForcePasswordChange bool   `json:"force_password_change,omitempty"`
	AccessLevel         string `json:"access_level,omitempty"`
	PrivilegeSet        string `json:"privilege_set,omitempty"`
	Privileges          struct {
		JssObjects  []string `json:"jss_objects,omitempty"`
		JssSettings []string `json:"jss_settings,omitempty"`
		JssActions  []string `json:"jss_actions,omitempty"`
		Recon       []string `json:"recon,omitempty"`
		CasperAdmin []string `json:"casper_admin,omitempty"`
	} `json:"privileges,omitempty"`
}

type JamfGroupResp struct {
	Group JamfGroup `json:"account,omitempty"`
}

type JamfGroup struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	AccessLevel  string `json:"access_level,omitempty"`
	PrivilegeSet string `json:"privilege_set,omitempty"`
	Privileges   struct {
		JssObjects    []string `json:"jss_objects,omitempty"`
		JssSettings   []string `json:"jss_settings,omitempty"`
		JssActions    []string `json:"jss_actions,omitempty"`
		Recon         []string `json:"recon,omitempty"`
		CasperAdmin   []string `json:"casper_admin,omitempty"`
		CasperRemote  []string `json:"casper_remote,omitempty"`
		CasperImaging []string `json:"casper_imaging,omitempty"`
	} `json:"privileges,omitempty"`
	Members []JamfUserId `json:"members,omitempty"`
}
