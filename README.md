## GIDEON


### Api to manage commander decks.

Gideon believed that those with power should protect the weak, and took this belief personally, using his invulnerability to protect any in need. He was one to never lose faith in those he called friends, even when all others had, a belief he took to his death


# project

TODO brief intro

## Why

TODO

## Testing

Run:

```
make test
```
Open generated coverage on a browser:

```
make coverage
```
To perform static analysis:

```
make lint
```

## Releasing

Run:

```
make release version=<version>
```

It will create a git tag with the provided **<version>**
and build and publish a docker image.

## Git Hooks

To install the project githooks run:

```
make githooks
```
