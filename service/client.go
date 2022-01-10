package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
)

type CrabClient struct {
	RestyClient *resty.Client
	Client      *ethclient.Client
}

func NewCrabClient() *CrabClient {
	client, err := ethclient.Dial(AvaxHttpUrl)
	if err != nil {
		log.Fatal("NewService dial", err)
	}

	return &CrabClient{
		RestyClient: resty.New(),
		Client:      client,
	}
}

func (c *CrabClient) StartGame(teamID int64) {
	startGameHex := "0xe5ed1d590000000000000000000000000000000000000000000000000000000000000" + Int642HexString(teamID)
	tx, err := c.SendTxWithHexData(SecretKey, CrabadaGame, startGameHex, GasPrice, GasLimit)
	if err != nil {
		log.Println("HandleBuyPotion SendTxWithHexData", err)
		return
	}
	log.Println("Play Game", tx)
}

func (c *CrabClient) EndGame(gameID int64) {
	endGameHex := "0x2d6ef31000000000000000000000000000000000000000000000000000000000000" + Int642HexString(gameID)
	tx, err := c.SendTxWithHexData(SecretKey, CrabadaGame, endGameHex, GasPrice, GasLimit)
	if err != nil {
		log.Println("HandleBuyPotion SendTxWithHexData", err)
		return
	}
	log.Println("Play Game", tx)
}

func (s *CrabClient) PrivateToPublicKey(secret string) string {
	privateKey, err := crypto.HexToECDSA(secret)
	if err != nil {
		log.Println("GetPublicKey error", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return ""
	}
	publicKeyHex := crypto.PubkeyToAddress(*publicKeyECDSA)
	return publicKeyHex.String()
}

func (s *CrabClient) SendTxWithHexData(secret, toAddress, txData string, gWei float64, gasLimit int64) (string, error) {
	privateKey, err := crypto.HexToECDSA(secret)
	if err != nil {
		log.Println("SendTxWithHexData ecdsa", err)
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("SendTxWithHexData Public", "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", err
	}
	publicKeyHex := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.Client.PendingNonceAt(context.Background(), publicKeyHex)
	fmt.Println(nonce)
	if err != nil {
		log.Println("SendTxWithHexData PendingNonceAt", err)
		return "", err
	}

	to := ethCommon.HexToAddress(toAddress)
	data, err := DecodeString(txData)
	if err != nil {
		log.Println("SendTxWithHexData DecodeString", err)
		return "", err
	}
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    big.NewInt(0),
		Gas:      uint64(gasLimit),
		GasPrice: big.NewInt(int64(gWei * 1e9)),
		Data:     data,
	})
	signer := types.LatestSignerForChainID(big.NewInt(ChainId))
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
	if err != nil {
		fmt.Println("SendTxWithHexData Sign", err)
		return "", err
	}
	signedTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		fmt.Println("SendTxWithHexData WithSignature", err)
		return "", err
	}
	return signedTx.Hash().String(), s.Client.SendTransaction(context.Background(), signedTx)
}
