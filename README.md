# TPic-Downloader

[![CircleCI](https://circleci.com/gh/yuki383/TPic-Downloader.svg?style=svg)](https://circleci.com/gh/yuki383/TPic-Downloader)

TPic-Downloaderは任意のtwimg URLから画像をダウンロードすることができるコマンドラインツールです。

## Build from source  
```bash
  $ go get github.com/yuki383/tpic-fetcher/cmd/tpicker
```

## Useage
```bash
 $ tpicker [name] [url] [...options]
 ```


| option | detail |
| --- | --- |
| p | ダウンロード先のパス |
| f | 画像の拡張子 (jpg or png) |
| s | 画像サイズ (small or midium or large or thumb) |

