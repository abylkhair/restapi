package cases

type CoinService struct {
	storage CoinStorage
	client  CryptoClient
	//logger
}

//GetACtualRates->Storage(lastvalueof coin string)
//GetActualRates->Storage(string coin)->CryptoClient(coin (if coin == ""))
//StoreActualRates->CryptoClient()->Storage(save)

// two adapters needed
//work with storage
//work with cryptoapi
