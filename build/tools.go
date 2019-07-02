// +build ignore

package tools

// This allows us to import the code-generator module as a dependency
// so that we can use it as part of the build process

import (
	_ "k8s.io/code-generator"
)
