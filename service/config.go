package service

const (
	CrabadaGame = "0x82a85407bd612f52577909f4a58bfc6873f14da8" // Crabada Contract
	AvaxHttpUrl = "https://api.avax.network/ext/bc/C/rpc"
	ChainId     = 43114
	GasLimit    = 500000
	GasPrice    = 50
)

var (
	EndGameHex   = "0x2d6ef31000000000000000000000000000000000000000000000000000000000000" + Int642HexString(TeamID)
	StartGameHex = "0xe5ed1d590000000000000000000000000000000000000000000000000000000000000" + Int642HexString(TeamID)
)

// Change this value
const (
	TeamID    = 4000                                                               // TODO: replace with your own team id
	SecretKey = "0000000000000000000000000000000000000000000000000000000000000000" // TODO: replace with your own secret key

)
