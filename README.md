# go-notion

![](https://img.shields.io/github/go-mod/go-version/kissy24/go-notion)
![](https://img.shields.io/github/license/kissy24/go-notion)

go-notion is a library for sending requests to Notion in the go language.  
The implementation follows the [Notion Developer reference](https://developers.notion.com/reference/intro).

I also found the following article helpful.  
https://zenn.dev/shinshin/articles/1f30809917f048c7d86a

## Usage

### Preparation

1. Prepare a config.toml like the following.

    ```toml
    [Secret]
    key = "secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    ```

2. toml file to `./config`.
3. Write the process you want to send to Notion under `./cmd/`.

### Sample

`./cmd/main.go` to check some operations.  
Due to relative paths `./cmd/` due to relative paths. (I will fix it eventually.)

```sh
cd cmd
go run main.go
```

## Requirements

- Go
- github.com/BurntSushi/toml

## License

MIT

## Author

kissy24(yuhei-kishida)
