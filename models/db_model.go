package models

type Chain struct {
	Id           uint64 `gorm:"primaryKey;type:bigint(20);not null;auto_increment"`
	Name         string `gorm:"size:64"`
	Height       uint64 `gorm:"type:bigint(20);not null"`
	StakeAmount  uint64 `gorm:"type:bigint(20);not null"`
	LastReward   uint64 `gorm:"type:bigint(20);not null"`
	MintPrice    uint64 `gorm:"type:bigint(20);not null"`
	RewardPeriod uint64 `gorm:"type:bigint(20);not null"`
	GasFee       uint64 `gorm:"type:bigint(20);not null"`
}

type Block struct {
	Hash         string         `gorm:"primaryKey;size:66;not null"`
	GasLimit     uint64         `gorm:"type:bigint(20);not null"`
	Size         uint64         `gorm:"type:bigint(20);not null"`
	Validators   uint64         `gorm:"type:bigint(20);not null"`
	GasUsed      uint64         `gorm:"type:bigint(20);not null"`
	Difficulty   uint64         `gorm:"type:bigint(20);not null"`
	Number       uint64         `gorm:"type:bigint(20);uniqueIndex;not null"`
	ParentHash   string         `gorm:"size:66;not null"`
	ReceiptHash  string         `gorm:"size:66;not null"`
	TxHash       string         `gorm:"size:66;not null"`
	TxNumber     uint64         `gorm:"type:bigint(20);not null0"`
	Time         uint64         `gorm:"type:bigint(20);not null"`
	Transactions []*Transaction `gorm:"foreignKey:BlockHash;references:Hash"`
}

type Transaction struct {
	Hash               string               `gorm:"primaryKey;size:66;not null"`
	From               string               `gorm:"size:42;not null"`
	Cost               uint64               `gorm:"type:bigint(20);not null"`
	Data               string               `gorm:"size:4096;not null"`
	Gas                uint64               `gorm:"type:bigint(20);not null"`
	GasPrice           uint64               `gorm:"type:bigint(20);not null"`
	To                 string               `gorm:"size:42;not null"`
	Value              uint64               `gorm:"type:bigint(20);not null"`
	Time               uint64               `gorm:"type:bigint(20);not null"`
	BlockNumber        uint64               `gorm:"type:bigint(20);index;not null"`
	Type               uint64               `gorm:"type:bigint(20);not null"`
	Status             uint64               `gorm:"type:bigint(20);not null"`
	BlockHash          string               `gorm:"size:66;not null"`
	Block              *Block               `gorm:"foreignKey:BlockHash;references:Hash"`
	Events             []*Event             `gorm:"foreignKey:TransactionHash;references:Hash"`
	TransactionDetails []*TransactionDetail `gorm:"foreignKey:TransactionHash;references:Hash"`
}

type TransactionDetail struct {
	Contract        string       `gorm:"size:42;not null"`
	From            string       `gorm:"size:42;not null"`
	To              string       `gorm:"size:42;not null"`
	Value           string       `gorm:"size:66;not null"`
	Time            uint64       `gorm:"type:bigint(20);not null"`
	TransactionHash string       `gorm:"size:66;not null"`
	Transaction     *Transaction `gorm:"foreignKey:TransactionHash;references:Hash"`
}

type TransactionDetailWithInfo struct {
	Contract        string        `gorm:"size:42;not null"`
	ContractInfo    *ContractInfo `gorm:"foreignKey:Contract;references:Contract"`
	From            string        `gorm:"size:42;not null"`
	To              string        `gorm:"size:42;not null"`
	Value           string        `gorm:"size:66;not null"`
	NFTHolder       *NFTHolder    `gorm:"foreignKey:Contract,Value;references:NFT,Token"`
	Time            uint64        `gorm:"type:bigint(20);not null"`
	TransactionHash string        `gorm:"size:66;not null"`
	Transaction     *Transaction  `gorm:"foreignKey:TransactionHash;references:Hash"`
}

func (TransactionDetailWithInfo) TableName() string {
	return "transaction_details"
}

