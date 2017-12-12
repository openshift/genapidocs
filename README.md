# genapidocs
Location of the OpenShift tool used to generate API docs.

## Build

```
$ make build 
```

Binary will be placed under `_output/local/bin/genapidocs`

## Use

```
$ _output/local/bin/genapidocs --root=api-docs --base=<path/to/openshift/origin>
```

Generates api docs under `./api-docs`.

`--base` is the relative path to the directory containing `api/swagger-spec/openshift-openapi-spec.json`

## Update

```
make update-deps
```
