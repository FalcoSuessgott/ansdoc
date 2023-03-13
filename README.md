# ansdoc
Out of the box documentation for any Ansible Role simply by using headcomments in your var file

<div align="center">

<img src="https://github.com/FalcoSuessgott/ansdoc/actions/workflows/test.yml/badge.svg" alt="drawing"/>
<img src="https://github.com/FalcoSuessgott/ansdoc/actions/workflows/lint.yml/badge.svg" alt="drawing"/>
<img src="https://codecov.io/gh/FalcoSuessgott/ansdoc/branch/master/graph/badge.svg" alt="drawing"/>
<img src="https://img.shields.io/github/downloads/FalcoSuessgott/ansdoc/total.svg" alt="drawing"/>
<img src="https://img.shields.io/github/v/release/FalcoSuessgott/ansdoc" alt="drawing"/>
<img src="https://img.shields.io/docker/pulls/falcosuessgott/ansdoc" alt="drawing"/>
</div>

# Description
`ansdoc` scans your `defaults/main.yml` (change by using `--file | -f` flag) file of the current ansible role and generates a simple markdown table containing your vars including their description

# Example
`defaults/main.yml`
```yaml
# this enables the deployment
enable: true

# specify your toplevel domain
domain: ansible.com

# number of clients
number_of_clients: 42 

# valid user IDs
user_ids: [1, 2, 3]

# supported OS
os:
  linux: true
  mac: false
  windows: false
```

by running:

```bash
> ansdoc
```

`ansdoc` creates:

|      VARIABLE       |         DESCRIPTION          |                DEFAULT VALUE                |
|---------------------|------------------------------|---------------------------------------------|
| `enable`            | this enables the deployment  | `"true"`                                    |
| `domain`            | specify your toplevel domain | `"ansible.com"`                             |
| `number_of_clients` | number of clients            | `"42"`                                      |
| `user_ids`          | valid user IDs               | `"[1 2 3]"`                                 |
| `os`                | supported OS                 | `"map[linux:true mac:false windows:false]"` |

# Installation
```bash
# curl
version=$(curl -S "https://api.github.com/repos/FalcoSuessgott/ansdoc/releases/latest" | jq -r '.tag_name[1:]')
curl -OL "https://github.com/FalcoSuessgott/ansdoc/releases/latest/download/ansdoc_${version}_$(uname)_$(uname -m).tar.gz"
tar xzf "ansdoc_${version}_$(uname)_$(uname -m).tar.gz"
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
ansdoc

# Fedora/CentOS/RHEL
version=$(curl -S "https://api.github.com/repos/FalcoSuessgott/ansdoc/releases/latest" | jq -r '.tag_name[1:]')
curl -OL "https://github.com/FalcoSuessgott/ansdoc/releases/latest/download/ansdoc_${version}_linux_amd64.rpm"
sudo dnf localinstall "./ansdoc_${version}_linux_amd64.rpm"
ansdoc

# Sources
git clone https://github.com/FalcoSuessgott/ansdoc && cd ansdoc
go build 
```