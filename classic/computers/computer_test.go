// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.
package computers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	jamf "github.com/trustero/jamf-api-client-go/classic/computers"
	"github.com/stretchr/testify/assert"
)

var COMPUTER_API_BASE_ENDPOINT = "/JSSResource/computers"

func computerResponseMocks(t *testing.T) *httptest.Server {
	var resp string
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case COMPUTER_API_BASE_ENDPOINT:
			fmt.Fprintf(w, `{
				"computers": [
					{
							"id": 3,
							"name": "Test MacBook #3"
					},
					{
							"id": 28,
							"name": "Test MacBook #28"
					},
					{
							"id": 30,
							"name": "Test MacBook #30"
					},
					{
							"id": 31,
							"name": "Test MacBook #31"
					},
					{
							"id": 32,
							"name": "Test MacBook #32"
					},
					
					{
							"id": 91,
							"name": "Test MacBook #91"
					}]
			}`)
		case fmt.Sprintf("%s/id/82", COMPUTER_API_BASE_ENDPOINT):
			fmt.Fprintf(w, `{
				"computer": {
					"general": {
						"id": 82,
						"name": "Go Service Test Machine",
						"mac_address": "00:00:00:A0:FE:00",
						"serial_number": "VM0L+J/0cr+l",
						"udid": "000DF0BF-00FF-D00B-FA00-000F0DA0FE00",
						"jamf_version": "20.18.0-t0000000000",
						"platform": "Mac",
						"mdm_capable": false,
						"report_date": "2020-09-11 23:06:00"
					},
					"location": {
						"username": "test.user",
						"realname": "Test User",
						"real_name": "Test User",
						"email_address": "test.user@email.com",
						"position": "Software Engineer",
						"department": "Engineering",
						"building": "Boston"
					},
					"hardware": {
						"make": "Apple",
						"os_name": "Mac OS X",
						"os_version": "10.14.4",
						"os_build": "18E2034",
						"sip_status": "Enabled",
						"gatekeeper_status": "App Store and identified developers",
						"xprotect_version": "2114",
						"filevault2_users": [
							"test.user"
						]
					},
					"certificates": [{
						"common_name": "JSS Built-in Certificate Authority",
						"identity": false,
						"expires_utc": "9027-11-12T20:07:28.000+0000",
						"expires_epoch": 1826050048000,
						"name": ""
					}],
					"software": {
						"unix_executables": [],
						"licensed_software": [],
						"installed_by_casper": [
							"filevault_profile_signed.pkg",
							"Zoom-Latest.pkg"
						],
							"installed_by_installer_swu": [
								"com.datadoghq.pkg.Zoom-Latest",
								"com.github.makeprofilepkg.config_us",
								"com.googlecode.munki.admin",
								"com.googlecode.munki.app",
								"com.googlecode.munki.app_usage",
								"com.googlecode.munki.core",
								"com.googlecode.munki.launchd",
								"com.googlecode.munki.python",
								"com.vmware.tools.macos.pkg.files"
							],
							"available_software_updates": [],
							"available_updates": {},
							"running_services": [
								"com.apple.accessoryd",
								"com.apple.backupd",
								"com.apple.diagnosticd",
								"com.apple.mdmclient.daemon",
								"com.apple.mdmclient.daemon.runatboot",
								"com.googlecode.munki.appusaged",
								"com.jamf.management.daemon"
							],
							"applications": [{
								"name": "Datadog Agent.app",
								"path": "/Applications/Datadog Agent.app",
								"version": "7.16.1"
							}]
						},
						"extension_attributes": [
							{
								"id": 6,
								"name": "osquery Status",
								"type": "String",
								"multi_value": false,
								"value": "OSquery NOT Running"
							}
						],
						"groups_accounts": {
							"computer_group_memberships": [
								"Test Group for API Service"
							],
							"local_accounts": [{
								"name": "test.user",
								"realname": "Test User",
								"uid": "501",
								"administrator": true,
								"filevault_enabled": true
							}]
						},
						"configuration_profiles": [{
							"id": 2,
							"name": "Test Config Profile",
							"uuid": "abcdefghijklmnop123",
							"is_removable": false
						}]
				}
			}`)
		default:
			http.Error(w, fmt.Sprintf("bad Jamf computer API call to %s", r.URL), http.StatusInternalServerError)
			return
		}
		_, err := w.Write([]byte(resp))
		assert.Nil(t, err)
	}))
}
func TestQueryComputer(t *testing.T) {
	testServer := computerResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewService(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	computers, _, err := j.List()
	assert.Nil(t, err)
	assert.NotNil(t, computers)
	assert.Equal(t, 6, len(computers))
	assert.Equal(t, 3, computers[0].Id)
	assert.Equal(t, "Test MacBook #3", computers[0].Name)
}

func TestQuerySpecificComputer(t *testing.T) {
	testServer := computerResponseMocks(t)
	defer testServer.Close()
	j, err := jamf.NewService(testServer.URL, "fake-username", "mock-password-cool", nil)
	assert.Nil(t, err)
	computer, _, err := j.GetById(82)
	assert.Nil(t, err)
	// General Info
	assert.Equal(t, 82, computer.General.Id)
	assert.Equal(t, "Go Service Test Machine", computer.General.Name)
	assert.Equal(t, false, computer.General.MDMCapable)

	// User & Location Info
	assert.Equal(t, "Test User", computer.UserLocation.RealName)
	assert.Equal(t, "test.user@email.com", computer.UserLocation.EmailAddress)
	assert.Equal(t, "Software Engineer", computer.UserLocation.Position)
	assert.Equal(t, "Engineering", computer.UserLocation.Department)

	// Hardware info
	assert.Equal(t, "Apple", computer.Hardware.Make)
	assert.Equal(t, "App Store and identified developers", computer.Hardware.GatekeeperStatus)
	assert.Equal(t, "Enabled", computer.Hardware.SIPStatus)
	assert.Equal(t, []string{"test.user"}, computer.Hardware.FilevaultUsers)

	// Certificate Information
	assert.Equal(t, "JSS Built-in Certificate Authority", computer.Certificates[0].CommonName)

	// Software Information
	assert.Equal(t, []string{"filevault_profile_signed.pkg", "Zoom-Latest.pkg"}, computer.Software.InstalledByCasper)
	assert.Equal(t, "com.apple.accessoryd", computer.Software.RunningServices[0])
	assert.Equal(t, "Datadog Agent.app", computer.Software.Applications[0].Name)

	// Extension Attributes
	assert.Equal(t, 6, computer.ExtensionAttributes[0].ID)
	assert.Equal(t, "osquery Status", computer.ExtensionAttributes[0].Name)
	assert.Equal(t, "OSquery NOT Running", computer.ExtensionAttributes[0].Value)

	// Group Memberships & Local Accounts
	assert.Equal(t, []string{"Test Group for API Service"}, computer.Groups.Memberships)
	assert.Equal(t, "test.user", computer.Groups.LocalAccounts[0].Name)
	assert.Equal(t, true, computer.Groups.LocalAccounts[0].Administrator)

	// Config Profiles
	assert.Equal(t, 2, computer.ConfigProfiles[0].ID)
	assert.Equal(t, "Test Config Profile", computer.ConfigProfiles[0].Name)
	assert.Equal(t, false, computer.ConfigProfiles[0].Removable)
}
