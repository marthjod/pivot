# pivot

Read in [pivio](https://github.com/pivio/) file and convert to OpenNebula-compatible YAML according to YAML-defined mapping.

## Install

```bash
for p in model mapping convert; do
	go get github.com/marthjod/pivot/$p
done
```

## PoC

```bash
go run main.go -mapping mapping.yaml -pivio pivio.yaml

Runtime:
{CPU:S RAM:CUSTOM Disk:L HostType:VM NetworkZone:DMZ Environments:[foo bar]}
Mapping:
&{RAM:{Alias:memory Values:map[s:1024 m:2048 l:4096 xxl:16384 custom:65536]} CPU:{Alias:vcpu Values:map[l:4 s:1 m:2] RatioFactor:0.5} Disk:{Alias:image Values:map[s:20 m:20 l:40]}}
Converted:
memory: 65536
vcpu: 0.5
image: 40

Writing converted YAML to converted.yaml
```

