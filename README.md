# pivot

[![Travis CI Build Status](https://travis-ci.org/marthjod/pivot.svg?branch=master)](https://travis-ci.org/marthjod/pivot)

Convert [pivio](https://github.com/pivio/) files.

## Install

```bash
go get github.com/marthjod/pivot
```

## Run

```
  -aliases string
    	Path to alias mapping (input) (default "aliases.yaml")
  -pivio string
    	Path to pivio.yaml (input) (default "pivio.yaml")
```

```
Read from pivio.yaml:
&{Id:CFPAnnouncement ShortName:CFPA Runtime:{CPU:XL RAM:S Disk:S HostType:Docker NetworkZone:DMZ Environments:[]} Services:{Provides:[{Description:Web Display of the Announcement ServiceName:web-announcement-service Protocol:https Port:80 TransportProtocol:tcp} {Description:REST API for updating CfP data ServiceName:rest-announcement-service Protocol:https Port:9449 TransportProtocol:tcp}] DependsOn:{Internal:[{Name:user-service} {Name:email-announcement-service}] External:[]}}}
Read from customformat/aliases.yaml:
&map[vcpu:map[XL:8 M:2] memory:map[S:2048 M:4096] image:map[S:ubuntu-latest-minimal M:ubuntu-latest] cpu:map[XL:2 M:0.5]]
2017/07/12 18:43:03 Output:
hg-CFPA:
  cpu: 2
  image: ubuntu-latest-minimal
  memory: 2048
  networks:
  - zone-ACCESS
  - zone-DMZ-CFPA
  - zone-LOGGING
  vcpu: 8
  ```
