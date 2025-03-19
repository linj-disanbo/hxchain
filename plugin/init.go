package plugin

import (
	_ "github.com/linj-disanbo/hxchain/plugin/consensus/init" //register consensus init package
	_ "github.com/linj-disanbo/hxchain/plugin/crypto/init"
	_ "github.com/linj-disanbo/hxchain/plugin/dapp/init"
	_ "github.com/linj-disanbo/hxchain/plugin/mempool/init"
	_ "github.com/linj-disanbo/hxchain/plugin/p2p/init"
	_ "github.com/linj-disanbo/hxchain/plugin/store/init"
)
