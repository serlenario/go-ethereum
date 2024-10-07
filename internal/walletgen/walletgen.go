package walletgen

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

func GenerateWallets(fileName string, numberOfWallets int) error {
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("File %s already exists, creation skipped.\n", fileName)
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking file: %v", err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Failed to close the file: %v\n", err)
		}
	}(file)

	for i := 0; i < numberOfWallets; i++ {
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			return fmt.Errorf("failed to generate private key: %v", err)
		}

		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		_, err = file.WriteString(fmt.Sprintf("%s\n", address.Hex()))
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	fmt.Printf("Successfully created %d Ethereum addresses and saved them to %s\n", numberOfWallets, fileName)
	return nil
}
