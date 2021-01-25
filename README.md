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
* [POST pltholders](#post-pltholders)
* [POST pltholderinfo](#post-pltholderinfo)
* [POST nfts](#post-nfts)
* [POST nfttokens](#post-nfttokens)
* [POST nfttokeninfo](#post-nfttokeninfo)
* [POST nftholdersofuser](#post-nftholdersofuser)
* [POST nftholders](#post-nftholders)
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
    "Height": 698,
    "StakeAmount": 150000000000000000,
    "LastReward": 695,
    "MintPrice": 0,
    "RewardPeriod": 5,
    "GasFee": 0
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
    "Hash": "0x13ad834f28b5860d46b36372de679178a5ae668760b9933303016e3f5ab239ea",
    "GasLimit": 3757178881,
    "GasUsed": 0,
    "Difficulty": 1,
    "Number": 1,
    "ParentHash": "0x6df81bb3d4b48a834ad0dacec41c95d3373e861a1bfffdcaa37f5aae4ee6de56",
    "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
    "TxHash": "0x4b84ad186c799266885db34cdbfc714edfd2cd37833449734f9520b5c8333129",
    "TxNumber": 1,
    "Time": 1606357620,
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
    "Hash": "0x13ad834f28b5860d46b36372de679178a5ae668760b9933303016e3f5ab239ea"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/blockbyhash/' \
--data-raw '{
    "Hash": "0x13ad834f28b5860d46b36372de679178a5ae668760b9933303016e3f5ab239ea"
}'
```

Example Response
```
{
    "Hash": "0x13ad834f28b5860d46b36372de679178a5ae668760b9933303016e3f5ab239ea",
    "GasLimit": 3757178881,
    "GasUsed": 0,
    "Difficulty": 1,
    "Number": 1,
    "ParentHash": "0x6df81bb3d4b48a834ad0dacec41c95d3373e861a1bfffdcaa37f5aae4ee6de56",
    "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
    "TxHash": "0x4b84ad186c799266885db34cdbfc714edfd2cd37833449734f9520b5c8333129",
    "TxNumber": 1,
    "Time": 1606357620,
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
    "PageNo":1,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/blocks' \
--data-raw '{
    "PageNo":1,
    "PageSize":10
}'

```

Example Response

```
{
    "PageSize": 10,
    "PageNo": 1,
    "TotalPage": 70,
    "TotalCount": 698,
    "Blocks": [
        {
            "Hash": "0xc74d139d947123100ee90387230922200da266c1afd4e3125093bebeec0968d1",
            "GasLimit": 3748016166,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 11,
            "ParentHash": "0x7d6b514b400daff394390b0e76158b2bfcea27d8ae821f8821059e2ed56a1645",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x124af027d3690993c381510c2f2e1c26d2dbfa0b60bb5c53b55dfdc8b108e0a9",
            "TxNumber": 1,
            "Time": 1606357675,
            "Transactions": [
                "0xd89d75a2906230f2f3d70e5b2d2ad504196b5af6f25078c041b9bb8d79b06487"
            ]
        },
        {
            "Hash": "0xa4f2870bda7cdc9a925ac2d4fc34ca6c1899c0b8237bf69680da09516add8b96",
            "GasLimit": 3747101124,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 12,
            "ParentHash": "0x7f9f5de29b5bcc956e1aa3113265504bab2761c29bbcedfc74e41a455aeeb95b",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x6944fe35714b72916c4f729178fdcfa229054366a82dc90455d7a08dd5ead52c",
            "TxNumber": 1,
            "Time": 1606357680,
            "Transactions": [
                "0xe71a6882536b1cd4bfe36e97c2c77ca6b01fb21b5498f28a78db9cf5e25f47e1"
            ]
        },
        {
            "Hash": "0xe2aa04f829cebae8d00a466491df8edff877a79cb840e47ef2cf20f2da6bad52",
            "GasLimit": 3746186306,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 13,
            "ParentHash": "0x2f68c0c877abc526755b92412efb5eb9b4dd32f5bc36a1ebca534f9cf9ebfda6",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x390e3554ea77d343bc0923ad9b74461c8956df1685d1034256ce045e8dfed64e",
            "TxNumber": 1,
            "Time": 1606357685,
            "Transactions": [
                "0x116d35309c411cbb088c703763cc3e45b28488fe38899c3c2ae2a813aaf3660f"
            ]
        },
        {
            "Hash": "0x72c5d73e6260028a2e42d56998db6d1e2ce0ec6fe38ac298ae93943601c83ab4",
            "GasLimit": 3745271711,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 14,
            "ParentHash": "0x3b666fbc004b0a7cf9e1f354b9ee5eb916b5c25a92780882b1262c585f4ca70e",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x122bbcf37083da036ade35f6335500a79f3c484f2e475688a31aa1fb5ebbcf0e",
            "TxNumber": 1,
            "Time": 1606357690,
            "Transactions": [
                "0x23236a99be2bb9b4f65a4cfeb9dfd43d942534726d0d237321fe19c8e0fbc941"
            ]
        },
        {
            "Hash": "0xe1a1a2a3aa54b603858df4c01da35585260fc10cc04a6eca7a520d4d88523693",
            "GasLimit": 3744357340,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 15,
            "ParentHash": "0x622502fed0216d06d6d8aad685ac6dfcd0a570b96972b9a373b357b4bc473fca",
            "ReceiptHash": "0xcc347bee4f6552d2a702bbc1f2588043cfc17c0f2299f9ce76557c04b5c01f30",
            "TxHash": "0xde759f1b69441fbb8fc1f53eee1605f946db05603e4beaa8cdc7c7212debd0b7",
            "TxNumber": 1,
            "Time": 1606357695,
            "Transactions": [
                "0x1d8a144e211e01c0768f0847259277b330f6119c1e15a8409af0b186549761ae"
            ]
        },
        {
            "Hash": "0x170f51267738edd2b51d43c345c479a294a2f92515de48cc587ef3f67ce68a80",
            "GasLimit": 3743443192,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 16,
            "ParentHash": "0x2c646d417fe1c63b75f1e978ac911f96ac188f9612283526cebbf13fdfab56f0",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x58c5c993e5a6ebd9555fb2210f782062dcca9528d5fd42167f6d22f6f55fa685",
            "TxNumber": 1,
            "Time": 1606357700,
            "Transactions": [
                "0xff0fac7c5fe44f4a0f295c0265fbb1af01c10ba15bb8fccb8e8b64549eb72c3e"
            ]
        },
        {
            "Hash": "0x0f72e1e670f0032a98ac004a396a8ee253da9d07cf0b974ee236bec5349ba555",
            "GasLimit": 3742529267,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 17,
            "ParentHash": "0xe06f4e90e4d728b01ba6bd4f3cd6092920075b95960ed47c5bd167a7dc035d36",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0xa7c7ea87a3ed1fed5f449558a4c40050ac3b7a8dfb289c06c598daef315c5587",
            "TxNumber": 1,
            "Time": 1606357705,
            "Transactions": [
                "0x58c3423e11dc05c4403f4aa116a686f522026498dc8bd1453e022f4f7ed532d3"
            ]
        },
        {
            "Hash": "0x91ad4c065600cf7bbaa63cf9a62fab74d5f67da246b1873ad6071c0e593f3aaa",
            "GasLimit": 3741615565,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 18,
            "ParentHash": "0x8312b9e1f758758f32aa9e77423b8b1e833b03aa560fbee1a16f8e14f51d5ba3",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x864cbf5c3723cc11a786e7095cd0758acbb434279f98f1825883a7198349d3f3",
            "TxNumber": 1,
            "Time": 1606357710,
            "Transactions": [
                "0x611a4aee28226363b36f074b5b8ed655b268fa464dd8a8fe38567f4770aeb2b4"
            ]
        },
        {
            "Hash": "0x8cac9330e5c3446f2f782219d70704f4ebb2ebe558a53cefe029ee7a5a82e412",
            "GasLimit": 3740702086,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 19,
            "ParentHash": "0x2e6efebef9ea66eb3d561a9bc47d2ccc73645947e7757d31123a7a1a40241266",
            "ReceiptHash": "0xb64408da6b8fe39ab764af88ece1e8cca1c35fd988db57806e99138c629365a0",
            "TxHash": "0x5b2afb15cfded9ac134db95e8c39d12b28c795058c089a1cc3c096cc26644574",
            "TxNumber": 1,
            "Time": 1606357715,
            "Transactions": [
                "0xc50786ef7108a51059c6dd7724c1c51f5ee198f13171e3fa4e08ec9f945dcc65"
            ]
        },
        {
            "Hash": "0x90789a2c32644a684d5936dd9cfc012b14f6de5f769a8a32016aca0b5ee799fd",
            "GasLimit": 3739788830,
            "GasUsed": 0,
            "Difficulty": 1,
            "Number": 20,
            "ParentHash": "0xc1acf6f9706ca888b5976aad7e666c341c57957ae46932093eeaee6f4d18411b",
            "ReceiptHash": "0xcc347bee4f6552d2a702bbc1f2588043cfc17c0f2299f9ce76557c04b5c01f30",
            "TxHash": "0x394f8d254cd993bb15e594f6e9f08d8d4b09f8d5c8142f4a49bec24c4a8a4b27",
            "TxNumber": 1,
            "Time": 1606357720,
            "Transactions": [
                "0x75001c84f2a7c99733bf91ffb8f80ca89b9d29a9c295230c9995cdc5b75cc3b6"
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
    "Hash": "0x856c56562ca99d61ff1780745d13b5338c5d3278bea94ef3ac987a66d3a5bdc3"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionbyhash/' \
--data-raw '{
    "Hash": "0x856c56562ca99d61ff1780745d13b5338c5d3278bea94ef3ac987a66d3a5bdc3"
}'
```

Example Response
```
{
    "Hash": "0x856c56562ca99d61ff1780745d13b5338c5d3278bea94ef3ac987a66d3a5bdc3",
    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
    "Cost": 0,
    "Data": "a9059cbb0000000000000000000000008fc2680e989c39eca4544c5961fb2415d0e8d907000000000000000000000000000000000000000000295be96e64066972000000",
    "Gas": 2100000,
    "GasPrice": 0,
    "To": "0x0000000000000000000000000000000000000103",
    "Value": 0,
    "Time": 1606357775,
    "BlockNumber": 31,
    "BlockHash": "0x65057996d9d7b0b0309d1f2ceb01af4a3361b5a94ef9cb63ab799717400561c8",
    "Events": [
        {
            "Number": 31,
            "Contract": "0x0000000000000000000000000000000000000103",
            "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
            "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
            "Topic2": "0x0000000000000000000000008fc2680e989c39eca4544c5961fb2415d0e8d907",
            "Topic3": "",
            "Topic4": "",
            "Data": "295be96e64066972000000",
            "TransactionHash": "0x856c56562ca99d61ff1780745d13b5338c5d3278bea94ef3ac987a66d3a5bdc3"
        }
    ],
    "TransactionDetails": [
        {
            "Contract": "0x0000000000000000000000000000000000000103",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "To": "0x8fc2680e989c39eca4544c5961fb2415d0e8d907",
            "Value": "50000000000000000",
            "TransactionHash": "0x856c56562ca99d61ff1780745d13b5338c5d3278bea94ef3ac987a66d3a5bdc3"
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
    "PageNo":1,
    "PageSize":20
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactions' \
--data-raw '{
    "PageNo":1,
    "PageSize":20
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 1,
    "TotalPage": 75,
    "TotalCount": 748,
    "Transactions": [
        {
            "Hash": "0xd89d75a2906230f2f3d70e5b2d2ad504196b5af6f25078c041b9bb8d79b06487",
            "From": "0x8c09d936a1b408d6e0afaa537ba4e06c4504a0ae",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357675,
            "BlockNumber": 11,
            "BlockHash": "0xc74d139d947123100ee90387230922200da266c1afd4e3125093bebeec0968d1",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0xe71a6882536b1cd4bfe36e97c2c77ca6b01fb21b5498f28a78db9cf5e25f47e1",
            "From": "0xc095448424a5ecd5ca7ccdadfaad127a9d7e88ec",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357680,
            "BlockNumber": 12,
            "BlockHash": "0xa4f2870bda7cdc9a925ac2d4fc34ca6c1899c0b8237bf69680da09516add8b96",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x116d35309c411cbb088c703763cc3e45b28488fe38899c3c2ae2a813aaf3660f",
            "From": "0xd47a4e56e9262543db39d9203cf1a2e53735f834",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000d0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357685,
            "BlockNumber": 13,
            "BlockHash": "0xe2aa04f829cebae8d00a466491df8edff877a79cb840e47ef2cf20f2da6bad52",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x23236a99be2bb9b4f65a4cfeb9dfd43d942534726d0d237321fe19c8e0fbc941",
            "From": "0xbfb558f0dceb07fbb09e1c283048b551a4310921",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357690,
            "BlockNumber": 14,
            "BlockHash": "0x72c5d73e6260028a2e42d56998db6d1e2ce0ec6fe38ac298ae93943601c83ab4",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x1d8a144e211e01c0768f0847259277b330f6119c1e15a8409af0b186549761ae",
            "From": "0x258af48e28e4a6846e931ddff8e1cdf8579821e5",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000f0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357695,
            "BlockNumber": 15,
            "BlockHash": "0xe1a1a2a3aa54b603858df4c01da35585260fc10cc04a6eca7a520d4d88523693",
            "Events": [
                {
                    "Number": 15,
                    "Contract": "0x0000000000000000000000000000000000000105",
                    "EventId": "0xa9fb763ca024083f3b32979bc27ec1ff7cb648a232ed0b591cb9d3a7bb1374ad",
                    "Topic1": "",
                    "Topic2": "",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0x1d8a144e211e01c0768f0847259277b330f6119c1e15a8409af0b186549761ae"
                }
            ],
            "TransactionDetails": null
        },
        {
            "Hash": "0xff0fac7c5fe44f4a0f295c0265fbb1af01c10ba15bb8fccb8e8b64549eb72c3e",
            "From": "0x8c09d936a1b408d6e0afaa537ba4e06c4504a0ae",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357700,
            "BlockNumber": 16,
            "BlockHash": "0x170f51267738edd2b51d43c345c479a294a2f92515de48cc587ef3f67ce68a80",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x58c3423e11dc05c4403f4aa116a686f522026498dc8bd1453e022f4f7ed532d3",
            "From": "0xc095448424a5ecd5ca7ccdadfaad127a9d7e88ec",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000110000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357705,
            "BlockNumber": 17,
            "BlockHash": "0x0f72e1e670f0032a98ac004a396a8ee253da9d07cf0b974ee236bec5349ba555",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x611a4aee28226363b36f074b5b8ed655b268fa464dd8a8fe38567f4770aeb2b4",
            "From": "0xd47a4e56e9262543db39d9203cf1a2e53735f834",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357710,
            "BlockNumber": 18,
            "BlockHash": "0x91ad4c065600cf7bbaa63cf9a62fab74d5f67da246b1873ad6071c0e593f3aaa",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0xc50786ef7108a51059c6dd7724c1c51f5ee198f13171e3fa4e08ec9f945dcc65",
            "From": "0xbfb558f0dceb07fbb09e1c283048b551a4310921",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000130000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357715,
            "BlockNumber": 19,
            "BlockHash": "0x8cac9330e5c3446f2f782219d70704f4ebb2ebe558a53cefe029ee7a5a82e412",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x75001c84f2a7c99733bf91ffb8f80ca89b9d29a9c295230c9995cdc5b75cc3b6",
            "From": "0x258af48e28e4a6846e931ddff8e1cdf8579821e5",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357720,
            "BlockNumber": 20,
            "BlockHash": "0x90789a2c32644a684d5936dd9cfc012b14f6de5f769a8a32016aca0b5ee799fd",
            "Events": [
                {
                    "Number": 20,
                    "Contract": "0x0000000000000000000000000000000000000105",
                    "EventId": "0xa9fb763ca024083f3b32979bc27ec1ff7cb648a232ed0b591cb9d3a7bb1374ad",
                    "Topic1": "",
                    "Topic2": "",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0x75001c84f2a7c99733bf91ffb8f80ca89b9d29a9c295230c9995cdc5b75cc3b6"
                }
            ],
            "TransactionDetails": null
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
    "PageSize":10,
    "Contract": "0x0000000000000000000000000000000000000105"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionsofcontract' \
--data-raw '{
    "PageNo":0,
    "PageSize":10,
    "Contract": "0x0000000000000000000000000000000000000105"
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 0,
    "TotalPage": 28,
    "TotalCount": 273,
    "Transactions": [
        {
            "Hash": "0x9b8b22fea998e8da751ab2971f597487d9bc9ec5cf78edc5af8bbc12a4eee528",
            "From": "0x8c09d936a1b408d6e0afaa537ba4e06c4504a0ae",
            "Cost": 0,
            "Data": "99210986",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357620,
            "BlockNumber": 1,
            "BlockHash": "0x13ad834f28b5860d46b36372de679178a5ae668760b9933303016e3f5ab239ea",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0xa4f49e77af95f02b960947b203258c16578e6b091fff6e3b583babc4e891da74",
            "From": "0xc095448424a5ecd5ca7ccdadfaad127a9d7e88ec",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357630,
            "BlockNumber": 2,
            "BlockHash": "0x85aab5f53d77d8309c5e9b69d2c4c50e69099237a6c63ea631b0b6436d7ad88e",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x08eb78a51561ddc03e6b1b53c13254972152702ff23777cd84eaa29a7a538fbe",
            "From": "0xd47a4e56e9262543db39d9203cf1a2e53735f834",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357635,
            "BlockNumber": 3,
            "BlockHash": "0xac4e749ea72bbeb6004aa81a066eec0b7e7d35b03684423c7f173581ab5ff4ba",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x1a5c37cb47ecf447b2356977a8cc07559490973c777fc754fc18d9c0478a12aa",
            "From": "0xbfb558f0dceb07fbb09e1c283048b551a4310921",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357640,
            "BlockNumber": 4,
            "BlockHash": "0xe1a0f974c7d67006b7fe49f8904141e12e7ae2638b12139b30ee511e41f76047",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0xf087dcc7eacfc8f47106931617ad7e33f9be9bb2b6b4745209aa34d1a64a792e",
            "From": "0x258af48e28e4a6846e931ddff8e1cdf8579821e5",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357645,
            "BlockNumber": 5,
            "BlockHash": "0x84567aec8b3bb3b941784037e547b280ae6f3aafc2a8491aecf0a6778f5581d1",
            "Events": [
                {
                    "Number": 5,
                    "Contract": "0x0000000000000000000000000000000000000105",
                    "EventId": "0xa9fb763ca024083f3b32979bc27ec1ff7cb648a232ed0b591cb9d3a7bb1374ad",
                    "Topic1": "",
                    "Topic2": "",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0xf087dcc7eacfc8f47106931617ad7e33f9be9bb2b6b4745209aa34d1a64a792e"
                }
            ],
            "TransactionDetails": null
        },
        {
            "Hash": "0x445bbcbe9b81790a4c42bf11d9a0d999432327da3858e6a1b46d1d0ecec7a758",
            "From": "0x8c09d936a1b408d6e0afaa537ba4e06c4504a0ae",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357650,
            "BlockNumber": 6,
            "BlockHash": "0x184f504a49876ea6ad4bb1f3f84ac3bc2c8894dbadb6a7b3bc321de50b90fdac",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0xda932e9f7dd06120ea1390f68566d8d1042f60faeebdcfc33d191833296b13ce",
            "From": "0xc095448424a5ecd5ca7ccdadfaad127a9d7e88ec",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000070000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357655,
            "BlockNumber": 7,
            "BlockHash": "0x9d7d014d996f01ec072b05af50909556a8b62ed9dd2d429c93046821a52459ab",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0xc8189341ee38452657bf639bee81e3e78812ed997e33b0b9461728d6a76b70d6",
            "From": "0xd47a4e56e9262543db39d9203cf1a2e53735f834",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357660,
            "BlockNumber": 8,
            "BlockHash": "0x2d63a738ebd07418bd9c04dc6c6bb1d76f15356b38368f01e357314a550cb1de",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x2d5ed3de7bae5622b4e4d523f5029c966c841f9bf1dbf00510882856a71e0526",
            "From": "0xbfb558f0dceb07fbb09e1c283048b551a4310921",
            "Cost": 0,
            "Data": "a348f75b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357665,
            "BlockNumber": 9,
            "BlockHash": "0x47578a4fe9233dc9c7788623617df8237b64eacbe2899f1378f8a787c48eff2f",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0xe4e19bf01187cf22aefda749d2214900bb24dbd4aeb997a529b86c8ea190647f",
            "From": "0x258af48e28e4a6846e931ddff8e1cdf8579821e5",
            "Cost": 0,
            "Data": "a348f75b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000258af48e28e4a6846e931ddff8e1cdf8579821e50000000000000000000000008c09d936a1b408d6e0afaa537ba4e06c4504a0ae000000000000000000000000c095448424a5ecd5ca7ccdadfaad127a9d7e88ec000000000000000000000000d47a4e56e9262543db39d9203cf1a2e53735f834000000000000000000000000bfb558f0dceb07fbb09e1c283048b551a4310921",
            "Gas": 1000000000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000000105",
            "Value": 0,
            "Time": 1606357670,
            "BlockNumber": 10,
            "BlockHash": "0x066b5a44652b1a808a8fa3a2b85b491a6c2e3c1e6b99d68aa958f865ada8e353",
            "Events": [
                {
                    "Number": 10,
                    "Contract": "0x0000000000000000000000000000000000000105",
                    "EventId": "0xa9fb763ca024083f3b32979bc27ec1ff7cb648a232ed0b591cb9d3a7bb1374ad",
                    "Topic1": "",
                    "Topic2": "",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0xe4e19bf01187cf22aefda749d2214900bb24dbd4aeb997a529b86c8ea190647f"
                }
            ],
            "TransactionDetails": null
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
    "PageNo":1,
    "PageSize":10,
    "User": "0xf3a9d42c01635a585f1721463842f8936075105f"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/transactionsofuser' \
--data-raw '{
    "PageNo":1,
    "PageSize":10,
    "User": "0xf3a9d42c01635a585f1721463842f8936075105f"
}'
```

Example Response
```
{
    "PageSize": 10,
    "PageNo": 1,
    "TotalPage": 5,
    "TotalCount": 46,
    "Transactions": [
        {
            "Hash": "0xa47fd4ce4f2019d50e5cd2f4b39d7cd5010543afb2fdfe2586ec4e9b8f27bcfa",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001001",
            "Value": 0,
            "Time": 1606358240,
            "BlockNumber": 124,
            "BlockHash": "0x76a8f203c256071f29c086a31575bd528559d7d1382cd649c91af234ed989afb",
            "Events": [
                {
                    "Number": 124,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000000000000000000000000000000000000000000105",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0xa47fd4ce4f2019d50e5cd2f4b39d7cd5010543afb2fdfe2586ec4e9b8f27bcfa"
                },
                {
                    "Number": 124,
                    "Contract": "0x0000000000000000000000000000000000001001",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "01",
                    "TransactionHash": "0xa47fd4ce4f2019d50e5cd2f4b39d7cd5010543afb2fdfe2586ec4e9b8f27bcfa"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x0000000000000000000000000000000000000105",
                    "Value": "0",
                    "TransactionHash": "0xa47fd4ce4f2019d50e5cd2f4b39d7cd5010543afb2fdfe2586ec4e9b8f27bcfa"
                },
                {
                    "Contract": "0x0000000000000000000000000000000000001001",
                    "From": "0x0000000000000000000000000000000000000000",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "1",
                    "TransactionHash": "0xa47fd4ce4f2019d50e5cd2f4b39d7cd5010543afb2fdfe2586ec4e9b8f27bcfa"
                }
            ]
        },
        {
            "Hash": "0xdfc249517fdd1485c8d3fa7cbdb3553b8560e84ba2084844357db0f5423e3f2b",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001001",
            "Value": 0,
            "Time": 1606358275,
            "BlockNumber": 131,
            "BlockHash": "0xef99f83c8f42123257fc5f9b12d3bae0509b6e3bb7267a8089d525aec77a9287",
            "Events": [
                {
                    "Number": 131,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000000000000000000000000000000000000000000105",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0xdfc249517fdd1485c8d3fa7cbdb3553b8560e84ba2084844357db0f5423e3f2b"
                },
                {
                    "Number": 131,
                    "Contract": "0x0000000000000000000000000000000000001001",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "02",
                    "TransactionHash": "0xdfc249517fdd1485c8d3fa7cbdb3553b8560e84ba2084844357db0f5423e3f2b"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x0000000000000000000000000000000000000105",
                    "Value": "0",
                    "TransactionHash": "0xdfc249517fdd1485c8d3fa7cbdb3553b8560e84ba2084844357db0f5423e3f2b"
                },
                {
                    "Contract": "0x0000000000000000000000000000000000001001",
                    "From": "0x0000000000000000000000000000000000000000",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "2",
                    "TransactionHash": "0xdfc249517fdd1485c8d3fa7cbdb3553b8560e84ba2084844357db0f5423e3f2b"
                }
            ]
        },
        {
            "Hash": "0x583ecf5c2414e9d8b7d35b1414527279b04731ecc2d51a575615806c15925c7f",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001001",
            "Value": 0,
            "Time": 1606358305,
            "BlockNumber": 137,
            "BlockHash": "0x125edde633344f35e1460ea7ab728350bf4d8c7af04c0ff26320940d5a9d1c2d",
            "Events": [
                {
                    "Number": 137,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000000000000000000000000000000000000000000105",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0x583ecf5c2414e9d8b7d35b1414527279b04731ecc2d51a575615806c15925c7f"
                },
                {
                    "Number": 137,
                    "Contract": "0x0000000000000000000000000000000000001001",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "03",
                    "TransactionHash": "0x583ecf5c2414e9d8b7d35b1414527279b04731ecc2d51a575615806c15925c7f"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x0000000000000000000000000000000000000105",
                    "Value": "0",
                    "TransactionHash": "0x583ecf5c2414e9d8b7d35b1414527279b04731ecc2d51a575615806c15925c7f"
                },
                {
                    "Contract": "0x0000000000000000000000000000000000001001",
                    "From": "0x0000000000000000000000000000000000000000",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "3",
                    "TransactionHash": "0x583ecf5c2414e9d8b7d35b1414527279b04731ecc2d51a575615806c15925c7f"
                }
            ]
        },
        {
            "Hash": "0x3956c59501cd5b4ac902e0a1940892a5d5f2484fcff65f048ba3a95fc07a41eb",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001002",
            "Value": 0,
            "Time": 1606358410,
            "BlockNumber": 158,
            "BlockHash": "0xe8fc36028921861c1b680eb3a64992f433ceae6756ca3e3dda2974364cab6875",
            "Events": [
                {
                    "Number": 158,
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x000000000000000000000000f3a9d42c01635a585f1721463842f8936075105f",
                    "Topic2": "0x0000000000000000000000000000000000000000000000000000000000000105",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "",
                    "TransactionHash": "0x3956c59501cd5b4ac902e0a1940892a5d5f2484fcff65f048ba3a95fc07a41eb"
                },
                {
                    "Number": 158,
                    "Contract": "0x0000000000000000000000000000000000001002",
                    "EventId": "0xbeabacc8ffedac16e9a60acdb2ca743d80c2ebb44977a93fa8e483c74d2b35a8",
                    "Topic1": "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "Topic2": "0x0000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Topic3": "",
                    "Topic4": "",
                    "Data": "01",
                    "TransactionHash": "0x3956c59501cd5b4ac902e0a1940892a5d5f2484fcff65f048ba3a95fc07a41eb"
                }
            ],
            "TransactionDetails": [
                {
                    "Contract": "0x0000000000000000000000000000000000000103",
                    "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
                    "To": "0x0000000000000000000000000000000000000105",
                    "Value": "0",
                    "TransactionHash": "0x3956c59501cd5b4ac902e0a1940892a5d5f2484fcff65f048ba3a95fc07a41eb"
                },
                {
                    "Contract": "0x0000000000000000000000000000000000001002",
                    "From": "0x0000000000000000000000000000000000000000",
                    "To": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
                    "Value": "1",
                    "TransactionHash": "0x3956c59501cd5b4ac902e0a1940892a5d5f2484fcff65f048ba3a95fc07a41eb"
                }
            ]
        },
        {
            "Hash": "0x689de850b37031ee6a5351e61bf5a3320f5bd5cd3eafe142e47d6f4c70c695b7",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001002",
            "Value": 0,
            "Time": 1606358410,
            "BlockNumber": 158,
            "BlockHash": "0xe8fc36028921861c1b680eb3a64992f433ceae6756ca3e3dda2974364cab6875",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x6ea38207ccb7e3ee000245517e26176843020ebc6b1ada00c366f0681425e5c3",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001002",
            "Value": 0,
            "Time": 1606358410,
            "BlockNumber": 158,
            "BlockHash": "0xe8fc36028921861c1b680eb3a64992f433ceae6756ca3e3dda2974364cab6875",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x775c37ef69003b85e1bd2494911286b588119593bba7149fa3e33fa892b72ce3",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001002",
            "Value": 0,
            "Time": 1606358410,
            "BlockNumber": 158,
            "BlockHash": "0xe8fc36028921861c1b680eb3a64992f433ceae6756ca3e3dda2974364cab6875",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x779ec5317d217f29bcb3af2fd9b1f7f6a591b6731141c1c0dc83882f5beb9cc9",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001002",
            "Value": 0,
            "Time": 1606358410,
            "BlockNumber": 158,
            "BlockHash": "0xe8fc36028921861c1b680eb3a64992f433ceae6756ca3e3dda2974364cab6875",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x96bfc952aa85631a873a18be34bbcdecac887cb0c7dcc2e4439b5b73147caf1b",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001002",
            "Value": 0,
            "Time": 1606358410,
            "BlockNumber": 158,
            "BlockHash": "0xe8fc36028921861c1b680eb3a64992f433ceae6756ca3e3dda2974364cab6875",
            "Events": null,
            "TransactionDetails": null
        },
        {
            "Hash": "0x9e609c3e1b5e2eb4f2c254c643079436febb9d39b6fcc3472e575910ed52fff7",
            "From": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Cost": 0,
            "Data": "85a66b910000000000000000000000002cd9d589d46122e4eddc495b49feda0b526c1af70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000057878787878000000000000000000000000000000000000000000000000000000",
            "Gas": 100000,
            "GasPrice": 0,
            "To": "0x0000000000000000000000000000000000001002",
            "Value": 0,
            "Time": 1606358410,
            "BlockNumber": 158,
            "BlockHash": "0xe8fc36028921861c1b680eb3a64992f433ceae6756ca3e3dda2974364cab6875",
            "Events": null,
            "TransactionDetails": null
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
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/pltholders' \
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
    "TotalCount": 6,
    "PLTHolderInfos": [
        {
            "Address": "0xf3a9d42c01635a585f1721463842f8936075105f",
            "Amount": 450000000000000000
        },
        {
            "Address": "0x0000000000000000000000000000000000000105",
            "Amount": 149992320000000000
        },
        {
            "Address": "0x3eacfd9c0adb0dd994193e4927ddb3e273d3b4c6",
            "Amount": 2048000000000
        },
        {
            "Address": "0x8fc2680e989c39eca4544c5961fb2415d0e8d907",
            "Amount": 2048000000000
        },
        {
            "Address": "0xf716a8e5c1324ae18181f7200b2da26038769782",
            "Amount": 2048000000000
        },
        {
            "Address": "0xa2ec66f9dee661e096db3c4de1187d32e48cc959",
            "Amount": 1536000000000
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
    "Address": "0xf3a9d42c01635a585f1721463842f8936075105f"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/pltholderinfo' \
--data-raw '{
    "Address": "0xf3a9d42c01635a585f1721463842f8936075105f"
}'
```

Example Response
```
{
    "Address": "0xf3a9d42c01635a585f1721463842f8936075105f",
    "Amount": 450000000000000000
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
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nfts' \
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
    "TotalCount": 4,
    "NFTs": [
        {
            "NFT": "0x0000000000000000000000000000000000001001"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001002"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001003"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001004"
        }
    ]
}
```

### POST nfttokens

Request 
```
http://localhost:8080/v1/nfttokens
```

BODY raw
```
{
    "NFT": "0x0000000000000000000000000000000000001001",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nfttokens' \
--data-raw '{
    "NFT": "0x0000000000000000000000000000000000001001",
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
    "NFTTokens": [
        {
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "1"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "2"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "3"
        }
    ]
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
    "NFT": "0x0000000000000000000000000000000000001001",
    "Token": "1"
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nfttokeninfo' \
--data-raw '{
    "NFT": "0x0000000000000000000000000000000000001001",
    "Token": "1"
}'
```

Example Response
```
{
    "NFT": "0x0000000000000000000000000000000000001001",
    "Token": "1",
    "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
    "URI": "xxxxx"
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
    "NFT": "0x0000000000000000000000000000000000001001",
    "Address": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nftholdersofuser' \
--data-raw '{
    "NFT": "0x0000000000000000000000000000000000001001",
    "Address": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
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
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "1",
            "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "URI": "xxxxx"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "2",
            "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "URI": "xxxxx"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "3",
            "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "URI": "xxxxx"
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
    "NFT": "0x0000000000000000000000000000000000001001",
    "PageNo":0,
    "PageSize":10
}
```

Example Request
```
curl --location --request POST 'http://localhost:8080/v1/nftholders' \
--data-raw '{
    "NFT": "0x0000000000000000000000000000000000001001",
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
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "1",
            "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "URI": "xxxxx"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "2",
            "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "URI": "xxxxx"
        },
        {
            "NFT": "0x0000000000000000000000000000000000001001",
            "Token": "3",
            "Owner": "0x2cd9d589d46122e4eddc495b49feda0b526c1af7",
            "URI": "xxxxx"
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
            "StakeAmount": 50000000000000000
        },
        {
            "Address": "0x6A708455C8777630AaC9d1e7702d13F7a865b27C",
            "DelegateFactor": 2000,
            "StakeAmount": 50000000000000000
        },
        {
            "Address": "0xAd3Bf5eD640CC72f37BD21d64A65C3C756e9C88C",
            "DelegateFactor": 2000,
            "StakeAmount": 50000000000000000
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
    "StakeAmount": 50000000000000000
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