package fast

import (
	"context"
	"encoding/json"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/venus/pkg/chain"
)

// ChainHead runs the chain head command against the filecoin process.
func (f *Filecoin) ChainHead(ctx context.Context) ([]cid.Cid, error) {
	var out []cid.Cid
	if err := f.RunCmdJSONWithStdin(ctx, nil, &out, "venus", "chain", "head"); err != nil {
		return nil, err
	}
	return out, nil

}

// ChainLs runs the chain ls command against the filecoin process.
func (f *Filecoin) ChainLs(ctx context.Context) (*json.Decoder, error) {
	return f.RunCmdLDJSONWithStdin(ctx, nil, "venus", "chain", "ls")
}

// ChainStatus runs the chain status command against the filecoin process.
func (f *Filecoin) ChainStatus(ctx context.Context) (*chain.Status, error) {
	var out *chain.Status
	if err := f.RunCmdJSONWithStdin(ctx, nil, &out, "venus", "chain", "status"); err != nil {
		return nil, err
	}
	return out, nil
}
