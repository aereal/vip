# vip

Vim plugin manager

**Caution: This software is still like a sketch. API or any design may change without notice.**

## Usage

TBD

## Philosophy

`vip` **does**:

* Install a plugin from Git repository
* Pin the plugin version to specified version
* Generate installed plugin information with versions
* Update installed plugins
* Run hooks when a plugin updated
* In manner of popular tools such as bundler, carton, or npm

`vip` **does not**:

* Manage `rtp`
* Supports lazy loading
* Takes care of dependencies

## Requirement

* Go
* make

## Build

```sh
make # Install dependencies and build from source
```

## Test

To run tests:

```sh
make test
```

## Author

aereal
