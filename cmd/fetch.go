package cmd

import (
	paramfetch "github.com/filecoin-project/go-paramfetch"
	cmds "github.com/ipfs/go-ipfs-cmds"
	"github.com/pkg/errors"

	"github.com/filecoin-project/venus/asset"
)

var fetchCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "fetch paramsters",
	},
	Options: []cmds.Option{
		cmds.Uint64Option(Size, "size to fetch"),
	},
	Run: func(req *cmds.Request, re cmds.ResponseEmitter, env cmds.Environment) error {
		// highest precedence is cmd line flag.
		if size, ok := req.Options[Size].(uint64); ok {
			ps, err := asset.Asset("_assets/proof-params/parameters.json")
			if err != nil {
				return err
			}
			if err := paramfetch.GetParams(req.Context, ps, size); err != nil {
				return errors.Wrapf(err, "fetching proof parameters: %v", err)
			}
			return nil
		}
		return errors.New("uncorrect parameters")
	},
}
