# gthsManage

<!-- TOC -->

- [gthsManage](#gthsmanage)
- [What's it for?](#whats-it-for)
- [Usage](#usage)
- [Config file example](#config-file-example)

<!-- /TOC -->

# What's it for?
Gthsmanage is used to keep the noticeboard that I made up and running in an easily managable way when I am no longer able to manage the noticeboard. I also made it because I was a bit bored really. But hey its a useful thing to make.
Also its basically just a wrapper for some shell scripts.

# Usage
```
Stuff to keep the GTHS Noticeboard running.

Usage:
  gthsManage [command]

Available Commands:
  config                 Config checker / creator
  deploy                 Redeploy the noticeboard if Chrome is already open.
  everythingisbrokenhelp Contact me directly.
  help                   Help about any command
  is                     Checks the status of various things.
  reboot                 Reboot the GTHS Noticeboard.
  resetrepo              Reset the repo on the noticeboard.
  update                 Update packages.

Flags:
  -c, --config string   config file (default is $HOME/.gthsManage.yaml)
  -h, --help            help for gthsManage
  -f, --idfile string   Full path to your private ssh key (default "/Users/willb/.ssh/id_rsa")
  -i, --ip string       IP of noticeboard (default "10.178.x.x")
  -p, --port string     SSH port of noticeboard (default "22")

Use "gthsManage [command] --help" for more information about a command.
```

# Config file example

```yml
ip: 10.178.1.170
port: 22
idfile: /Users/willb/.ssh/id_rsa
```
