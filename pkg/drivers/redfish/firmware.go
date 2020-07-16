package redfish

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/bm-metamorph/MetaMorph/pkg/db/models/node"
	"github.com/manojkva/go-redfish-plugin/pkg/isogen"
	version "github.com/hashicorp/go-version"
)

func (bmhnode *BMHNode) UpgradeEachFirmware(filepath string) bool {
	redfishClient := getRedfishClient(bmhnode)
	return redfishClient.UpgradeFirmware(filepath)
}

func IsVersionHigher(providedVersion string, versionfromNode string) bool {
	var err error
	vprovided, err := version.NewVersion(providedVersion)
	if err != nil {
		fmt.Printf("Provided version could not be parsed")
		return false
	}
	vfromNode, err := version.NewVersion(versionfromNode)
	if err != nil {
		fmt.Printf("Firmware version from node could not be parsed")
		return false
	}
	if vfromNode.Equal(vprovided) {
		fmt.Printf("Version provided is equal to one in the node. Proceeding with installation..")
		return  true
	}
	if vfromNode.LessThan(vprovided) {
		fmt.Printf("Version provided is lower than one in the node")
		return false
	}
	return true
}

func (bmhnode *BMHNode) CheckUpgradeAllowed(providedName string, providedVersion string) bool {
	redfishClient := getRedfishClient(bmhnode)
	name, version, updateavailable := redfishClient.GetFirmwareDetails(providedName)
	if (name == "") || (version == "") {
		fmt.Printf("Failed to retrieve firmware details for %v\n", providedName)
		return false
	}
	if updateavailable == false {
		fmt.Printf("Firmware not updateable")
		return false
	}
	if IsVersionHigher(providedVersion, version) {
		return true
	}
	return true
}

func (bmhnode *BMHNode) UpdateFirmware() error {
	//iterate through the list of firmwares
	firwares, err := node.GetFirmwares(bmhnode.NodeUUID.String())
	if err != nil {
		return fmt.Errorf("Failed to retreive firmwareURL. %v", err)
	}
	for _, firmware := range firwares {

		if bmhnode.CheckUpgradeAllowed(firmware.Name, firmware.Version) {

			filename := path.Base(firmware.URL)

			tempdir, err := ioutil.TempDir("/tmp", "firmware")
			if err != nil {
				return fmt.Errorf("Failed to create temporary directory. %v", err)
			}
			defer os.RemoveAll(tempdir)
			firmwarefilepath := path.Join(tempdir, filename)
			err = isogen.DownloadUrl(firmwarefilepath, firmware.URL)

			if err != nil {
				return fmt.Errorf("Failed to Download URL, %v",err)
			}
			res := bmhnode.UpgradeEachFirmware(firmwarefilepath)
			if res == false {
				return  fmt.Errorf("Failed to Upgrade Firmware")
			}
		} else{
			return fmt.Errorf("Check for Upgrade version info failed for %v and version %v", firmware.Name, firmware.Version)
		}

	}

	return nil

}
