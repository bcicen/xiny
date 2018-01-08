# xiny

CLI tool for unit conversions

## Install

Fetch the [latest release](https://github.com/bcicen/xiny/releases) for your platform:

#### Linux

```bash
sudo wget https://github.com/bcicen/xiny/releases/download/v0.1/xiny-0.1-linux-amd64 -O /usr/local/bin/xiny
sudo chmod +x /usr/local/bin/xiny
```

#### OS X

```bash
sudo curl -Lo /usr/local/bin/xiny https://github.com/bcicen/xiny/releases/download/v0.1/xiny-0.1-darwin-amd64
sudo chmod +x /usr/local/bin/xiny
```

## Usage

Conversions may be passed in long form:
```bash
xiny 20 kilograms in pounds
```

or shortened form with symbols:
```bash
xiny 20kg in lb
```

Use the verbose flag(`-v`) to print the formula used for the conversion:
```bash
# xiny -v 32C in F
celsius -> farenheit: (x * 1.8 + 32)
89.6 farenheit
```

### Options
Option | Description
--- | ---
-i | start xiny in interactive mode
-v | enable verbose output
-vv | enable debug output
-list | list all potential unit names and exit
