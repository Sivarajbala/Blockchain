package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"github.com/Sivarajbala/blocks"

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
