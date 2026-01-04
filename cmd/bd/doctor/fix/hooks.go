package fix

import (
	"fmt"
)

// GitHooks fixes missing or broken git hooks by calling bd hooks install
// DISABLED: Git hooks installation has been disabled in this build.
func GitHooks(path string) error {
	// DISABLED: Git hooks installation - too invasive for minimal setup
	_ = path // silence unused variable warning
	fmt.Println("Git hooks installation is disabled in this build.")
	return nil
}
