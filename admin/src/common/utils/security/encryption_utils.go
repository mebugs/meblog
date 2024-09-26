package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

/**
 *
 * AES加解密工具类
 * 使用AES-128-CBC加密模式，key需要为16位。
 * key 由方法传入(生成估计16位的随机数）
 *
 * @author 米虫丨www.mebugs.com
 * @since 2022-08-24
 */
var iv = []byte("AES_SiteOL_Stone")

// AESEncrypt 加密工具
func AESEncrypt(origData, key string) (string, error) {
	origBytes := []byte(origData)
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origBytes = PKCS5Padding(origBytes, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cry := make([]byte, len(origBytes))

	blockMode.CryptBlocks(cry, origBytes)
	return base64.StdEncoding.EncodeToString(cry), nil
}

// AESDecrypt 解密工具
func AESDecrypt(cryData, key string) (string, error) {
	keyBytes := []byte(key)
	decodeData, err := base64.StdEncoding.DecodeString(cryData)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origData, []byte(decodeData))
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return string(origData), nil
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func ZeroPadding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(cipherText, padText...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
