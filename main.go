/*	Code to Retrive data from CSV to private blocks in blockchain
*
*
 */
package main

import (
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)
var block_number = 1
//test def to assign transactions
type Trans struct{
	adhaar_num 	[]string
	full_name 	[]string
	ror_data 	[]string
	bank_data 	[]string
	income_tax 	[]string	
}

type Block struct {
	timestamp    time.Time
	transactions1 	[]string
	transactions2 	[]string
	transactions3 	[]string
	transactions4 	[]string
	transactions5 	[]string
	prevhash     	[]byte
	Hash         	[]byte
}

func main() {
	records := csv_call()
	// fmt.Println(records)
	//block 1
	abc1 := []string{records[0][0]}
	abc2 := []string{records[0][1]}
	abc3 := []string{records[0][2]}
	abc4 := []string{records[0][3]}
	abc5 := []string{records[0][4]}
	xyz := Blocks(abc1, abc2, abc3, abc4, abc5, []byte{})
	fmt.Println("This is our First Block")
	Print(xyz)

	//block 2
	// pqrs1 := []string{records[1][0]}
	// pqrs2 := []string{records[1][1]}
	// pqrs3 := []string{records[1][2]}
	// pqrs4 := []string{records[1][3]}
	// pqrs5 := []string{records[1][4]}
	// klmn := Blocks(pqrs1,pqrs2, pqrs3, pqrs4, pqrs5, xyz.Hash)
	// fmt.Println("This is our Second Block")
	// Print(klmn)
	UID := dataset_block(records, xyz)
	dataset_block(records, UID)

}

//user defined function to keep adding the new blocks and hashids
func dataset_block(records [][]string, xyz *Block) *Block {
	for i := block_number;i<5;block_number++ {
		adhaar_num := []string{records[i][0]}
		full_name := []string{records[i][1]}
		ror_data := []string{records[i][2]}
		bank_data := []string{records[i][3]}
		income_tax := []string{records[i][4]}
		UID := Blocks(adhaar_num,full_name, ror_data, bank_data, income_tax, xyz.Hash)
		fmt.Println("This is our Block :",block_number)
		Print(UID)
		return UID
	}
	return nil
}

func Blocks(transactions1 []string, transactions2 []string, transactions3 []string, transactions4 []string, transactions5 []string, prevhash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:    currentTime,
		transactions1: transactions1,
		transactions2: transactions2,
		transactions3: transactions3,
		transactions4: transactions4,
		transactions5: transactions5,
		prevhash:     prevhash,
		Hash:         NewHash(currentTime, transactions1, transactions2, transactions3, transactions4, transactions5, prevhash),
	}
}

func NewHash(time time.Time, transactions1 []string, transactions2 []string, transactions3 []string, transactions4 []string, transactions5 []string, prevhash []byte) []byte {
	input := append(prevhash, time.String()...)
	for transactions1 := range transactions1 {
		input = append(input, string(rune(transactions1))...)
	}
	for transactions2 := range transactions2 {
		input = append(input, string(rune(transactions2))...)
	}
	for transactions3 := range transactions3 {
		input = append(input, string(rune(transactions3))...)
	}
	for transactions4 := range transactions4 {
		input = append(input, string(rune(transactions4))...)
	}
	for transactions5 := range transactions5 {
		input = append(input, string(rune(transactions5))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func Print(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevhash: %x\n", block.prevhash)
	fmt.Printf("\thash: %x\n", block.Hash)
	Transaction(block)
}

func Transaction(block *Block) {
	fmt.Println("\tTransactions1:")
	for i, transaction1 := range block.transactions1 {
		fmt.Printf("\t\tAdhaar Number: %q\n", transaction1)
		i++
	}
	//fmt.Println("\tTransactions2:")
	for i, transaction2 := range block.transactions2 {
		fmt.Printf("\t\tFull Name: %q\n", transaction2)
		i++
	}
	//fmt.Println("\tTransactions3:")
	for i, transaction3 := range block.transactions3 {
		fmt.Printf("\t\tROR: %q\n", transaction3)
		i++
	}
	//fmt.Println("\tTransactions2:")
	for i, transaction4 := range block.transactions4 {
		fmt.Printf("\t\tBank Details: %q\n", transaction4)
		i++
	}
	//fmt.Println("\tTransactions2:")
	for i, transaction5 := range block.transactions5 {
		fmt.Printf("\t\tIncome Tax: %q\n", transaction5)
		i++
	}
}

// csv or excel files to import and read and store in variable
func csv_call() [][]string {
	// open CSV file
	fd, error := os.Open("data.csv")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Successfully opened the CSV file")
	defer fd.Close()

	// read CSV file
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}
	// fmt.Println(records[0])
	// dataset_extract(records)
	return records

}

// extracts the data from the csv file individually
func dataset_extract(records [][]string) {
	for i := 1; i <= 3; i++ {
		adhaar_num := records[i][0]
		full_name := records[i][1]
		ror_data := records[i][2]
		bank_data := records[i][3]
		income_tax := records[i][4]
		fmt.Println(adhaar_num, full_name, ror_data, bank_data, income_tax)
	}
}

