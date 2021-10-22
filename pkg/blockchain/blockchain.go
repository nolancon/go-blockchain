package blockchain

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func InitBlockchain(data string) *Blockchain {
	return &Blockchain{[]*Block{Genesis(data)}}
}

func Genesis(data string) *Block {
	return CreateBlock(data, []byte{})
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	block.Nonce, block.Hash = NewProofOfWork(block).Run()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	bc.Blocks = append(bc.Blocks, CreateBlock(data, prevBlock.Hash))
}

/*
func (b *Block) DeriveHash() {
	hash := sha256.Sum256(bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}))
	b.Hash = hash[:]
}
*/
