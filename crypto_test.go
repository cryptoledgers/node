package main

import (
	"testing"

	"github.com/btcd/chaincfg/chainhash"
	cryptoutil "github.com/polygonledger/node/crypto"
)

func TestBasicSign(t *testing.T) {

	//sign newly created keypair should be valid signature
	keypair := cryptoutil.PairFromSecret("test")
	message := "test"

	signature := cryptoutil.SignMsgHash(keypair, message)
	verified := cryptoutil.VerifyMessageSign(signature, keypair, message)
	if !verified {
		t.Error("msg failed")
	}

	messagefalse := "testshouldbefalse"
	verifiedfalse := cryptoutil.VerifyMessageSign(signature, keypair, messagefalse)

	if verifiedfalse {
		t.Error("sign verify should fail")
	}

	otherkeypair := cryptoutil.PairFromSecret("testother")
	verifiedother := cryptoutil.VerifyMessageSign(signature, otherkeypair, message)
	if verifiedother {
		t.Error("sign other should fail")
	}

}

func TestRecordPrivkey(t *testing.T) {
	// Decode the hex-encoded pubkey of the recipient.
	// pubKeyBytes, err := hex.DecodeString("04115c42e757b2efb7671c578530ec191a1" +
	// 	"359381e6a71127a9d37c486fd30dae57e76dc58f693bd7e7010358ce6b165e483a29" +
	// 	"21010db67ac11b1b51b651953d2") // uncompressed pubkey

	// pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())

	// // Encrypt a message decryptable by the private key corresponding to pubKey
	// // message := "test message"
	// // ciphertext, err := btcec.Encrypt(pubKey, []byte(message))

	// // Decode the hex-encoded private key
	// pkBytes, err := hex.DecodeString("a11b0a4e1a132305652ee7a8eb7848f6ad" +
	// 	"5ea381e3ce20a2c086a2e388230811")

	// note that we already have corresponding pubKey
	//privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	// Try decrypting and verify if it's the same message
	// plaintext, err := btcec.Decrypt(privKey, ciphertext)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(plaintext))
}

func TestDecode(t *testing.T) {

	pubKey := cryptoutil.PubKeyFromHex("02a673638cb9587cb68ea08dbef685c6f2d2a751a8b3c6f2a7e9a4999e6e4bfaf5")

	h := "30450220090ebfb3690a0ff115bb1b38b8b323a667b7653454f1bccb06d4bbdca42c2079022100ec95778b51e7071cb1205f8bde9af6592fc978b0452dafe599481c46d6b2e479"
	signature := cryptoutil.SignatureFromHex(h)

	message := "test message"
	messageHash := chainhash.DoubleHashB([]byte(message))
	verified := signature.Verify(messageHash, &pubKey)

	if !verified {
		t.Error("signature decoding failed")
	}
}

func TestAddress(t *testing.T) {

	keypair := cryptoutil.PairFromSecret("test")
	pubkey_string := cryptoutil.PubKeyToHex(keypair.PubKey)
	if pubkey_string != "03dab2d148f103cd4761df382d993942808c1866a166f27cafba3289e228384a31" {
		t.Error("expected different hex of pubkey")
	}

	hexString := "a11b0a4e1a132305652ee7a8eb7848f6ad5ea381e3ce20a2c086a2e388230811"
	privKey := cryptoutil.PrivKeyFromHex(hexString)
	privKeyHex := cryptoutil.PrivKeyToHex(privKey)

	if privKeyHex != hexString {
		t.Error("privkey encoding")
	}

	// pubKeyBytes, err := hex.DecodeString(pubkey_string)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//FAIL
	// if *pubKey != keypair.PubKey {
	// 	log.Println(*pubKey, keypair.PubKey)
	// 	t.Error("error recoding pubkey")
	// }

	addr := cryptoutil.Address(pubkey_string)
	if addr[0] != 'P' {
		t.Error("address should start with P ", addr[0])
	}

	if len(addr) != 13 {
		t.Error("length of address should be 13 ", len(addr))
	}
}

func TestSignHardcoded(t *testing.T) {
	pub := "03dab2d148f103cd4761df382d993942808c1866a166f27cafba3289e228384a31"
	pubkey := cryptoutil.PubKeyFromHex(pub)

	keypair := cryptoutil.PairFromSecret("test")
	h := cryptoutil.PubKeyToHex(keypair.PubKey)

	if h != pub {
		t.Error("hardcoded pubkey wrong")
	}

	sig := "3045022100b6789781f032512fc9ae06e9621ca4ce40d317a83a6b6f4ee1ad35942a3c928602204d8f03b330efc416b822447862333140d0acb3323d4575f1efba6e5418a036f7"
	sign := cryptoutil.SignatureFromHex(sig)
	msg := "test"
	verified := cryptoutil.VerifyMessageSignPub(sign, pubkey, msg)
	if !verified {
		t.Error("should verify standard ", verified)
	}

}

func TestGenkeys(t *testing.T) {
	// Decode a hex-encoded private key.
	h := "22a47fa09a223f2aa079edf85a7c2d4f87" +
		"20ee63e502ee2869afab7de234b80c"

	keypair := cryptoutil.PairFromHex(h)

	if cryptoutil.PubKeyToHex(keypair.PubKey) == "" {
		t.Error("keypair is nil")
	}

	//log.Println("pubkey example %v", keypair.PubKey)
	//log.Println(keypair.PrivKey)

	//hash := sha256.Sum256(pubKey.Serialize())

	// Sign a message using the private key.
	// message := "test message"
	// messageHash := chainhash.DoubleHashB([]byte(message))
	// signature, err := privKey.Sign(messageHash)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Serialize and display the signature.
	// fmt.Printf("Serialized Signature: %x\n", signature.Serialize())
	// // Verify the signature for the message using the public key.
	// verified := signature.Verify(messageHash, pubKey)
	// fmt.Printf("Signature Verified? %v\n", verified)

	// data := []byte("hello")
	// hash := sha256.Sum256(data)
	// fmt.Printf("%x", hash[:])

	// timestamp := time.Now().Unix()
	// b := []byte(string(timestamp))
	// hash = sha256.Sum256(b)

	// fmt.Printf("\n%x", hash[:])

}
