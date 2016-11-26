# pivot

Read in [pivio](https://github.com/pivio/) file and convert to other format(s) according to YAML-defined mapping.

## Install

```
for p in model mapping convert; do
	go get github.com/marthjod/pivot/$p
done
```

## PoC

```
go run main.go -mapping mapping.yaml -pivio pivio.yaml
Runtime:
{CPU:S RAM:M Disk:L HostType:VM NetworkZone:DMZ Environments:[foo bar]}
Mapping:
&{RAM:{Alias:memory Values:{S:1024 M:2048 L:4096}} CPU:{Alias:vcpu Values:{S:1 M:2 L:4}} Disk:{Alias:image Values:{S:20 M:20 L:40}}}
Converted:
vcpu: 1, memory: 2048, image: 40
```