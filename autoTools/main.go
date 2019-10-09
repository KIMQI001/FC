package main

import "github.com/FC/FC/autoTools/tools"

func main() {
	// gen LocalDeploy bin file
	tools.LocalDeploy()

	// gen Remote bin file
	tools.RemoteDeploy()
}