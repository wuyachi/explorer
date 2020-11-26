package models

type Chain struct {
	Id     uint64 `gorm:"primaryKey;type:bigint(20);not null;auto_increment"`
	Name   string `gorm:"size:64"`
	Height uint64 `gorm:"type:bigint(20);not null;default:0"`
	StakeAmount uint64 `gorm:"type:bigint(20);not null;default:0"`
	LastReward uint64 `gorm:"type:bigint(20);not null;default:0"`
	MintPrice uint64 `gorm:"type:bigint(20);not null;default:0"`
	RewardPeriod uint64 `gorm:"type:bigint(20);not null;default:0"`
	GasFee uint64 `gorm:"type:bigint(20);not null;default:0"`
}

type Block struct {
	Hash    string `gorm:"primaryKey;size:66;not null"`
	GasLimit  uint64 `gorm:"type:bigint(20);not null;default:0"`
	GasUsed  uint64 `gorm:"type:bigint(20);not null;default:0"`
	Difficulty uint64 `gorm:"type:bigint(20);not null;default:0"`
	Number     uint64 `gorm:"type:bigint(20);not null;default:0"`
	ParentHash string `gorm:"size:66;not null"`
	ReceiptHash  string `gorm:"size:66;not null"`
	TxHash  string `gorm:"size:66;not null"`
	TxNumber uint64 `gorm:"type:bigint(20);not null;default:0"`
	Time  uint64 `gorm:"type:bigint(20);not null;default:0"`
	Transactions []*Transaction  `gorm:"foreignKey:BlockHash;references:Hash"`
}

type Transaction struct {
	Hash  string `gorm:"primaryKey;size:66;not null"`
	From  string `gorm:"size:42;not null"`
	Cost  uint64 `gorm:"type:bigint(20);not null;default:0"`
	Data  string `gorm:"size:4096;not null"`
	Gas   uint64 `gorm:"type:bigint(20);not null;default:0"`
	GasPrice uint64 `gorm:"type:bigint(20);not null;default:0"`
	To    string `gorm:"size:42;not null"`
	Value uint64 `gorm:"type:bigint(20);not null;default:0"`
	Time  uint64 `gorm:"type:bigint(20);not null;default:0"`
	BlockNumber  uint64 `gorm:"type:bigint(20);not null;default:0"`
	BlockHash  string `gorm:"size:66;not null"`
	Block  *Block `gorm:"foreignKey:BlockHash;references:Hash"`
	Events []*Event `gorm:"foreignKey:TransactionHash;references:Hash"`
	TransactionDetails []*TransactionDetail `gorm:"foreignKey:TransactionHash;references:Hash"`
}

type TransactionDetail struct {
	Contract string `gorm:"size:42;not null"`
	From  string `gorm:"size:42;not null"`
	To string `gorm:"size:42;not null"`
	Value string `gorm:"size:66;not null"`
	TransactionHash  string `gorm:"size:66;not null"`
	Transaction *Transaction `gorm:"foreignKey:TransactionHash;references:Hash"`
}

type Event struct {
	Number   uint64 `gorm:"type:bigint(20);not null;default:0"`
	Contract string `gorm:"size:42;not null"`
	EventId  string `gorm:"size:66;not null"`
	Topic1 string `gorm:"size:66;not null"`
	Topic2 string `gorm:"size:66;not null"`
	Topic3 string `gorm:"size:66;not null"`
	Topic4 string `gorm:"size:66;not null"`
	Data string `gorm:"size:4096;not null"`
	TransactionHash  string `gorm:"size:66;not null"`
	Transaction *Transaction `gorm:"foreignKey:TransactionHash;references:Hash"`
}

type PLTContract struct {
	Address string `gorm:"primaryKey;size:42;not null"`
	Amount uint64 `gorm:"type:bigint(20);not null;default:0"`
}

type NFTContract struct {
	NFT string `gorm:"primaryKey;size:42;not null"`
	Token string  `gorm:"primaryKey;size:42;not null"`
	Owner string `gorm:"size:42;not null"`
	Uri  string `gorm:"size:66;not null"`
}

type Stake struct {
	Owner string `gorm:"primaryKey;size:42;not null"`
	Validator string `gorm:"primaryKey;size:42;not null"`
	StakeAmount uint64 `gorm:"type:bigint(20);not null;default:0"`
}

type Validator struct {
	Address string `gorm:"primaryKey;size:42;not null"`
	DelegateFactor uint64 `gorm:"type:bigint(20);not null;default:0"`
	StakeAmount uint64 `gorm:"type:bigint(20);not null;default:0"`
}

type Propose struct {
	ProposeId    string `gorm:"primaryKey;size:42;not null"`
	Proposer     string `gorm:"size:42;not null"`
	Value        string `gorm:"size:66;not null"`
	EndBlock     uint64 `gorm:"type:bigint(20);not null;default:0"`
	Type         uint8 `gorm:"size:8;not null;default:0"`
	State        uint8 `gorm:"size:8;not null;default:0"`
}

