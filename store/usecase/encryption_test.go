package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {

	file := []byte("123456791234567912345679123456791234567912345679123456791234567912345679")
	// expect := []byte{122, 20, 181, 106, 196, 60, 166, 99, 101, 194, 186, 81, 140, 253, 96, 24, 28, 165, 72, 169, 97, 242, 1, 97, 190, 184, 129, 4, 190, 209, 98, 216, 148, 246, 159, 238, 107, 114, 101, 78, 0, 213, 167, 133, 120, 243, 240, 234, 237, 182, 207, 52, 161, 144, 129, 192, 118, 31, 151, 190, 38, 2, 234, 7, 233, 212, 181, 148, 124, 43, 128, 226, 202, 28, 225, 139, 117, 76, 63, 30, 240, 68, 144, 47, 142, 124, 211, 4}

	_, err := NewEncryption().Encrypt(context.Background(), &file)

	if err != nil {
		t.Fatal(err)
	}

}

func TestDecrypt(t *testing.T) {

	file := []byte{230, 30, 113, 100, 42, 233, 121, 94, 112, 128, 171, 87, 232, 35, 26, 50, 87, 255, 213, 168, 140, 151, 183, 102, 87, 102, 243, 177, 207, 98, 48, 108, 3, 173, 55, 193, 114, 106, 53, 126, 232, 221, 152, 67, 134, 232, 44, 246, 166, 58, 163, 20, 61, 233, 130, 181, 77, 12, 204, 241, 38, 24, 3, 82, 242, 186, 210, 22, 153, 31, 230, 4, 188, 224, 173, 5, 195, 123, 228, 214, 218, 143, 132, 113, 7, 22, 203, 112}
	expect := []byte("123456791234567912345679123456791234567912345679123456791234567912345679")
	result, err := NewEncryption().Decrypt(context.Background(), &file)

	fmt.Println(string(*result))

	if err != nil {
		t.Fatal(err)
	}
	equal := reflect.DeepEqual(*result, expect)

	if !equal {
		t.Fatal("not equal")
	}

}