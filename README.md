# randomfs

Generates x files with random data.

Options:
- specify exact content in files
- repeat exact content at certain frequencies (eg every 10th file created)
- specify max length of data in files

-----------------------------------------------------------
## Install

### Binary
[Download](../../releases) the latest stable release.

### Source
```
go get -u github.com/techjacker/randomfs
```


-----------------------------------------------------------
## Usage
```Shell
Usage: randomfs [options] <no_of_files_to_create>

Eg: randomfs 100

  -freq int
        frequency of the repeat text, eg every 5 files (default 5)
  -location string
        the location to put the files (default "/Users/andy/go/src/github.com/techjacker/randomfs")
  -max int
        the max number of characters per file (default 512)
  -prev int
        how often to repeat the content of all the previous files (default 11)
  -repeat string
        the repeat text (default "repeat_text")
```

-----------------------------------------------------------
## Tests
```Shell
go test
```

-----------------------------------------------------------
### Docker Example

Generates 100 files with random data.
1. Every 5th file contains "Nothing to see here"
2. Every 11th file contains the contents of all previous files
3. Random data files not exceed 512 characters in text

#### Build
```Shell
docker build -t randomfs/golang .
```

#### Run
```Shell
docker run -it --name randomfs_golang randomfs/golang /bin/bash

app -location /home 100

cd /home
ls -l /home

# random data
cat -v 000

# Every 5th file should contain "repeat_text"
cat -v 004

# Every 11th file should contain the contents of all previous files.
cat -v 010

# Every 11th file should contain the contents of all previous files.
# But capped at 512 characters
cat -v 098
```


