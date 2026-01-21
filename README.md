# edit_project

Quick navigation for projects

## Installation:

```sh
# download/install with go
go install github.com/timhugh/edit_project/cmd/edit-cli

# activate in your shell (or add to your shell rc)
source <(edit-cli install zsh)
# replace zsh with bash if that's how you roll

# I also highly recommend some aliases
alias ep="edit_project"
alias op="open_project"
```

If you don't want to install with Go, you can also use a [pre-built binary](https://github.com/timhugh/edit_project/releases) and put it somewhere on your $PATH.

## Usage

### The short version

To navigate to a project directory:
```sh
# if you set up aliases above
op
# if you didn't
open_project
```

To open a project in your editor (see [configuration](#configuration)):
```sh
# with aliases
ep
# without
edit_project
```

### The long version

Let's say I want to navigate to my dotfiles repo (on my system that's `~/git/timhugh/dotfiles`):
```sh
op dotf
```

That will open fzf with a list of all the projects I have (based on my [configuration](#configuration)), and pre-filter the fuzzy finder with the query "dotf". When I select "timhugh/dotfiles", fzf will exit and my shell will cd to `~/git/timhugh/dotfiles`. Easy!

Alternatively, maybe I want to open _this_ project in my [configured](#configuration) editor. In that case:
```sh
ep edit
```
Again, fzf opens with a list of all of the projects I have pre-filtered with the query "edit". When I select "timhugh/edit_project", my shell will cd to `~/git/timhugh/edit_project` and neovim will open in that directory. Double easy!

## Configuration

`edit-cli` has some convenience commands for working with your config (which you can see by running `edit-cli config`):

- `edit-cli create` will write a config file (by default `~/.config/edit_project/config.json`)
- `edit-cli show` will print your current configuration
- `edit-cli edit` will open your config file in your configured editor (admittedly a bit chicken and egg if you haven't configured an editor yet)
[//]: # (TODO: when $EDITOR fallback is added, update that last bullet)

The most basic configuration (which is the default) anticipates you having a workspace folder (~/git by default) that contains projects/repos organized by user/repo, like:

- ~/git/timhugh/edit_project
- ~/git/junegunn/fzf

If you are me (which would be weird), then you would want your configuration to look something like this:
```json
{
  "workspaces": [
    { "path": "~/git", "user_prefixes": true }
  ],
  "git_users": [ "timhugh" ],
  "editor": "nvim"
}
```

This means that:
1. There is a `~/git` directory that contains repos _inside_ of username directories
2. A repo like `timhugh/edit_project` is owned by me (and can generally just be referred to as `edit_project`)
3. When I invoke `edit_project`, I want neovim to open

If you are also me, you might have some repositories for work that have to be in a specific folder because some paths are hard-coded (I don't want to talk about it), but you still want to be able to navigate quickly. In that case, your config might have this:
```json
{
  "workspaces": [
    { "path": "~/git", "user_prefixes": true },
    { "path": "~/workspace", "user_prefixes": false }
  ],
  // the rest
```

Now if we imagine my home directory looks like this:
```
~
  git
    timhugh
      dotfiles
      edit_project
  workspace
    server
    frontend
```

I can run `op` and fzf will list:
```
timhugh/dotfiles
timhugh/edit_project
workspace/server
workspace/frontend
```

Triple easy!

## Contributing

Issues and PRs are welcome!

This project doesn't use anything fancy, just standard Go stuff:
- `go test ./...` to run tests
- `go run ./cmd/edit-cli` to run the CLI app

For ease of testing, I highly recommend that after making a change to the cli, you can use `go install ./cmd/edit-cli` to build it and put it in your $GOPATH. Once you're done testing, you can use `go clean -i ./cmd/edit-cli` to uninstall! Handy.
