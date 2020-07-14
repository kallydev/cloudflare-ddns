# CloudFlare-DDNS

![GitHub Last Commit](https://img.shields.io/github/last-commit/kallydev/forward?style=flat-square)
![GitHub License](https://img.shields.io/github/license/kallydev/forward?style=flat-square)

A CloudFlare DDNS service.

## Instructions

### 1. Install

- Build from source

You need to install `upx-ucl` and `p7zip-full`, the compiled files are located in the `release` folder.

```bash
git clone https://github.com/kallydev/cloudflare-ddns
cd cloudflare-ddns
bash bash script/build.sh --release
```

### 2. Config

If you need to enable `TLS`, please see [config.json](config.json).

- Client example

```json
{
  "cloud_flare": {
    "key": "b57a34a824e81f05c6fee5be3cef10f015665"
  },
  "client": {
    "server": "example.com:50051",
    "domain": "www.example.com"
  }
}
````

- Server example

```json
{
  "cloud_flare": {
    "email": "example@email.com",
    "key": "b57a34a824e81f05c6fee5be3cef10f015665",
    "zone_id": "c0a6482d6cb17960b5bfdfa3d256f5f1",
    "account_id": "95cbb560688820be3be3e08571da12b4"
  },
  "server": {
    "host": "0.0.0.0",
    "port": 50051
  }
}
```

### 3. Run

- Server

```bash
./cloudflare-ddns_linux_amd64 server -c config.json
```

- Client

```bash
./cloudflare-ddns_linux_amd64 client -c config.json
```

## License

Copyright (c) KallyDev. All rights reserved.

Licensed under the [MIT](LICENSE) license.
