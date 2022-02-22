package main

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/tamalsaha/image-digester-demo/pkg/version"
)

func main() {
	image := "kubedb/operator:v0.25.0-dbg_linux_arm64"
	digest, err := crane.Digest(image,
		crane.WithUserAgent(fmt.Sprintf("%s/%s", "digester", version.Version)))
	if err != nil {
		panic(err)
	}
	fmt.Println(digest)
}
