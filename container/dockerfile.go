package container

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// packageExists determines if a package and tag exist under autamus
func packageExists(pkg string) bool {
	var cli = &http.Client{Timeout: 10 * time.Second}
	_, err := cli.Get("https://crane.ggcr.dev/manifest/ghcr.io/autamus/" + pkg)
	if err != nil {
		return false
	}
	return true
}

func Dockerfile(packages []string, validate bool) string {

	dockerfile := ""

	// Keep a list of package names to COPY from after
	var packageNames []string

	// Ensure that each package exists as we add
	for _, pkg := range packages {

		parts := strings.Split(pkg, ":")
		name := parts[0]
		version := strings.ReplaceAll(parts[1], ".", "-")
		name = name + "-" + version

		if validate {
			if packageExists(pkg) {
				dockerfile += ("FROM ghcr.io/autamus/" + pkg + " as " + name + "\n")
				packageNames = append(packageNames, name)
			} else {
				fmt.Println("Warning, %s does not exist.", pkg)
			}
		} else {
			dockerfile += ("FROM ghcr.io/autamus/" + pkg + " as " + name + "\n")
			packageNames = append(packageNames, name)
		}
	}

	// Add the spack base container
	dockerfile += "FROM spack/ubuntu-bionic\n"

	// Add the COPY commands
	for _, pkg := range packageNames {
		dockerfile += ("COPY --from=" + pkg + " /opt/software /opt/spack/opt/spack\n")
	}

	// Finish up!
	dockerfile += "ENV PATH=/opt/spack/bin:$PATH\n"
	dockerfile += "WORKDIR /opt/spack\n"

	// We need to remove the spack db so it is re-created
	dockerfile += "RUN rm -rf opt/spack/.spack-db/\n"
	dockerfile += "ENTRYPOINT [\"/bin/bash\"]"
	return dockerfile
}
