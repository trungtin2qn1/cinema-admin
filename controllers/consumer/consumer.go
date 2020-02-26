package consumer

import "github.com/qor/admin"

// SetupResource ...
func SetupResource(resource *admin.Resource) *admin.Resource {
	setupSearchHandler(resource)

	return resource
}
