# palette explorer

提供palette链的信息。

## 测试节点

http://40.115.153.174:22000

## API

* [GET /](#get-/)
* [GET info](#get-info)
* [GET chain](#get-chain)
* [POST blockbynumber](#post-blockbynumber)
* [POST blockbyhash](#post-blockbyhash)
* [POST blocks](#post-blocks)
* [POST transactionbyhash](#post-transactionbyhash)
* [POST transactions](#post-transactions)
* [POST transactionsofcontract](#post-transactionsofcontract)
* [POST transactionsofuser](#post-transactionsofuser)
* [POST transactiondetails](#post-transactiondetails)
* [POST pltholders](#post-pltholders)
* [POST plttransactions](#post-plttransactions)
* [POST pltholderinfo](#post-pltholderinfo)
* [POST nfts](#post-nfts)
* [POST nft](#post-nft)
* [POST nfttokeninfo](#post-nfttokeninfo)
* [POST nfttokentransactions](#post-nfttokentransactions)
* [POST nftholdersofuser](#post-nftholdersofuser)
* [POST nftholders](#post-nftholders)
* [POST nftusers](#post-nftusers)
* [POST nfttransactions](#post-nfttransactions)
* [POST stakesofowner](#post-stakesofowner)
* [POST stakesofvalidator](#post-stakesofvalidator)
* [POST stakeinfo](#post-stakeinfo)
* [POST validators](#post-validators)
* [POST validatorinfo](#post-validatorinfo)
* [POST proposes](#post-proposes)
* [POST proposesofuser](#post-proposesofuser)
* [POST proposeinfo](#post-proposeinfo)

## API

### GET /

Request 
```
http://localhost:8080/
```

Example Request
```
curl --location --request GET 'http://localhost:8080/'
```

Example Response
```
{
    "Version": "v1",
    "URL": "http://localhost:8080/v1"
}
```

### GET info

Request 
```
http://localhost:8080/v1
```

Example Request
```
curl --location --request GET 'http://localhost:8080/v1'
```

Example Response
```
{
    "Version": "v1",
    "URL": "http://localhost:8080/v1"
}
```

### GET chain

Request 
```
http://localhost:8080/v1/chain
```

Example Request
```
curl --location --request GET 'http://localhost:8080/v1/chain'
```

Example Response
```
{
    "Id": 1,
    "Name": "palette",
    "Height": 42675,
    "StakeAmount": "150000000",
    "LastReward": 70515,
    "MintPrice": "0",
    "RewardPeriod": 5,
    "GasFee": "0"
}
```

### POST blockbynumber

Request 
```
http://localhost:8080/v1/blockbynumber
```

BODY raw
```
{
    "Number": 1
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/blockbynumber' \
--data-raw '{
    "Number": 1
}'
```

Example Response
```
{
    "Hash": "0x1a7397488ec4fb21bfa1d133194315e4d3601af9d9ce611ff64a5ffc47970338",
    "GasLimit": 3757178881,
    "Size": 1096,
    "Validators": 5,
    "GasUsed": 0,
    "Difficulty": 1,
    "Number": 1,
    "ParentHash": "0x2c515b3b0af55ba04df4ca18cbabcc8ac7b95a49b880af704b644b549b910bdd",
    "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
    "TxHash": "0x4b84ad186c799266885db34cdbfc714edfd2cd37833449734f9520b5c8333129",
    "TxNumber": 1,
    "Time": 1612335579,
    "Transactions": [
        "0x9b8b22fea998e8da751ab2971f597487d9bc9ec5cf78edc5af8bbc12a4eee528"
    ]
}
```

### POST blockbyhash

Request 
```
http://localhost:8080/v1/blockbyhash/
```

BODY raw
```
{
    "Hash": "0x1a7397488ec4fb21bfa1d133194315e4d3601af9d9ce611ff64a5ffc47970338"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/blockbyhash/' \
--data-raw '{
    "Hash": "0x1a7397488ec4fb21bfa1d133194315e4d3601af9d9ce611ff64a5ffc47970338"
}'
```

Example Response
```
{
    "Hash": "0x1a7397488ec4fb21bfa1d133194315e4d3601af9d9ce611ff64a5ffc47970338",
    "GasLimit": 3757178881,
    "Size": 1096,
    "Validators": 5,
    "GasUsed": 0,
    "Difficulty": 1,
    "Number": 1,
    "ParentHash": "0x2c515b3b0af55ba04df4ca18cbabcc8ac7b95a49b880af704b644b549b910bdd",
    "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
    "TxHash": "0x4b84ad186c799266885db34cdbfc714edfd2cd37833449734f9520b5c8333129",
    "TxNumber": 1,
    "Time": 1612335579,
    "Transactions": [
        "0x9b8b22fea998e8da751ab2971f597487d9bc9ec5cf78edc5af8bbc12a4eee528"
    ]
}
```

### POST blocks

Request 
```
http://localhost:8080/v1/blocks
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":5
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/blocks' \
--data-raw '{
    "PageNo":0,
    "PageSize":5
}'

```

Example Response

```
{
    "PageSize": 5,
    "PageNo": 0,
    "TotalPage": 8583,
    "TotalCount": 42915,
    "Blocks": [
        {
            "Hash": "0x17c500c150975e8414e2cba3c75f2afc3aad06f1511ae2d5435e58c30a7ed842",
            "GasLimit": 700000000,
            "Size": 1653,
            "Validators": 8,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 42913,
            "ParentHash": "0x4045e4d5cda541cfd3f689220779b4d1fad4c9adfb9ae8b2fcc7e7c58c36cfad",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0xf1db1ddfc4d08a6b1b5f261f1131af84f052dfdf47b97143fd00bb14485e5c7c",
            "TxNumber": 1,
            "Time": 1612551516,
            "Transactions": [
                "0x536468d1274c69227841b168b84defebfc4106944f523651b9daba81cf3f31ea"
            ]
        },
        {
            "Hash": "0xf5025b4dc6fcfd0e17e608f7259f8ac4706760e2b165e912ffebe9aca85a87cb",
            "GasLimit": 700000000,
            "Size": 1653,
            "Validators": 8,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 42912,
            "ParentHash": "0x2d9fe508df839f64b0833c51621dc2feba5fc65972ef1a86d365dd8df1113dd3",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x86a5dff31066ecfb52f232509005aba6fa470a6ac757b8e90477325f0234e3a3",
            "TxNumber": 1,
            "Time": 1612551511,
            "Transactions": [
                "0xbe61cb17e538cea210ba433b72549698478bc68e1d26852a4e632022a2b027b5"
            ]
        },
        {
            "Hash": "0x1e58e91b02f5d2316609a971f04fef093f404f0ee5730a0c03dc82c5f02939e9",
            "GasLimit": 700000000,
            "Size": 1653,
            "Validators": 8,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 42911,
            "ParentHash": "0x2ea3fc96a436e43e7a874936d1645b80af50a11f9d4cf2c788d907b9ce056437",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x07bb14529829c471e8f8676336e62accc5358535caf00f581a131987510f20c9",
            "TxNumber": 1,
            "Time": 1612551506,
            "Transactions": [
                "0xd0d33b9defb01d29b861661697a3a49526a411aee6d0a6c5c6c7b31d3a4e0083"
            ]
        },
        {
            "Hash": "0xc54e16e55614d928e28714da956e1182aa97129b3e2492ce2cccedb65d65b2fd",
            "GasLimit": 700000000,
            "Size": 1653,
            "Validators": 8,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 42910,
            "ParentHash": "0xdd205ddbd717a143dfa561a4cc9e453b6b2dca60ef02f754f851f336754f9075",
            "ReceiptHash": "0x87570201ab52828e9e768c4239e8cbb034ff3776a93fc363abb19814be7e1451",
            "TxHash": "0x07926dffd3f355af8e0ab2325cc0a0878256ecd59b67ba46864cc43d53d121da",
            "TxNumber": 1,
            "Time": 1612551501,
            "Transactions": [
                "0x169cb7757eab6716ccc11cfe838c6b3c1f787b37344d3633ac348597180ec9d4"
            ]
        },
        {
            "Hash": "0xe2cdb7e665488a9ec52e5f738732cd283d683637649fdbdf2f68c507804c4cfa",
            "GasLimit": 700000000,
            "Size": 1653,
            "Validators": 8,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 42909,
            "ParentHash": "0xe3c34a489efdb6b9c0d873ce24b45ce2769745d3071721e3841469c74b9fd180",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x706201ed4da4c034816af0f83587d660a426861b722e14752f860d20ac37d40f",
            "TxNumber": 1,
            "Time": 1612551496,
            "Transactions": [
                "0x11c92de0f7962d8c56b7ed02affadb61e9ac681c5006c6c317f238e91ef34cef"
            ]
        }
    ]
}
```

### POST transactionbyhash

Request 
```
http://localhost:8080/v1/transactionbyhash/
```

BODY raw
```
{
    "Hash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionbyhash/' \
--data-raw '{
    "Hash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0"
}'
```

Example Response
```
{
    "Hash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0",
    "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
    "Cost": 0,
    "Data": "23b872dd0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000000000000000000000000000000000000000000001",
    "Gas": 2100000,
    "GasPrice": "0",
    "To": "0x0000000000000000000000000000000000001001",
    "Value": "0",
    "Time": 1612434309,
    "BlockNumber": 19746,
    "Type": 1,
    "Status": 1,
    "BlockHash": "0x2744a9b14f8876aa9a926202197aba4aea2a46dc2d155537a1ae4930c87e6e5e",
    "Events": [
        {
            "Number": 19746,
            "Contract": "0x0000000000000000000000000000000000001001",
            "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
            "Topic1": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Topic2": "0x0000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c",
            "Topic3": "",
            "Topic4": "",
            "Data": "01",
            "Time": 1612434309,
            "TransactionHash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0"
        }
    ],
    "TransactionDetails": [
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "1",
            "Time": 1612434309,
            "Status": 1,
            "TransactionHash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0"
        }
    ]
}
```

### POST transactions

Request 
```
http://localhost:8080/v1/transactions
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":5
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactions' \
--data-raw '{
    "PageNo":0,
    "PageSize":5
}'
```

Example Response
```
{
    "PageSize": 5,
    "PageNo": 0,
    "TotalPage": 8873,
    "TotalCount": 44365,
    "Transactions": [
        {
            "Hash": "0x2b5a5068e2fd949bf64e294938c12f036f9cd6be0b9a4269eca8fae35b3b1076",
            "From": "0xbfb558f0dceb07fbb09e1c283048b551a4310921",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000aac0000000000000000000000000000000000000000000000000000000000000000800000000000000000000000003ff6beb65feb5da87ca1b5468b3e95da767255e000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000105",
            "Value": "0",
            "Time": 1612555511,
            "BlockNumber": 43712,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0x555b90e9513ccfd584220e1d99a3fbb833a87b3b1b01dc16dd0439e59ff5c51f"
        },
        {
            "Hash": "0x9917f53ec2e0de876e3c5abc379ecab2c5c823e1029616195f9a7bfc66e8af64",
            "From": "0xd47a4e56e9262543db39d9203cf1a2e53735f834",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000aabf000000000000000000000000000000000000000000000000000000000000000800000000000000000000000003ff6beb65feb5da87ca1b5468b3e95da767255e000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000105",
            "Value": "0",
            "Time": 1612555506,
            "BlockNumber": 43711,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0x4262b84cd3b844f31295b147d5e7823a0229d8573d35c447a4d5bd72df4ce1b0"
        },
        {
            "Hash": "0xa87c11a0dd775ef889e6def89b6a52dbda8a596e258b4e0a4241ca9c0f7a7d27",
            "From": "0xc095448424a5ecd5ca7ccdadfaad127a9d7e88ec",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000aabe000000000000000000000000000000000000000000000000000000000000000800000000000000000000000003ff6beb65feb5da87ca1b5468b3e95da767255e000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000105",
            "Value": "0",
            "Time": 1612555501,
            "BlockNumber": 43710,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xc83c430e74be761a270aa3f8957b4c3c57e1971a4738522b58d4c7534dc8aa14"
        },
        {
            "Hash": "0xc5c1ee9b1acd4e0824df3b545b701f9601d894dc5c60699962626b00ae9d3bde",
            "From": "0xad3bf5ed640cc72f37bd21d64a65c3c756e9c88c",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000aabd000000000000000000000000000000000000000000000000000000000000000800000000000000000000000003ff6beb65feb5da87ca1b5468b3e95da767255e000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000105",
            "Value": "0",
            "Time": 1612555496,
            "BlockNumber": 43709,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xa7f1dde9d852bf1ba39b9ce3fa2e8eac2ce2f498a078fc3529104a6685de4804"
        },
        {
            "Hash": "0x8b0fca00f3c6ad573e91d865f114ea5df9968ef34fae1211b1763b61a19544cd",
            "From": "0x8c09d936a1b408d6e0afaa537ba4e06c4504a0ae",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000aabc000000000000000000000000000000000000000000000000000000000000000800000000000000000000000003ff6beb65feb5da87ca1b5468b3e95da767255e000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000ad3bf5ed640cc72f37bd21d64a65c3c756e9c88c000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000105",
            "Value": "0",
            "Time": 1612555491,
            "BlockNumber": 43708,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0x289a995e2ec0c3f537a927ca7d4067dbf24afbaad9c6a0d702c4ca08bf558d61"
        }
    ]
}
```

### POST transactionsofcontract

Request 
```
http://localhost:8080/v1/transactionsofcontract
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":5,
    "Contract": "0x0000000000000000000000000000000000001001"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionsofcontract' \
--data-raw '{
    "PageNo":0,
    "PageSize":5,
    "Contract": "0x0000000000000000000000000000000000001001"
}'
```

Example Response
```
{
    "PageSize": 5,
    "PageNo": 0,
    "TotalPage": 2,
    "TotalCount": 9,
    "Transactions": [
        {
            "Hash": "0x39a019e2fd0e17168763dcb6c11ab06103a9f9c5a98bc260fb5ddddddacee0b3",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Cost": 0,
            "Data": "23b872dd0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000000000000000000000000000000000000000000003",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000001001",
            "Value": "0",
            "Time": 1612434369,
            "BlockNumber": 19758,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xe1afab662e1e3104762e750c4f8e9d88c617715ec79ec6ffce8ceba8105e962e"
        },
        {
            "Hash": "0xa16958915989fac051715356ce6f8a87f13acad230d68229a4de7852ebad93ec",
            "From": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Cost": 0,
            "Data": "23b872dd0000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000003",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000001001",
            "Value": "0",
            "Time": 1612434364,
            "BlockNumber": 19757,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xa5cb7ac0fcdb44e6fd8122046bae155f130314dd83e820642342d32be2d7fe0b"
        },
        {
            "Hash": "0x7683e2f61d6168d838841cb4e6f0bccdce4488b7ef2abad797498362eeb38058",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Cost": 0,
            "Data": "23b872dd0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000000000000000000000000000000000000000000002",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000001001",
            "Value": "0",
            "Time": 1612434344,
            "BlockNumber": 19753,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0x268a134067cc4bb836c5bf6db376dba304c6af2776a83508f645100cedb69c32"
        },
        {
            "Hash": "0xc9e62f19bbf4956db22117e4f80eb1dee0a421e0db7e95c7e9b1e7c02a2afd3b",
            "From": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Cost": 0,
            "Data": "23b872dd0000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000002",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000001001",
            "Value": "0",
            "Time": 1612434339,
            "BlockNumber": 19752,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xd50a9880b0f5c54d4dbf675e2b48141b6e91c4c25c02fb2584f31ddcfad5f633"
        },
        {
            "Hash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Cost": 0,
            "Data": "23b872dd0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000006a708455c8777630aac9d1e7702d13f7a865b27c0000000000000000000000000000000000000000000000000000000000000001",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000001001",
            "Value": "0",
            "Time": 1612434309,
            "BlockNumber": 19746,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0x2744a9b14f8876aa9a926202197aba4aea2a46dc2d155537a1ae4930c87e6e5e"
        }
    ]
}
```

### POST transactionsofuser

Request 
```
http://localhost:8080/v1/transactionsofuser
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":5,
    "User": "0xf3a9d42c01635a585f1721463842f8936075105f"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionsofuser' \
--data-raw '{
    "PageNo":0,
    "PageSize":5,
    "User": "0xf3a9d42c01635a585f1721463842f8936075105f"
}'
```

Example Response
```
{
    "PageSize": 5,
    "PageNo": 0,
    "TotalPage": 4,
    "TotalCount": 16,
    "Transactions": [
        {
            "Hash": "0xea8c1510cc92cdd8c842844adc25c7f1d854442d8157755f2138256fcbd42991",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "a9059cbb0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000de0b6b3a7640000",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000103",
            "Value": "0",
            "Time": 1612516381,
            "BlockNumber": 35886,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xed45057c0b8a3d193ca193100b020d8cc68a87287a1764e75c45ef3bc1c76cb7",
            "Events": [
                {
                    "Number": 35886,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "0de0b6b3a7640000",
                    "Time": 1612516381,
                    "TransactionHash": "0xea8c1510cc92cdd8c842844adc25c7f1d854442d8157755f2138256fcbd42991"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "1000000000",
                    "Time": 1612516381,
                    "Status": 1,
                    "TransactionHash": "0xea8c1510cc92cdd8c842844adc25c7f1d854442d8157755f2138256fcbd42991"
                }
            ]
        },
        {
            "Hash": "0x93c34d597f144a0192589ccbd0fe1f19b39a9b12944d5e3aefdb12c2c935c4f2",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "a9059cbb0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000de0b6b3a7640000",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000103",
            "Value": "0",
            "Time": 1612516006,
            "BlockNumber": 35811,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xcbfe35423da5fc64162a33a334d5bc5f686949ac62bea460956772fa58f11d2f",
            "Events": [
                {
                    "Number": 35811,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "0de0b6b3a7640000",
                    "Time": 1612516006,
                    "TransactionHash": "0x93c34d597f144a0192589ccbd0fe1f19b39a9b12944d5e3aefdb12c2c935c4f2"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "1000000000",
                    "Time": 1612516006,
                    "Status": 1,
                    "TransactionHash": "0x93c34d597f144a0192589ccbd0fe1f19b39a9b12944d5e3aefdb12c2c935c4f2"
                }
            ]
        },
        {
            "Hash": "0x87b438276fedcc58c016930dfb376e39a92132b05ee5faf4cbd00fd38c259230",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "a9059cbb0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000de0b6b3a7640000",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000103",
            "Value": "0",
            "Time": 1612418849,
            "BlockNumber": 16654,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xe06a7288d0d6208554bef3214c690a79b9a5f1048d2bfdfcf1502d8bd7088cc0",
            "Events": [
                {
                    "Number": 16654,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "0de0b6b3a7640000",
                    "Time": 1612418849,
                    "TransactionHash": "0x87b438276fedcc58c016930dfb376e39a92132b05ee5faf4cbd00fd38c259230"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "1000000000",
                    "Time": 1612418849,
                    "Status": 1,
                    "TransactionHash": "0x87b438276fedcc58c016930dfb376e39a92132b05ee5faf4cbd00fd38c259230"
                }
            ]
        },
        {
            "Hash": "0x4b7ee200c046dd204df89774d808b7a3ec098ad4febc83f5e67f4e56632d3c0c",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "a9059cbb0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000de0b6b3a7640000",
            "Gas": 2100000,
            "GasPrice": "0",
            "To": "0x0000000000000000000000000000000000000103",
            "Value": "0",
            "Time": 1612417944,
            "BlockNumber": 16473,
            "Type": 1,
            "Status": 1,
            "BlockHash": "0xbc39212b1eadebcde4d806aa6bb539d6e9018e13b100ef4c11792d0b4c78ec89",
            "Events": [
                {
                    "Number": 16473,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "0de0b6b3a7640000",
                    "Time": 1612417944,
                    "TransactionHash": "0x4b7ee200c046dd204df89774d808b7a3ec098ad4febc83f5e67f4e56632d3c0c"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "1000000000",
                    "Time": 1612417944,
                    "Status": 1,
                    "TransactionHash": "0x4b7ee200c046dd204df89774d808b7a3ec098ad4febc83f5e67f4e56632d3c0c"
                }
            ]
        },
        {
            "Hash": "0xfd82f4b8489999e83b02a960b5d34f23c39548f68cbc405e70a82023e8d90088",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "34a773eb0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000005cf00000000ffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000c25191cf7592b15c04ca1bdcb07677a1bf3c995353ee4e68e35f798ee83fdcee00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008e305f000000001dac2b7c00000000fd1a057b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a372c2263223a322c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303365663434626562613834343232626437366135393935333163396665353039363961393239613066656533356466363636393066333730636531396661386330227d2c7b22696e646578223a322c226964223a2231323035303338323437656663666561653066646637363036383564316163316330383362653366663565396134613534386263336132653938663034333466303932343833227d2c7b22696e646578223a332c226964223a2231323035303232303932653334653031373664636366386162623439366238333364353931643235353333343639623363616630653237396239373432393535646438666333227d2c7b22696e646578223a342c226964223a2231323035303237626437373165363861646238383339383238326532316138623033633132663634633233353165613439613262613036613033323763383362323339636139227d2c7b22696e646578223a352c226964223a2231323035303264306430653838336337336438323536636634333134383232646464393733633031373962373364386564336466383561616433386433366138623262306337227d2c7b22696e646578223a362c226964223a2231323035303361346634346464363563626363353262316431616335313734373337386137663834373533623566376266323736306361323133393063656436623137326262227d2c7b22696e646578223a372c226964223a2231323035303236393663306362653734663031656538356533633065626534656264633562656134303466313939643032363266313934316664333966663064313030323537227d5d2c22706f735f7461626c65223a5b362c362c312c352c322c322c372c332c362c352c362c322c372c322c372c322c362c362c322c312c352c312c312c352c332c372c372c332c332c332c342c352c332c332c342c322c352c342c352c322c312c372c372c372c312c332c342c362c372c342c372c362c352c332c352c372c312c322c332c352c372c342c342c362c332c332c342c322c352c352c322c322c342c342c332c312c352c362c342c322c372c322c312c362c312c342c362c312c332c312c372c332c312c312c342c362c352c342c362c362c352c372c312c342c325d2c226d61785f626c6f636b5f6368616e67655f76696577223a36303030307d7d3cc22b9403d96ee5c9422ca9d502e0907617ccb20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001d51205042092e34e0176dccf8abb496b833d591d25533469b3caf0e279b9742955dd8fc3899a042cd338e82698b5284720f85b309f2b711c05cb37836488371741168da6120504696c0cbe74f01ee85e3c0ebe4ebdc5bea404f199d0262f1941fd39ff0d100257a2f2a11aaf2f0baccf6c9e30aa3b204bd4b935f3c1bb5b20349c7afd35565f2e1205047bd771e68adb88398282e21a8b03c12f64c2351ea49a2ba06a0327c83b239ca9420cf3852f7991d2a53afd008d1f6c356294b83aeeb4aad769f8c95ffeb4d5ac1205048247efcfeae0fdf760685d1ac1c083be3ff5e9a4a548bc3a2e98f0434f092483760cb1d3138a9beadf9f784d60604f37f1a51464ba228ec44f89879df1c10e07120504a4f44dd65cbcc52b1d1ac51747378a7f84753b5f7bf2760ca21390ced6b172bbf4d03e2cf4e0e79e46f7a757058d240e542853341e88feb1610ff03ba785cfc1120504d0d0e883c73d8256cf4314822ddd973c0179b73d8ed3df85aad38d36a8b2b0c7696f0c66330d243b1bc7bc8d05e694b4d642ac68f741d2b7f6ea4037ef46b992120504ef44beba84422b",
            "Gas": 10000000,
            "GasPrice": "0",
            "To": "0x42fba453e4201df556b06d02400e944b33e98ddf",
            "Value": "0",
            "Time": 1612417274,
            "BlockNumber": 16339,
            "Type": 1,
            "Status": 0,
            "BlockHash": "0x651a7b80c98e77e9b108c4dfa7c03a91bbcb7ea154756d7a0db59a2b9d673c2d"
        }
    ]
}
```

### POST transactiondetails

Request 
```
http://localhost:8080/v1/transactiondetails
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":10,
    "Hash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactiondetails' \
--data-raw '{
    "PageNo":0,
    "PageSize":10,
    "Hash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0"
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 1,
    "TransactionDetails": [
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "1",
            "Time": 1612434309,
            "Status": 1,
            "TransactionHash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0",
            "ContractInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Type": 1,
                "Name": "JpDigitalCat01",
                "Symbol": "JDC-01",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "",
                "Site": "",
                "Time": 1612421689,
                "TotalSupply": 3,
                "AddressNum": 1,
                "TransferNum": 9
            },
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "1",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat1.jpeg"
            }
        }
    ]
}
```

### POST pltholders

Request 
```
http://localhost:8080/v1/pltholders
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":5
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/pltholders' \
--data-raw '{
    "PageNo":0,
    "PageSize":5
}'
```

Example Response
```
{
    "PageSize": 5,
    "PageNo": 0,
    "TotalPage": 2,
    "TotalCount": 8,
    "PLTHolderInfos": [
        {
            "Address": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Amount": "449999996",
            "Percent": "75%"
        },
        {
            "Address": "0x0000000000000000000000000000000000000105",
            "Amount": "149457060",
            "Percent": "24.9%"
        },
        {
            "Address": "0xf716a8e5c1324ae18181f7200b2da26038769782",
            "Amount": "144784",
            "Percent": "0.02%"
        },
        {
            "Address": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
            "Amount": "144784",
            "Percent": "0.02%"
        },
        {
            "Address": "0x8fc2680e989c39eca4544c5961fb2415d0e8d907",
            "Amount": "144784",
            "Percent": "0.02%"
        }
    ]
}
```

### POST plttransactions

Request 
```
http://localhost:8080/v1/plttransactions
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":5
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/plttransactions' \
--data-raw '{
    "PageNo":0,
    "PageSize":5
}'
```

Example Response
```
{
    "PageSize": 5,
    "PageNo": 0,
    "TotalPage": 7264,
    "TotalCount": 36318,
    "TransactionDetails": [
        {
            "Contract": "0x0000000000000000000000000000000000000103",
            "From": "0x0000000000000000000000000000000000000105",
            "To": "0x8fc2680e989c39eca4544c5961fb2415d0e8d907",
            "Value": "16000000000",
            "Time": 1612563901,
            "Status": 1,
            "TransactionHash": "0x27c5065b994510cd7f80733a29c17ba471e75fcdbd922f439f24183a1806b35b"
        },
        {
            "Contract": "0x0000000000000000000000000000000000000103",
            "From": "0x0000000000000000000000000000000000000105",
            "To": "0xa2ec66f9dee661e096db3c4de1187d32e48cc959",
            "Value": "12000000000",
            "Time": 1612563901,
            "Status": 1,
            "TransactionHash": "0x27c5065b994510cd7f80733a29c17ba471e75fcdbd922f439f24183a1806b35b"
        },
        {
            "Contract": "0x0000000000000000000000000000000000000103",
            "From": "0x0000000000000000000000000000000000000105",
            "To": "0xf716a8e5c1324ae18181f7200b2da26038769782",
            "Value": "16000000000",
            "Time": 1612563901,
            "Status": 1,
            "TransactionHash": "0x27c5065b994510cd7f80733a29c17ba471e75fcdbd922f439f24183a1806b35b"
        },
        {
            "Contract": "0x0000000000000000000000000000000000000103",
            "From": "0x0000000000000000000000000000000000000105",
            "To": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
            "Value": "16000000000",
            "Time": 1612563901,
            "Status": 1,
            "TransactionHash": "0x27c5065b994510cd7f80733a29c17ba471e75fcdbd922f439f24183a1806b35b"
        },
        {
            "Contract": "0x0000000000000000000000000000000000000103",
            "From": "0x0000000000000000000000000000000000000105",
            "To": "0x8fc2680e989c39eca4544c5961fb2415d0e8d907",
            "Value": "16000000000",
            "Time": 1612563876,
            "Status": 1,
            "TransactionHash": "0x331094ac0bdef2e3bfdb3708da9e43b8de5d995fffdef46c52ff7c9609cc1148"
        }
    ]
}
```

### POST pltholderinfo

Request 
```
http://localhost:8080/v1/pltholderinfo
```

BODY raw
```
{
    "Address": "0x0000000000000000000000000000000000000105"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/pltholderinfo' \
--data-raw '{
    "Address": "0x0000000000000000000000000000000000000105"
}'
```

Example Response
```
{
    "Address": "0x0000000000000000000000000000000000000105",
    "Amount": "149453700"
}
```

### POST nfts

Request 
```
http://localhost:8080/v1/nfts
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":5
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nfts' \
--data-raw '{
    "PageNo":0,
    "PageSize":5
}'
```

Example Response
```
{
    "PageSize": 5,
    "PageNo": 0,
    "TotalPage": 2,
    "TotalCount": 6,
    "Contracts": [
        {
            "Contract": "0x0000000000000000000000000000000000001005",
            "Type": 1,
            "Name": "JpDigitalCat05",
            "Symbol": "JDC-05",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "",
            "Site": "",
            "Time": 1612507094,
            "TotalSupply": 3,
            "AddressNum": 2,
            "TransferNum": 12
        },
        {
            "Contract": "0x0000000000000000000000000000000000001004",
            "Type": 1,
            "Name": "JpDigitalCat02",
            "Symbol": "JDC-02",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "",
            "Site": "",
            "Time": 1612507049,
            "TotalSupply": 0,
            "AddressNum": 0,
            "TransferNum": 0
        },
        {
            "Contract": "0x0000000000000000000000000000000000001003",
            "Type": 1,
            "Name": "JpDigitalCat02",
            "Symbol": "JDC-02",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "",
            "Site": "",
            "Time": 1612502119,
            "TotalSupply": 1,
            "AddressNum": 1,
            "TransferNum": 1
        },
        {
            "Contract": "0x0000000000000000000000000000000000001002",
            "Type": 1,
            "Name": "JpDigitalCat02",
            "Symbol": "JDC-02",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "",
            "Site": "",
            "Time": 1612434549,
            "TotalSupply": 3,
            "AddressNum": 2,
            "TransferNum": 3
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "Type": 1,
            "Name": "JpDigitalCat01",
            "Symbol": "JDC-01",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "",
            "Site": "",
            "Time": 1612421689,
            "TotalSupply": 3,
            "AddressNum": 1,
            "TransferNum": 9
        }
    ]
}
```

### POST nft

Request 
```
http://localhost:8080/v1/nft
```

BODY raw
```
{
   "Contract": "0x0000000000000000000000000000000000001001"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nft' \
--data-raw '{
    "Contract": "0x0000000000000000000000000000000000001001"
}'
```

Example Response
```
{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Type": 1,
    "Name": "JpDigitalCat01",
    "Symbol": "JDC-01",
    "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
    "Uri": "",
    "Site": "",
    "Time": 1612421689,
    "TotalSupply": 3,
    "AddressNum": 1,
    "TransferNum": 9
}
```

### POST nfttokeninfo

Request 
```
http://localhost:8080/v1/nfttokeninfo
```

BODY raw
```
{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Token": "1"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nfttokeninfo' \
--data-raw '{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Token": "1"
}'
```

Example Response
```
{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Token": "1",
    "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
    "Uri": "cat1.jpeg"
}
```


### POST nfttokentransactions

Request 
```
http://localhost:8080/v1/nfttokentransactions
```

BODY raw
```
{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Token": "1",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nfttokentransactions' \
--data-raw '{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Token": "1",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 3,
    "TransactionDetails": [
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "1",
            "Time": 1612434309,
            "Status": 1,
            "TransactionHash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "1",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat1.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Value": "1",
            "Time": 1612434299,
            "Status": 1,
            "TransactionHash": "0xeef1decdeb98a12ad4bba6e759c71b2b53b852b8b5edbcd927b9299d791a8814",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "1",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat1.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x0000000000000000000000000000000000000000",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "1",
            "Time": 1612427249,
            "Status": 1,
            "TransactionHash": "0xc449123ec72caea8fc25b8ad03b570287ac71da18a2eda1534337a076062d040",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "1",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat1.jpeg"
            }
        }
    ]
}
```

### POST nftholdersofuser

Request 
```
http://localhost:8080/v1/nftholdersofuser
```

BODY raw
```
{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Address": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nftholdersofuser' \
--data-raw '{
    "Contract": "0x0000000000000000000000000000000000001001",
    "Address": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 3,
    "NFTTokenInfos": [
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "Token": "1",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "cat1.jpeg"
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "Token": "2",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "cat2.jpeg"
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "Token": "3",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "cat3.jpeg"
        }
    ]
}
```

### POST nftholders

Request 
```
http://localhost:8080/v1/nftholders
```

BODY raw
```
{
    "Contract": "0x0000000000000000000000000000000000001001",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nftholders' \
--data-raw '{
    "Contract": "0x0000000000000000000000000000000000001001",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 3,
    "NFTTokenInfos": [
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "Token": "1",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "cat1.jpeg"
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "Token": "2",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "cat2.jpeg"
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "Token": "3",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Uri": "cat3.jpeg"
        }
    ]
}
```

### POST nfttransactions

Request 
```
http://localhost:8080/v1/nfttransactions
```

BODY raw
```
{
    "Contract": "0x0000000000000000000000000000000000001001",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nfttransactions' \
--data-raw '{
    "Contract": "0x0000000000000000000000000000000000001001",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 9,
    "TransactionDetails": [
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "3",
            "Time": 1612434369,
            "Status": 1,
            "TransactionHash": "0x39a019e2fd0e17168763dcb6c11ab06103a9f9c5a98bc260fb5ddddddacee0b3",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "3",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat3.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Value": "3",
            "Time": 1612434364,
            "Status": 1,
            "TransactionHash": "0xa16958915989fac051715356ce6f8a87f13acad230d68229a4de7852ebad93ec",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "3",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat3.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "2",
            "Time": 1612434344,
            "Status": 1,
            "TransactionHash": "0x7683e2f61d6168d838841cb4e6f0bccdce4488b7ef2abad797498362eeb38058",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "2",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat2.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Value": "2",
            "Time": 1612434339,
            "Status": 1,
            "TransactionHash": "0xc9e62f19bbf4956db22117e4f80eb1dee0a421e0db7e95c7e9b1e7c02a2afd3b",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "2",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat2.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "1",
            "Time": 1612434309,
            "Status": 1,
            "TransactionHash": "0x11e2014375403533664e1e7c92854e654eb375a83ad28d1239e39f80a84e4cc0",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "1",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat1.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "Value": "1",
            "Time": 1612434299,
            "Status": 1,
            "TransactionHash": "0xeef1decdeb98a12ad4bba6e759c71b2b53b852b8b5edbcd927b9299d791a8814",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "1",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat1.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x0000000000000000000000000000000000000000",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "1",
            "Time": 1612427249,
            "Status": 1,
            "TransactionHash": "0xc449123ec72caea8fc25b8ad03b570287ac71da18a2eda1534337a076062d040",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "1",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat1.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x0000000000000000000000000000000000000000",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "3",
            "Time": 1612427124,
            "Status": 1,
            "TransactionHash": "0x1548a65af2c7205ff5b345238b7e0555864d2da1e9a0130e50384e69c3bc748c",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "3",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat3.jpeg"
            }
        },
        {
            "Contract": "0x0000000000000000000000000000000000001001",
            "From": "0x0000000000000000000000000000000000000000",
            "To": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "Value": "2",
            "Time": 1612427049,
            "Status": 1,
            "TransactionHash": "0xdeab8611e43a4fb6a7de9115da057e6d048e7b11bbfacd87e1ad53c2b0720432",
            "TokenInfo": {
                "Contract": "0x0000000000000000000000000000000000001001",
                "Token": "2",
                "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
                "Uri": "cat2.jpeg"
            }
        }
    ]
}
```

### POST nftusers

Request 
```
http://localhost:8080/v1/nftusers
```

BODY raw
```
{
    "Contract": "0x0000000000000000000000000000000000001002",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nftusers' \
--data-raw '{
    "Contract": "0x0000000000000000000000000000000000001002",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 2,
    "NFTUsers": [
        {
            "Contract": "0x0000000000000000000000000000000000001002",
            "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "TokenNumber": 2,
            "Percent": "66.66%"
        },
        {
            "Contract": "0x0000000000000000000000000000000000001002",
            "Owner": "0x6a708455c8777630aac9d1e7702d13f7a865b27c",
            "TokenNumber": 1,
            "Percent": "33.32%"
        }
    ]
}
```

### POST stakesofowner

Request 
```
http://localhost:8080/v1/stakesofowner
```

BODY raw
```
{
    "Owner": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/stakesofowner' \
--data-raw '{
    "Owner": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 1,
    "StakeInfos": [
        {
            "Owner": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
            "Validator": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
            "Amount": 50000000000000000
        }
    ]
}
```

### POST stakesofvalidator

Request 
```
http://localhost:8080/v1/stakesofvalidator
```

BODY raw
```
{
    "Validator": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/stakesofvalidator' \
--data-raw '{
    "Validator": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 1,
    "StakeInfos": [
        {
            "Owner": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
            "Validator": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
            "Amount": 50000000000000000
        }
    ]
}
```

### POST stakeinfo

Request 
```
http://localhost:8080/v1/stakeinfo
```

BODY raw
```
{
    "Owner": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
    "Validator": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/stakeinfo' \
--data-raw '{
    "Owner": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
    "Validator": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e"
}'
```

Example Response
```
{
    "Owner": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
    "Validator": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
    "Amount": 50000000000000000
}
```

### POST validators

Request 
```
http://localhost:8080/v1/validators
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/validators' \
--data-raw '{
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 3,
    "ValidatorInfos": [
        {
            "Address": "0x03fF6Beb65fEb5Da87cA1B5468b3e95da767255E",
            "DelegateFactor": 2000,
            "StakeAmount": "50000000",
            "Name": "",
            "Uri": "",
            "Percent": "33.32%"
        },
        {
            "Address": "0x6A708455C8777630AaC9d1e7702d13F7a865b27C",
            "DelegateFactor": 2000,
            "StakeAmount": "50000000",
            "Name": "",
            "Uri": "",
            "Percent": "33.32%"
        },
        {
            "Address": "0xAd3Bf5eD640CC72f37BD21d64A65C3C756e9C88C",
            "DelegateFactor": 2000,
            "StakeAmount": "50000000",
            "Name": "",
            "Uri": "",
            "Percent": "33.32%"
        }
    ]
}
```

### POST validatorinfo

Request 
```
http://localhost:8080/v1/validatorinfo
```

BODY raw
```
{
    "Address": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/validatorinfo' \
--data-raw '{
    "Address": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e"
}'
```

Example Response
```
{
    "Address": "0x03fF6Beb65fEb5Da87cA1B5468b3e95da767255E",
    "DelegateFactor": 2000,
    "StakeAmount": "50000000",
    "Name": "",
    "Uri": ""
}
```

### POST proposes

Request 
```
http://localhost:8080/v1/proposes
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/proposes' \
--data-raw '{
    "PageNo":0,
    "PageSize":10
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 2,
    "ProposeInfos": [
        {
            "ProposeId": "0x0000000000000000000000010000000000000000",
            "Proposer": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
            "Value": "10000000000000000000",
            "EndBlock": 245591,
            "Type": 1,
            "State": 0
        },
        {
            "ProposeId": "0x0000000000000000000000010000000000000001",
            "Proposer": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
            "Value": "10000000000000000000",
            "EndBlock": 245646,
            "Type": 1,
            "State": 0
        }
    ]
}
```

### POST proposesofuser

Request 
```
http://localhost:8080/v1/proposesofuser
```

BODY raw
```
{
    "PageNo":0,
    "PageSize":10,
    "Proposer": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/proposesofuser' \
--data-raw '{
    "PageNo":0,
    "PageSize":10,
    "Proposer": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6"
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 1,
    "TotalCount": 2,
    "ProposeInfos": [
        {
            "ProposeId": "0x0000000000000000000000010000000000000000",
            "Proposer": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
            "Value": "10000000000000000000",
            "EndBlock": 245591,
            "Type": 1,
            "State": 0
        },
        {
            "ProposeId": "0x0000000000000000000000010000000000000001",
            "Proposer": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
            "Value": "10000000000000000000",
            "EndBlock": 245646,
            "Type": 1,
            "State": 0
        }
    ]
}
```

### POST proposeinfo

Request 
```
http://localhost:8080/v1/proposeinfo
```

BODY raw
```
{
    "ProposeId": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/proposeinfo' \
--data-raw '{
    "ProposeId": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e"
}'
```

Example Response
```
{
    "ProposeId": "0x0000000000000000000000010000000000000001",
    "Proposer": "0x03ff6beb65feb5da87ca1b5468b3e95da767255e",
    "Value": "10000000000000000000",
    "EndBlock": 245646,
    "Type": 1,
    "State": 0
}
```