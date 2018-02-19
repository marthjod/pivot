# pivot

[![Travis CI Build Status](https://travis-ci.org/marthjod/pivot.svg?branch=master)](https://travis-ci.org/marthjod/pivot)

Retrieve/convert [pivio](https://github.com/pivio/) files.

## Example (API backend)

```
./pivot -h
  -a string
    	Pivio query API endpoint (default "http://pivio.example.com:9123")
  -s string
    	Artifact shortname to query (default "AAA")
```

## Example (file backend)

```
./pivot -h
  -aliases string
    	Path to alias mapping (input) (default "aliases.yaml")
  -format string
    	Conversion output format (default "default")
  -pivio string
    	Path to pivio.yaml (input) (default "pivio.yaml")
```

### Output formats

Input: example [pivio.yaml](https://github.com/marthjod/pivot/blob/master/pivio.yaml).

```
./pivot -format simple
```
```json
{
  "CFPA": {
    "cpu": "XL",
    "disk": "S",
    "memory": "S",
    "zone": "DMZ"
  }
}
```

```
./pivot -format custom -aliases formats/custom/aliases.yaml
```
```yaml
service-CFPA:
  cpu: 2
  image: ubuntu-latest-minimal
  memory: 2048
  networks:
  - zone-ACCESS
  - zone-DMZ-CFPA
  - zone-LOGGING
  vcpu: 8
```
