package fix

import (
	"fmt"
)

// MergeDriver fixes the git merge driver configuration to use correct placeholders.
// DISABLED: Git merge driver installation has been disabled in this build.
func MergeDriver(path string) error {
	// DISABLED: Git merge driver installation - too invasive for minimal setup
	_ = path // silence unused variable warning
	fmt.Println("Git merge driver installation is disabled in this build.")
	return nil
}
