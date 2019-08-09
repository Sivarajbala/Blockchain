package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"encoding/hex"
)

var id int=0

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Iter     int
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	fmt.Println(string(info))
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{id,[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	block.Iter= id
	return block
}

func (chain *BlockChain) AddBlock(b *Block){
 	 chain.blocks = append(chain.blocks, b)
}
func (chain *BlockChain) NewBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	id = id+1
	new := CreateBlock(data, prevBlock.Hash)
	chain.AddBlock(new)
}

func Start() *Block {
	return CreateBlock("Starting Node", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Start()}}
}

func main() {
	chain := InitBlockChain()
	fmt.Println("********************************")
	chain.NewBlock("Sivaraj")
	chain.NewBlock("This is II")
	chain.NewBlock("Third Block")

	for _, block := range chain.blocks {
		fmt.Println("")
		fmt.Printf("Block Id: %d\n", block.Iter)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

	HashInBytes := sha256.Sum256([]byte("Starting Node"))
	HashInString := hex.EncodeToString(HashInBytes[:])
	fmt.Println(HashInString)
}

