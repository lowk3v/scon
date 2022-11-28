# Smart Contract Discovery which detect chain of addresses and dump source theme

[![asciicast](https://asciinema.org/a/etH8bvQVq1hGyL83eRiunu9HH.svg)](https://asciinema.org/a/etH8bvQVq1hGyL83eRiunu9HH)

```
==================================================
 
        ███████╗ ██████╗ ██████╗ ███╗   ██╗
        ██╔════╝██╔════╝██╔═══██╗████╗  ██║
        ███████╗██║     ██║   ██║██╔██╗ ██║
        ╚════██║██║     ██║   ██║██║╚██╗██║
        ███████║╚██████╗╚██████╔╝██║ ╚████║
        ╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝  By @LowK3v
 Discovery contract of address provided on stdin
 Credits: https://github.com/lowk3v/scon
 Twitter: https://twitter.com/Low_K_
 ==================================================

Usage of: SCON_ENV=mainnet && scon <options> <mode> [address:-]
Options:
  -c, --chain <chain-name>      Specific chain name
  -d, --delay <delay>           DelayOpt between issuing requests (ms)
  -o, --output <dir>            Directory to save responses in (will be created)
  -x, --proxy <proxyURL>        Use the provided HTTP proxy

Mode:
  -dc, --detect-chain           Detect blockchain of address
  -l,  --supported-chains       List supported chain
  -ds  --dump-source            Dump contract source code. Output to folder specified by -o option. Default is ./contracts

```


```bash
go install github.com/lowk3v/scon@latest
```