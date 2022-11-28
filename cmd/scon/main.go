package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"scon/pkg/scon"
)

var args scon.Arguments

func banner() string {
	// http://patorjk.com/software/taag/#p=testall&f=Cards&t=SCON
	return fmt.Sprintln(
		color.YellowString("==================================================\n"),
		color.HiBlueString(`
	███████╗ ██████╗ ██████╗ ███╗   ██╗
	██╔════╝██╔════╝██╔═══██╗████╗  ██║
	███████╗██║     ██║   ██║██╔██╗ ██║
	╚════██║██║     ██║   ██║██║╚██╗██║
	███████║╚██████╗╚██████╔╝██║ ╚████║
	╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝ `+" By @LowK3v\n"),
		color.BlueString("Discovery contract of address provided on stdin\n"),
		"Credits: https://github.com/lowk3v/scon\n",
		"Twitter: https://twitter.com/Low_K_\n",
		color.YellowString("=================================================="),
	)
}

func init() {
	args = scon.Arguments{}

	// delay time between requests
	flag.IntVar(&args.DelayOpt, "delay", 200, "DelayOpt between requests (ms)")
	flag.IntVar(&args.DelayOpt, "d", 200, "DelayOpt between requests (ms)")

	// specified chain id
	flag.StringVar(&args.ChainNameOpt, "chain", "", "Specified chain name split on comma")
	flag.StringVar(&args.ChainNameOpt, "c", "", "Specified chain name split on comma")

	// output folder path
	flag.StringVar(&args.OutputOpt, "output", "contracts", "Specified output folder path")
	flag.StringVar(&args.OutputOpt, "o", "contracts", "Specified output folder path")

	// set proxy options
	flag.StringVar(&args.ProxyOpt, "proxy", "", "Specified proxy options")
	flag.StringVar(&args.ProxyOpt, "x", "", "Specified proxy options")

	// MODE

	// detect chain only
	flag.BoolVar(&args.DetectChainMode, "detect-chain", false, "Detect chain only")
	flag.BoolVar(&args.DetectChainMode, "dc", false, "Detect chain only")

	// list supported chain
	flag.BoolVar(&args.SupportedChainMode, "list-supported-chains", false, "List supported chains")
	flag.BoolVar(&args.SupportedChainMode, "l", false, "List supported chains")

	// dump source
	flag.BoolVar(&args.DumpSourceMode, "dump-source", false, "Dump source code")
	flag.BoolVar(&args.DumpSourceMode, "ds", false, "Dump source code")

	flag.Usage = func() {
		h := []string{
			banner(),
			"Usage of: SCON_ENV=mainnet && scon <options> <mode> [address:-]",
			"Options:",
			"  -c, --chain <chain-name>  	Specific chain name",
			"  -d, --delay <delay>       	DelayOpt between issuing requests (ms)",
			"  -o, --output <dir>        	Directory to save responses in (will be created)",
			"  -x, --proxy <proxyURL>    	Use the provided HTTP proxy",
			"",
			"Mode:",
			"  -dc, --detect-chain      	Detect blockchain of address",
			"  -l,  --supported-chains	List supported chain",
			"  -ds  --dump-source 		Dump contract source code. Output to folder specified by -o option. Default is ./contracts\n",
		}

		_, _ = fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
	flag.Parse()

}

func main() {
	fmt.Println(banner())
	scon.Run(&args)
}
