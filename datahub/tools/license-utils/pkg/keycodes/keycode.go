package keycodes

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"prophetstor.com/alameda/datahub/tools/license-utils/pkg/utils"
	"prophetstor.com/api/datahub/keycodes"
)

var (
	datahubAddress *string
)

func KeycodeInit(address *string) {
	datahubAddress = address
}

func Executor() (string, error) {
	prompt := promptui.Select{
		Label: "Select Option",
		Items: []string{"Add", "Read", "Delete", "Activate", "Generate Registration Data", "Back"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Invalid input value %v\n", err)
		return "", err
	}

	switch result {
	case "Add":
		keycode, _ := utils.InputText("Keycode")
		return "Add", AddKeycode(keycode)
	case "Read":
		keycode, _ := utils.InputText("Keycode")
		return "Read", ListKeycodes(keycode)
	case "Delete":
		keycode, _ := utils.InputText("Keycode")
		return "Delete", DeleteKeycode(keycode)
	case "Activate":
		filePath, _ := utils.InputText("Registration File Path")
		return "Activate", Activate(filePath)
	case "Generate Registration Data":
		return "Generate Registration Data", GenerateRegistrationData()
	default:
		return "Back", nil
	}
}

func PrintKeycode(keycode *keycodes.Keycode) {
	fmt.Println(fmt.Sprintf("Keycode: %s", keycode.Keycode))
	fmt.Println(fmt.Sprintf("KeycodeType: %s", keycode.KeycodeType))
	fmt.Println(fmt.Sprintf("KeycodeVersion: %d", keycode.KeycodeVersion))
	fmt.Println(fmt.Sprintf("ApplyTime: %s", keycode.ApplyTime))
	fmt.Println(fmt.Sprintf("ExpireTime: %s", keycode.ExpireTime))
	fmt.Println(fmt.Sprintf("Registered: %t", keycode.Registered))
	fmt.Println(fmt.Sprintf("LicenseState: %s", keycode.LicenseState))
	fmt.Println(fmt.Sprintf("Max Users: %d", keycode.Capacity.Users))
	fmt.Println(fmt.Sprintf("Max Hosts: %d", keycode.Capacity.Hosts))
	fmt.Println(fmt.Sprintf("Max Disks: %d", keycode.Capacity.Disks))
	fmt.Println(fmt.Sprintf("Diskprophet enabled: %t", keycode.Functionality.DiskProphet))
	fmt.Println(fmt.Sprintf("Workload enabled: %t", keycode.Functionality.Workload))
}
