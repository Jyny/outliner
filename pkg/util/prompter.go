package util

import (
	"fmt"

	"github.com/Songmu/prompter"
)

func ContinueInteractive() bool {
	fmt.Println("The next step is to deploy.")
	fmt.Println("Deploy outliner VPN to server just create.")
	fmt.Println("You can continue to auto deploy or Do it your self.")
	fmt.Println("This is same as `./outliner deploy -i {SERVER_IP}`.")
	return prompter.YN("Do you want to continue to Auto Deploying", true)
}
