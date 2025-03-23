# pair-api: wtf i build this?

## Installation

```sh
# check if go is installed
go -v
```

```sh
# clone repo locally
git clone <url>
cd <dir>

# install project dependencies
go mod tidy

# compile and run go program
go run .
```

## Usage

```sh
# to get first 10 items in descending order [defalut]
http :3000/api/get

# to get top "n" number of responses
http :3000/api/get?limit={n}

# to search `aname` or `bname` in database
http :3000/api/get?query={query}

# to get item in ascending order
http :3000/api/get?order={asc}

# to get item in descending order
http :3000/api/get?order={desc}

# to get first 3 items in descending order
http :3000/api/get?limit=3&order=desc

# to add an item with values "aintnodev" and "stelafella"
http :3000/api/add aname="aintnodev" bname="stelafella"

# to delete a value with id "id"
http DELETE :3000/api/delete/{id}
```
