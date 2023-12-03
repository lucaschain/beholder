![main branch status checks](https://github.com/lucaschain/beholder/actions/workflows/build_and_test.yml/badge.svg)
![coverage](https://raw.githubusercontent.com/lucaschain/beholder/badges/.badges/main/coverage.svg)

# Beholder

_Under development_


A simple CLI tool that watches files and runs commands based on it.

## Installation

Check the [releases page](https://github.com/lucaschain/beholder/releases) to get the correct URL for your platform install with:

```bash
export BEHOLDER_VERSION="0.0.8"
export BEHOLDER_PLATFORM="linux_amd64"
curl -LO https://github.com/lucaschain/beholder/releases/download/${BEHOLDER_VERSION}/beholder_${BEHOLDER_VERSION}_${BEHOLDER_PLATFORM}.tar.gz

tar -xvzf beholder_${BEHOLDER_VERSION}_${BEHOLDER_PLATFORM}.tar.gz

sudo mv beholder /usr/local/bin

```

## Usage

```bash
# this will start watching the files
beholder /tmp -- echo "I see that {file} was changed"

# Whenever a file in /tmp is changed, with a command like:
echo "hello" > /tmp/beholder_test.txt

# this will be printed:
# I see that /tmp/beholder_test.txt was changed
```
