# dnscli

dnscli is an app that interacts with DNSimple API from the CLI.

### Introduction

Basic operations such as delete, create or update your DNS entries should be quick. If you  
deploy a new server, it should be pretty fast to have your DNS resolution,
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


```
$ dnscli records list
Type         Name                     TTL       RecordID     Content
A            .example.com             3600      2193074      2.3.4.5
A            test.example.com         3600      4712596      7.7.7.7
A            nil.example.com          3600      2253342      1.2.3.6
A            lab.example.com          3600      2796691      27.3.14.37
A            monitoring.example.com   3600      2253149      134.213.137.66
A            dashb.example.com        3601      2193521      78.23.4.55
A            wiki.example.com         3600      2193141      62.24.1.32
CNAME        mail.example.com         3600      3009110      www.coolmail.com
CNAME        www.example.com          3600      3642864      example.com
MX           .example.com             3600      3009108      in1-smtp.com
MX           .example.com             3600      3009109      in2-smtp.com
SPF          .example.com             3600      3009111      v=spf1 include
TXT          .example.com             3600      3009112      v=spf1 include
```

You can provide a different domain from the CLI.

```$ dnscli -d example2.com records  list```

For operations such as ```delete, get or update``` you need to provide the **RecordID** as an option.
