# GIDEON

![alt text](https://themanabase.com/wp-content/uploads/2017/10/Gideon-Amonkhet-2-e1506932824569.jpg)

## Api to manage mtg decks.

Gideon believed that those with power should protect the weak, and took this belief personally, using his invulnerability to protect any in need. He was one to never lose faith in those he called friends, even when all others had, a belief he took to his death.
This is a backend api to save and manage mtg decks, focused in [Commander format](https://magic.wizards.com/pt-br/content/commander-format).


## Requirements

To Run this project you need [Docker](https://www.docker.com/), [docker-compose](https://docs.docker.com/compose/) and [Golang](https://golang.org/doc/install)

## Run
```
make run
```

## Testing

### Unit tests

Run:

```
make check
```

### Integration tests

Run:

```
make check-integration
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
