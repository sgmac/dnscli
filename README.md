# dnscli

dnscli is an app that interacts with DNSimple API from the CLI.

### Introduction

Basic operations such as delete, create or update your DNS entries should be quick. If you  
deploy a new server, it should be pretty fast to have DNS resolution,
this is why I came up with ```dnscli```.

### Installation

To install simply run:
```
$ go get -u github.com/sgmac/dnscli
```
The only dependency so far is [github.com/codegangsta/cli](https://github.com/codegangsta/cli)

Make sure you have properly set your GO environment before using the command. [See the install instructions](http://golang.org/doc/install.html).

### Configuration

The default path for the configuration is located in ```$HOME/.dnscli/config.json```. The first time you run ```dnscli``` it creates an empty configuration, 
you need to update the information with your credentials.

```
{
    "ApiURL": "https://api.dnsimple.com/v1/domains/",
    "Domain": "example.com",
    "Mail": "nobody@example.com",
    "Token": "your_token"
}
```

### Getting started

You can manage your records on a daily basis from the CLI, operations such as ```list, delete, create, update and get``` are fully supported. The first time you try to run it asks to configure your credentials. It will not work until you configure them.

#### Records
```
$ dnscli records list
Type         Name                     TTL       RecordID     Content
A            .example.com             3600      5163024      2.3.4.5
A            test.example.com         3600      4212532      7.7.7.7
A            nil.example.com          3600      6253203      1.2.3.6
A            lab.example.com          3600      4689691      27.3.14.37
A            monitoring.example.com   3600      3203149      134.213.137.66
A            dashb.example.com        3601      1133621      78.23.4.55
A            wiki.example.com         3600      5193141      62.24.1.32
CNAME        mail.example.com         3600      1307110      www.coolmail.com
CNAME        www.example.com          3600      2642364      example.com
MX           .example.com             3600      1189408      in1-smtp.com
MX           .example.com             3600      5389609      in2-smtp.com
SPF          .example.com             3600      4779156      v=spf1 include
TXT          .example.com             3600      8549313      v=spf1 include
```

You can provide a different domain from the CLI.

```$ dnscli -d example2.com records  list```

For operations such as ```delete, get or update``` you need to provide the **RecordID** as an option.

#### Auto renewal

Auto renewal can be enabled/disabled:
```
$ dnscli  autorenewal  -e
Domain                Lockable              AutoRenew
example.com             true                  true
```

#### Bash completion
It has built-in bash completion thanks to [codegangsta/cli](https://github.com/codegangsta/cli), however you will need to source it from
some place:

```
PROG=dnscli source $GOPATH/src/github.com/codegangsta/cli/autocomplete/bash_autocomplete
```

#### TODO
I just wanted something that works, some improvements :

- Move ```records``` and ```autorenewal``` to packages so if somebody wants to use it as a library may integrate it 
 on their apps.
- Better testing
