# get

A *very* simple file/web viewer and fetcher.

`get [-f] [-o [filename]] <URL or path>`


## Examples:

| Command | Result |
|---------|--------|
| `get /etc/passwd` | Syntax-highlight `/etc/passwd` |
| `get example.com` | Syntax-hlighting `https://example.com` |
| `get -o example.com` | Save to `get.output.html` |
| `get -o example.com/foo.gif` | Save to `./foo.gif` |
| `get -o outp.gif example.com/foo.gif` | Save to `./outp.gif` |


## Features:

- Works on local paths and remote URLs (HTTP/HTTPS only)
  - If it's an existing file, it'll display that
  - URLs can be specified with or without a protocol
    (ex. `example.com` is the same as `https://example.com`)
  - Unless `http://` is specified, will always try HTTPS first
    and fall back to HTTP
- Automatic syntax highlighting if it's outputting to a terminal
- `-o` without a `filename` will guess an appropriate name
  - Extension is determined based on `Content-Type`
- Refuses to overwrite files, unless `-f` is passed
