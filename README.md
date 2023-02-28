# go-notion

![](https://img.shields.io/github/go-mod/go-version/kissy24/go-notion)
![](https://img.shields.io/github/license/kissy24/go-notion)

go-notion is a library for sending requests to Notion in the go language.

## Usage

1. Prepare a config.toml like the following.

    ```toml
    [Secret]
    key = "secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    ```

2. toml file to . /config.
3. Write the process you want to send to Notion under ./cmd/.

## Requirements

- Go
- github.com/BurntSushi/toml

## License

MIT

## Author

kissy24(yuhei-kishida)
