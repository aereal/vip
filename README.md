# vip [![wercker status](https://app.wercker.com/status/e8df97161d75205b16e7e0656f3b1022/s "wercker status")](https://app.wercker.com/project/bykey/e8df97161d75205b16e7e0656f3b1022) [![Coverage Status](https://coveralls.io/repos/aereal/vip/badge.svg?branch=HEAD&service=github)](https://coveralls.io/github/aereal/vip?branch=HEAD)

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
* Support lazy loading
* Take care of dependencies

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
