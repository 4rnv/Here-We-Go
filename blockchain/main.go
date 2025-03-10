package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    int64
	PreviousHash []byte
	Hash         []byte
	Data         []byte
	Nonce        int
}

type Blockchain struct {
	Blocks []*Block
}

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

const targetBits = 16 // lower value makes it easier to crack, higher value makes it more difficult
const maxNonce = math.MaxInt64

// func IntToHex(integer int64) []byte {
// 	hex_value := fmt.Sprintf("%x", integer)
// 	return []byte(hex_value)
// }

func IntToHex(integer int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, integer)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func printBlock(block *Block) {
	fmt.Printf("Block Index: %d\n", block.Index)
	fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
	fmt.Printf("Block Data: %s\n", block.Data)
	fmt.Printf("Block Hash: %x\n", block.Hash)
	fmt.Println("---------------------------------------------------------------------------------------")
}

func (block *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PreviousHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

func createBlock(data string, previousBlock *Block) *Block {
	block := &Block{previousBlock.Index + 1, time.Now().Unix(), previousBlock.Hash, []byte{}, []byte(data), 0}
	//block.setHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (bc *Blockchain) addBlock(data string) {
	prevBlocks := bc.Blocks[:len(bc.Blocks)-1] //This creates a sub-blockchain with all elements except the last one
	fmt.Println(prevBlocks)
	prevBlock := bc.Blocks[len(bc.Blocks)-1] //This gets the last block in the chain. Noootice the different syntax in slice addressing
	newBlock := createBlock(data, prevBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func neonGenesisBlock() *Block {
	return createBlock("Init Block", &Block{Index: 0, PreviousHash: []byte{}})
}

func newBlockchain() *Blockchain {
	return &Blockchain{[]*Block{neonGenesisBlock()}}
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits)) // Will set target as 2^256-targetBits, higher the target, easier it is to crack since you just have to find a value lower than target
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreviousHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func main() {
	bc := newBlockchain()
	bc.addBlock("Sent 1BTC to RA")
	bc.addBlock("Sent 4BTC to MMI")
	bc.addBlock("Seized 20BTC from ALS")
	bc.addBlock("Sent 5BTC to SI")
	bc.addBlock("Sent 2BTC to KN")

	for _, block := range bc.Blocks {
		printBlock(block)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
