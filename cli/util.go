package cli

import cli "github.com/urfave/cli/v2"

// ConcatSlices concats multiple slices of flags together
func ConcatSlices(flags ...[]cli.Flag) []cli.Flag {
	var tmp []cli.Flag
	for _, s := range flags {
		tmp = append(tmp, s...)
	}
	return tmp
}
