package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"gossh/test/helpers"
	"time"
)

//参考文档
//http://www.topgoer.com/%E5%85%B6%E4%BB%96/%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86/%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86.html
//高级加密标准（Adevanced Encryption Standard ,AES）

//16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
//key不能泄露
//var PwdKey = []byte("DIS**#KKKDJJSKDI")
var PwdKey = "FJLASFDJL;ADJFLD"

//PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充的反向操作，删除填充字符串
func PKCS7UnPadding1(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

//实现加密
func AesEcrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//实现解密
func AesDeCrypt(cypted []byte, key []byte) (string, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = PKCS7UnPadding1(origData)
	if err != nil {
		return "", err
	}
	return string(origData), err
}

//加密base64
func EnPwdCode(pwdStr string) string {
	pwd := []byte(pwdStr)
	result, err := AesEcrypt(pwd, []byte(PwdKey))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(result)
}

//解密
func DePwdCode(pwd string) string {
	temp, _ := hex.DecodeString(pwd)
	//执行AES解密
	res, _ := AesDeCrypt(temp, []byte(PwdKey))
	return res
}

func testAes() {
	//加密
	str, _ := helpers.EncryptByAes([]byte("test"))
	//解密
	str1, _ := helpers.DecryptByAes(str)
	//打印
	fmt.Printf(" 加密：%v\n 解密：%s\n ",
		str, str1,
	)
}

//测试速度
func testAesTime() {
	startTime := time.Now()
	count := 1000000
	for i := 0; i < count; i++ {
		str, _ := helpers.EncryptByAes([]byte("test"))
		helpers.DecryptByAes(str)
	}
	fmt.Printf("%v次 - %v", count, time.Since(startTime))
}

func main() {
	//aes加密
	fmt.Println("test")

	testAes()

	destring := `{"name":"test","site":"https://www.test.com"}`
	deStr := EnPwdCode(destring)
	fmt.Println(deStr) // f9bb31ab1ff35c51d9698cff7636490ae2e7c98abd58d2477c67ce53c463fe4d82ef1f058341973b81b68ad72070636b

	//aes解密
	decodeStr := DePwdCode(deStr)
	fmt.Println(decodeStr) //{"name":"test","site":"https://www.test.com"}
}
