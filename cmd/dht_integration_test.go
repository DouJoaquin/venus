package cmd_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/filecoin-project/venus/app/node"
	"github.com/filecoin-project/venus/app/node/test"
	tf "github.com/filecoin-project/venus/pkg/testhelpers/testflags"
)

func TestDhtFindPeer(t *testing.T) {
	tf.IntegrationTest(t)
	ctx := context.Background()

	builder1 := test.NewNodeBuilder(t)
	n1 := builder1.BuildAndStart(ctx)
	defer n1.Stop(ctx)
	cmdClient, done := test.RunNodeAPI(ctx, n1, t)
	defer done()

	builder2 := test.NewNodeBuilder(t)
	n2 := builder2.BuildAndStart(ctx)
	defer n2.Stop(ctx)

	node.ConnectNodes(t, n1, n2)

	n2Id := n2.Network().API().NetworkGetPeerID()
	findpeerOutput := cmdClient.RunSuccess(ctx, "dht", "findpeer", n2Id.String()).ReadStdoutTrimNewlines()
	n2Addr := n2.Network().API().NetworkGetPeerAddresses()[0]

	assert.Contains(t, findpeerOutput, n2Addr.String())
}

// TODO: findprovs will have to be untested until
//  https://github.com/filecoin-project/venus/issues/2357
//  original tests were flaky; testing may need to be omitted entirely
//  unless it can consistently pass.
