# Gokrazy StatusPage

Tidy looking statuspage for running in gokrazy

## How do I use this?

```
$ make
$ ./statuspage
```

Or, if you want to be able to change web assets live and see how it looks quickly:

```
$ make dev
$ ./statuspage-dev
```

You will need to configure the site to be useful. Here is a configuration I drew up earlier.

```
base-url: /look/ma/you/can/put/me/behind/a/reverse/proxy/
columns: 3
tests:
  - type: vaelatern-ping
    host: Server

  - type: vaelatern-port-open
    desc: Server Up
    host: server
    port: 443
    proto: tcp

  - type: vaelatern-http-200
    desc: Server Responding
    url: https://server/

  - type: vaelatern-ping
    host: Kiosk1

  - type: vaelatern-ping
    host: Kiosk2

  - type: vaelatern-ping
    host: Kiosk3
```

## Why are the types prefixed with your github handle?

I wanted to lower the cost to add new modules. Namespacing with github handles allows each contributor to write their own modules and I, as a maintainer, can just merge any modules that don't break existing behaviour and match the naming requirement.

Also lets someone else implement e.g. ping in a different way if need be.

## Docs plz

TODO

## This looks like [Monitoror](https://monitoror.com/)

Indeed. I took a screenshot of Monitoror, fed it into Vercel's V0.dev to [generate](https://v0.dev/t/PgjBM0G63Wj) the initial look of the site.

Monitoror was so close to what I wanted, but it required a nodejs build step. That meant I couldn't reasonably force the software to build on `go get`. I loved their look (and even had their icon in mind when I drew the favicon) but did not use their code at all, not even as a reference when doing my plugins.
