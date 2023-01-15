<div align="center">

# **torrent-client**


</div>

# BitTorrent client written in Go

## Install

```sh
go get github.com/pchchv/torrent-client
```
## Usage
### Try downloading [Debian](https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/#indexlist)!

```sh
torrent-client debian-10.2.0-amd64-netinst.iso.torrent debian.iso
```

### Limitations
* #### Only supports `.torrent` files (no magnet links)
* #### Only supports HTTP trackers
* #### Does not support multi-file torrents
* #### Strictly leeches (does not support uploading pieces)

#

<div align="right">

###### *Inspired by [Jesse Li](https://blog.jse.li/posts/torrent/)*

</div>
