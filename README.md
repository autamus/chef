# Chef

Chef is an autamus container recipe writer. Using package from [autamus.io](autamus.io),
we can generate a Dockerfile with any set of packages!

![img/chef.png](img/chef.png)

## Usage

### Build

To run chef, you can either run the go file directly:

```bash
$ go run chef.go
USAGE: chef CMD [OPTIONS]

DESCRIPTION: Generate scientific software container recipes

COMMANDS:

    NAME      ALIAS  DESCRIPTION
    generate  g      Generate a Dockerfile from a chef config file.
    help      ?      Get help with a specific subcommand
    version          Print the version of this command
```

You can also build the binary first, and then run it.

```bash
$ go build
$ ./chef
```

### Generate

Currently, the main command `generate` requires a chef config file that lists
packages that you want to install from [autamus.io](autamus.io). For example:

```yaml
# chef.yaml
packages:
 - clingo:latest
 - python:latest
```

We would then generate the Dockerfile as follows:

```bash
$ chef generate
```
or
```bash
$ go run chef.go generate
FROM ghcr.io/autamus/clingo:latest as clingo
FROM ghcr.io/autamus/python:latest as python
FROM spack/ubuntu-bionic
COPY --from=clingo /opt/software /opt/spack/opt/spack
COPY --from=python /opt/software /opt/spack/opt/spack
ENV PATH=/opt/spack/bin:$PATH
WORKDIR /opt/spack
RUN rm -rf opt/spack/.spack-db/
ENTRYPOINT ["/bin/bash"]
```

You could easily pipe this into a Dockerfile:

```bash
$ go run chef.go generate chef.yaml > Dockerfile
```

To skip validation that the images/tags exist (you'd find out when you build the Dockerfile) do:

```bash
$ go run chef.go generate --skip-validation
```

If you don't provide a chef.yaml file, it defaults to chef.yaml. However you
can provide a custom filename:

```bash
$ go run chef.go generate custom.yaml
```
