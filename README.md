## Prerequisites
- make ([Installation guide for Ubuntu](https://zoomadmin.com/HowToInstall/UbuntuPackage/make))
- Docker ([Installation guide](https://docs.docker.com/engine/install/ubuntu))

## Installation
### Using Docker
#### Preparation
```sh
make network pull
```
#### Run
```sh
make run
```
#### Clear
```sh
make clear
```

## Endpoints
### Health check
```sh
make check
```

### Count
```sh
make count
```