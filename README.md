# xiny

A simple command line tool for converting between various units of measurement

<p align="center"><img src="https://xiny.sh/img/screencap.gif" alt="xiny"/></p>

## Install

Fetch the [latest release](https://github.com/bcicen/xiny/releases) for your platform:

#### Linux

```bash
sudo wget https://github.com/bcicen/xiny/releases/download/v0.3.3/xiny-0.3.3-linux-amd64 -O /usr/local/bin/xiny
sudo chmod +x /usr/local/bin/xiny
```

#### OS X

```bash
sudo curl -Lo /usr/local/bin/xiny https://github.com/bcicen/xiny/releases/download/v0.3.3/xiny-0.3.3-darwin-amd64
sudo chmod +x /usr/local/bin/xiny
```

#### Docker

```bash
docker run --rm -ti \
           --name=xiny \
           quay.io/vektorlab/xiny:latest
```

## Usage

Conversions may be passed in long form:
```
$ xiny 20 kilograms in pounds
44.092452 pounds
```

or shortened form with symbols:
```
$ xiny 20kg in lb
44.092452 pounds
```

Use the verbose flag(`-v`) to print the formula used for the conversion:
```
$ xiny -v 32C in F
celsius -> farenheit: (x * 1.8 + 32)
89.6 farenheit
```

### Interactive mode

If no positional arguments are provided, `xiny` will be started in interactive mode, providing a prompt for conversions with autocomplete and other useful features

### Options
Option | Description
--- | ---
-n | display only numeric output (exclude units)
-v | enable more verbose output (twice for debug output)