type Event struct {
	Number          uint64       `gorm:"type:bigint(20);not null;default:0"`
	Contract        string       `gorm:"size:42;not null"`
	EventId         string       `gorm:"size:66;not null"`
	Topic1          string       `gorm:"size:66;not null"`
	Topic2          string       `gorm:"size:66;not null"`
	Topic3          string       `gorm:"size:66;not null"`
	Topic4          string       `gorm:"size:66;not null"`
	Data            string       `gorm:"size:4096;not null"`
	Time            uint64       `gorm:"type:bigint(20);not null"`
	TransactionHash string       `gorm:"size:66;not null"`
	Transaction     *Transaction `gorm:"foreignKey:TransactionHash;references:Hash"`
}

type PLTHolder struct {
	Address string `gorm:"primaryKey;size:42;not null"`
	Amount  uint64 `gorm:"type:bigint(20);not null"`
}

type PLTHolderWithPercent struct {
	Address string  `gorm:"primaryKey;size:42;not null"`
	Amount  uint64  `gorm:"type:bigint(20);not null"`
	Percent float64 `gorm:"not null"`
}

func (PLTHolderWithPercent) TableName() string {
	return "plt_holders"
}

type NFTHolder struct {
	NFT   string `gorm:"primaryKey;size:42;not null"`
	Token string `gorm:"primaryKey;size:42;not null"`
	Owner string `gorm:"size:42;not null"`
	Uri   string `gorm:"size:66;not null"`
}

type NFTUser struct {
	NFT         string  `gorm:"primaryKey;size:42;not null"`
	Owner       string  `gorm:"size:42;not null"`
	TokenNumber uint64  `gorm:"type:bigint(20);not null"`
	Percent     float64 `gorm:"not null"`
}

func (NFTUser) TableName() string {
	return "nft_holders"
}

type ContractInfo struct {
	Contract    string `gorm:"primaryKey;size:42;not null"`
	Type        uint64 `gorm:"type:bigint(20);not null"`
	Name        string `gorm:"size:42;not null"`
	Symbol      string `gorm:"size:42;not null"`
	Owner       string `gorm:"size:42;not null"`
	Uri         string `gorm:"size:66;not null"`
	Site        string `gorm:"size:66;not null"`
	Time        uint64 `gorm:"type:bigint(20);not null"`
	TotalSupply uint64 `gorm:"type:bigint(20);not null"`
	AddressNum  uint64 `gorm:"type:bigint(20);not null"`
	TransferNum uint64 `gorm:"type:bigint(20);not null"`
}

type Stake struct {
	Owner       string `gorm:"primaryKey;size:42;not null"`
	Validator   string `gorm:"primaryKey;size:42;not null"`
	StakeAmount uint64 `gorm:"type:bigint(20);not null"`
}

type Validator struct {
	Address        string `gorm:"primaryKey;size:42;not null"`
	DelegateFactor uint64 `gorm:"type:bigint(20);not null"`
	StakeAmount    uint64 `gorm:"type:bigint(20);not null"`
	Name           string `gorm:"size:66;not null"`
	Uri            string `gorm:"size:66;not null"`
}

type ValidatorWithPercent struct {
	Address        string  `gorm:"primaryKey;size:42;not null"`
	DelegateFactor uint64  `gorm:"type:bigint(20);not null"`
	StakeAmount    uint64  `gorm:"type:bigint(20);not null"`
	Name           string  `gorm:"size:66;not null"`
	Uri            string  `gorm:"size:66;not null"`
	Percent        float64 `gorm:"not null"`
}

func (ValidatorWithPercent) TableName() string {
	return "validators"
}

type Propose struct {
	ProposeId string `gorm:"primaryKey;size:42;not null"`
	Proposer  string `gorm:"size:42;not null"`
	Value     string `gorm:"size:66;not null"`
	EndBlock  uint64 `gorm:"type:bigint(20);not null"`
	Type      uint8  `gorm:"size:8;not null;default:0"`
	State     uint8  `gorm:"size:8;not null;default:0"`
}
