package httpcfg

import (
	"github.com/ledgerwatch/erigon-lib/kv/kvcache"
	"github.com/ledgerwatch/erigon/eth/ethconfig"
	"github.com/ledgerwatch/erigon/node/nodecfg/datadir"
	"github.com/ledgerwatch/erigon/rpc/rpccfg"
)

type HttpCfg struct {
	Enabled        bool
	PrivateApiAddr string
	WithDatadir    bool // Erigon's database can be read by separated processes on same machine - in read-only mode - with full support of transactions. It will share same "OS PageCache" with Erigon process.
	WithSnapdir    bool // This is to cover when snapshots directory is not located in the same directory
	DataDir        string
	SnapDir        string
	Dirs           datadir.Dirs

	HttpListenAddress        string
	AuthRpcHTTPListenAddress string
	TLSCertfile              string
	TLSCACert                string
	TLSKeyFile               string

	HttpPort           int
	AuthRpcPort        int
	HttpCORSDomain     []string
	HttpVirtualHost    []string
	AuthRpcVirtualHost []string
	HttpCompression    bool
	API                []string
	Gascap             uint64
	MaxTraces          uint64

	WebsocketEnabled     bool
	WebsocketCompression bool
	RpcAllowListFilePath string
	RpcBatchConcurrency  uint
	RpcStreamingDisable  bool
	DBReadConcurrency    int
	TraceCompatibility   bool // Bug for bug compatibility for trace_ routines with OpenEthereum
	TxPoolApiAddr        string
	TevmEnabled          bool
	StateCache           kvcache.CoherentConfig
	Snap                 ethconfig.Snapshot
	Sync                 ethconfig.Sync

	GRPCServerEnabled      bool
	GRPCListenAddress      string
	GRPCPort               int
	GRPCHealthCheckEnabled bool
	StarknetGRPCAddress    string
	JWTSecretPath          string // Engine API Authentication
	TraceRequests          bool   // Always trace requests in INFO level
	HTTPTimeouts           rpccfg.HTTPTimeouts
	AuthRpcTimeouts        rpccfg.HTTPTimeouts
}
