# Beholder

_Under development_

A simple CLI tool that watches files and runs commands based on it.


## Usage

```bash
# this will start watching the files
beholder /tmp -- echo "I see that {file} was changed"

# Whenever a file in /tmp is changed, with a command like:
echo "hello" > /tmp/beholder_test.txt

# this will be printed:
# I see that /tmp/beholder_test.txt was changed
```
