# Docker Self-Hosted Infra

This project is the successor of [this project](https://github.com/sirber/infra).
It is written in GoLang instead of Makefile. Ths biggest advantage is that in order 
to use the infra, you won't have to merge two git repositories, one for the infra, 
the other for `services/`.

# Installation

A binary is automatically compiled by GitHub.

- Download [latest release](https://github.com/sirber/infra2/releases)
- Move the downloaded binary `infra` to a folder in path
- Make it executable: `chmod a+x infra`

# Usage

In a peculiar folder, like `~/infra`