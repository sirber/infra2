# Docker Self-Hosted Infra

This project is the successor of [this project](https://github.com/sirber/infra).
It is written in GoLang instead of Makefile. Ths biggest advantage is that in order 
to use the infra, you won't have to merge two git repositories, one for the infra, 
the other for `services/`.

# Installation

A binary is automatically compiled by GitHub.

- Download [latest release](https://github.com/sirber/infra2/releases)
  - example: `wget -O infra https://github.com/sirber/infra2/releases/download/0.2.2/infra` 
- Move the downloaded binary `infra` to a folder in path
- Make it executable: `chmod a+x infra`

# Usage

In a peculiar folder, like `~/infra`, init your infra by running `infra init`. 
A `infra.jon` config file will be created. Then, you can create folders and sub-folders
for your services. 

At minimum, a service must have a `docker-compose.yml` and be enabled by creating
an empty file `enabled`. 

Then, you are ready to start the infra! With `infra up`.

## Best Practices

- Use [bind mount](https://docs.docker.com/engine/storage/bind-mounts/), to be able to backup your data
- Set [restart policy](https://www.baeldung.com/ops/docker-compose-restart-policies) to "always"
- Down your service before deleting it, or else it will run forever like a ghost!

## Backup

You can do a full cold backup by running `sudo infra backup`. The backup will keep file 
permission, for the bind mounts. The archive file name is in the `infra.json` config file.

To restore, just extract the archive somewhere else (as root) and up it!

# Development

TODO