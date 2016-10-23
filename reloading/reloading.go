package reloading

import "github.com/theplant/containers"

func ReloadingScript() containers.Container {
	return containers.Script("reload.js")
}
