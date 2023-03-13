# ansdoc
Out of the box documentation for any Ansible Role by using simple Comments

<div align="center">

<img src="https://github.com/FalcoSuessgott/ansdoc/actions/workflows/test.yml/badge.svg" alt="drawing"/>
<img src="https://github.com/FalcoSuessgott/ansdoc/actions/workflows/lint.yml/badge.svg" alt="drawing"/>
<img src="https://codecov.io/gh/FalcoSuessgott/ansdoc/branch/main/graph/badge.svg" alt="drawing"/>
<img src="https://img.shields.io/github/downloads/FalcoSuessgott/ansdoc/total.svg" alt="drawing"/>
<img src="https://img.shields.io/github/v/release/FalcoSuessgott/ansdoc" alt="drawing"/>
<img src="https://img.shields.io/docker/pulls/falcosuessgott/ansdoc" alt="drawing"/>
</div>

# Description
`ansdoc` scans your `defaults/main.yml` (change by using `--file | -f` flag) file of the current ansible role and generates a simple markdown table containing your vars including their description

# Example
```yaml
# bool
var1: true

# string
var2: string

# int
var3: 42 

# list
var4: [1, 2, 3]

# map
var5:
  linux: true
  mac: false
  windows: false

# multiline 
# comment
var6: "value with a multiline comment"
```

creates:

| VARIABLE | DESCRIPTION |                DEFAULT VALUE                |
|----------|-------------|---------------------------------------------|
| `var1`   | bool        | `"true"`                                    |
| `var2`   | string      | `"string"`                                  |
| `var3`   | int         | `"42"`                                      |
| `var4`   | list        | `"[1 2 3]"`                                 |
| `var5`   | map         | `"map[linux:true mac:false windows:false]"` |


# Installation
```bash
# curl
version=$(curl -S "https://api.github.com/repos/FalcoSuessgott/ansdoc/releases/latest" | jq -r '.tag_name[1:]')
curl -OL "https://github.com/FalcoSuessgott/ansdoc/releases/latest/download/ansdoc_${version}_$(uname)_$(uname -m).tar.gz"
tar xzf "vops_${version}_$(uname)_$(uname -m).tar.gz"
./ansdoc version

# Go 
go install github.com/FalcoSuessgott/ansdoc@latest
ansdoc version

# Docker/Podman
docker run ghcr.io/falcosuessgott/ansdoc version

# Ubuntu/Debian
version=$(curl -S "https://api.github.com/repos/FalcoSuessgott/ansdoc/releases/latest" | jq -r '.tag_name[1:]')
curl -OL "https://github.com/FalcoSuessgott/ansdoc/releases/latest/download/ansdoc_${version}_linux_amd64.deb"
sudo dpkg -i "./ansdoc_${version}_linux_amd64.deb"
vops version

# Fedora/CentOS/RHEL
version=$(curl -S "https://api.github.com/repos/FalcoSuessgott/ansdoc/releases/latest" | jq -r '.tag_name[1:]')
curl -OL "https://github.com/FalcoSuessgott/ansdoc/releases/latest/download/ansdoc_${version}_linux_amd64.rpm"
sudo dnf localinstall "./ansdoc_${version}_linux_amd64.rpm"
vops version

# Sources
git clone https://github.com/FalcoSuessgott/vops && cd ansdoc
go build 
```