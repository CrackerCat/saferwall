// Copyright 2021 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package avast

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/saferwall/saferwall/pkg/utils"
)

const (
	cmd         = "scan"
	avastDaemon = "/usr/bin/avast"
	licenseFile = "/etc/avast/license.avastlic"
	vpsUpdate   = "/var/lib/avast/Setup/avast.vpsupdate"
)

// Result represents detection results.
type Result struct {
	Infected bool   `json:"infected"`
	Output   string `json:"output"`
}

// GetVPSVersion returns VPS version.
func GetVPSVersion() (string, error) {

	// Run the scanner to grab the version
	out, err := utils.ExecCommand(cmd, "-V")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

// GetProgramVersion returns program version.
func GetProgramVersion() (string, error) {

	// Run the scanner to grab the version
	out, err := utils.ExecCommand(cmd, "-v")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

// ScanFilePath scans from a filepath.
func ScanFilePath(filepath string) (Result, error) {

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return Result{}, err
	}

	// Execute the scanner with the given file path
	//  -a         Print all scanned files/URLs, not only infected.
	//  -b         Report decompression bombs as infections.
	//  -f         Scan full files.
	//  -u         Report potentionally unwanted programs (PUP).
	out, err := utils.ExecCommand(cmd, "-abfu", filepath)

	// Exit status:
	//  0 - no infections were found
	//  1 - some infected file was found
	//  2 - an error occurred
	if err != nil && err.Error() != "exit status 1" {
		return Result{}, err
	}

	// Check if the file is clean.
	if strings.Contains(out, "[OK]") {
		return Result{}, nil
	}

	// Sanitize the detection output
	det := strings.Split(out, "\t")
	if len(det) < 1 {
		errUnexpectedOutput := fmt.Errorf("Unexpected output: %s", out)
		return Result{}, errUnexpectedOutput
	}

	res := Result{}
	res.Output = strings.TrimSpace(det[1])
	res.Infected = true
	return res, nil
}

// ScanFileBinary scan from a given buffer of bytes.
func ScanFileBinary(r io.Reader) (Result, error) {
	_, err := utils.WriteBytesFile("sample", r)
	if err != nil {
		return Result{Output: ""}, err
	}

	return ScanFilePath("sample")
}

// ScanURL scans a given URL.
func ScanURL(url string) (string, error) {

	// Execute the scanner with the given URL
	out, err := utils.ExecCommand(cmd, "-U", url)

	// Exit status:
	//  0 - no infections were found
	//  1 - some infected file was found
	//  2 - an error occurred
	if err != nil && err.Error() != "exit status 1" {
		return "", err
	}

	// Check if we got a clean URL
	if out == "" {
		return "[OK]", nil
	}

	// Sanitize the output
	str := strings.Split(out, "\t")
	result := strings.TrimSpace(str[1])
	return result, nil
}

// UpdateVPS performs a VPS update.
func UpdateVPS() error {
	_, err := utils.ExecCommand(vpsUpdate)
	return err
}

// IsLicenseExpired returns true if license was expired.
func IsLicenseExpired() (bool, error) {

	if _, err := os.Stat(licenseFile); os.IsNotExist(err) {
		return true, errors.New("License not found")
	}

	out, err := utils.ExecCommand(avastDaemon, "status")
	if err != nil {
		return true, err
	}

	if strings.Contains(out, "License expired") {
		return true, nil
	}
	return false, nil
}

// ActivateLicense activate the license.
func ActivateLicense(r io.Reader) error {
	// Write the license file to disk
	_, err := utils.WriteBytesFile(licenseFile, r)
	if err != nil {
		return err
	}

	// Change the owner of the license file to `avast` user
	err = utils.ChownFileUsername(licenseFile, "avast")
	if err != nil {
		return err
	}

	// Restart the service to apply the license
	err = RestartService()
	if err != nil {
		return err
	}

	isExpired, err := IsLicenseExpired()
	if err != nil {
		return err
	}

	if isExpired {
		return errors.New("License was expired ")
	}

	return nil
}

// RestartService re-starts the Avast service.
func RestartService() error {
	// check if service is running
	_, err := utils.ExecCommand(avastDaemon, "status")
	if err != nil && err.Error() != "exit status 3" {
		return err
	}

	// exit code 3 means program is not running
	if err.Error() == "exit status 3" {
		_, err = utils.ExecCommand(avastDaemon, "start")
	} else {
		_, err = utils.ExecCommand(avastDaemon, "restart")
	}

	return err
}

// StartDaemon starts the Avast daemon.
func StartDaemon() error {
	err := utils.StartCommand("sudo", avastDaemon, "-n", "-D")
	time.Sleep(5 * time.Second)
	err = utils.StartCommand("sudo", avastDaemon, "-n", "-D")
	return err
}
