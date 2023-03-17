<div align="center">

<h1>ansdoc</h1>
<img src="https://github.com/FalcoSuessgott/ansdoc/actions/workflows/test.yml/badge.svg" alt="drawing"/>
<img src="https://github.com/FalcoSuessgott/ansdoc/actions/workflows/lint.yml/badge.svg" alt="drawing"/>
<img src="https://codecov.io/gh/FalcoSuessgott/ansdoc/branch/master/graph/badge.svg" alt="drawing"/>
<img src="https://img.shields.io/github/downloads/FalcoSuessgott/ansdoc/total.svg" alt="drawing"/>
<img src="https://img.shields.io/github/v/release/FalcoSuessgott/ansdoc" alt="drawing"/>
<img src="https://img.shields.io/docker/pulls/falcosuessgott/ansdoc" alt="drawing"/>
</div>

# Description
`ansdoc` is a dead-simple CLI tool written in Go, that scans your Ansible Role Vars file (usually `defaults/main.yml`, change by using `--file | -f` flag) and generates a Markdown Table accordingly. 

The variable description is taken be the leading headcomment of each variable (see [Example](https://github.com/FalcoSuessgott/ansdoc#example)).

`ansdoc` allows you to insert the generated Markdown Table in a specified output-file (`--output-file | -o`) between two `<!--ansdoc -->` separators (see [README.md](https://github.com/FalcoSuessgott/ansdoc/blob/b1808887c5cce60f45e80f36848547e08137840b/README.md?plain=1#L67)).

# Features
* support multiline comments
* configurable via env vars (`ANSDOC_FILE`, `ANSDOC_OUTPUT_FILE`, `ANSDOC_INSERT`, `ANSDOC_BACKUP`)
* insert table in a specified output file
* keep backup of output-file when inserting content

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

#we also support jinja & you dont even need a leading space :)
jinja: |
  [
  {% for server in groups[vault_raft_group_name] %}
    {
      "peer": "{{ server }}",
      "api_addr": "{{ hostvars[server]['vault_api_addr'] |
      default(vault_protocol + '://' + hostvars[server]['ansible_' + hostvars[server]['ansible_default_ipv4']['interface']]['ipv4']['address'] + ':' + (vault_port|string)) }}"
    },
  {% endfor %}
  ]
```

by running:

```bash
> ansdoc
```

`ansdoc` creates:

<!--ansdoc -->
<table>
<tr>
<th>Name</th>
<th>Description</th>
<th>Default Value</th>
</tr>

<tr>
<td>

```
enable
```

</td><td>enable the deployment</td><td>

```
true
```

</td></tr>

<tr>
<td>

```
domain
```

</td><td>specify your toplevel domain</td><td>

```
ansible.com
```

</td></tr>

<tr>
<td>

```
number_of_clients
```

</td><td>number of clients</td><td>

```
42
```

</td></tr>

<tr>
<td>

```
user_ids
```

</td><td>valid user IDs</td><td>

```
[1 2 3]
```

</td></tr>

<tr>
<td>

```
os
```

</td><td>supported OS</td><td>

```
map[linux:true mac:false windows:false]
```

</td></tr>

<tr>
<td>

```
jinja
```

</td><td>we also support jinja & you dont even need a leading space :)</td><td>

```
[
{% for server in groups[vault_raft_group_name] %}
  {
    "peer": "{{ server }}",
    "api_addr": "{{ hostvars[server]['vault_api_addr'] |
    default(vault_protocol + '://' + hostvars[server]['ansible_' + hostvars[server]['ansible_default_ipv4']['interface']]['ipv4']['address'] + ':' + (vault_port|string)) }}"
  },
{% endfor %}
]
```

</td></tr>

</table>
<!--ansdoc -->

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
