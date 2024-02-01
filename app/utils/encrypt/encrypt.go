package encrypt

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/tjfoc/gmsm/sm3"
)

func GetSignBySM3(json string) (sign string) {
	arrBytePlain := []byte(json)
	hash := sm3.New()
	_, errWrite := hash.Write(arrBytePlain)
	err.ErrCheck(errWrite)
	arrByteCipher := hash.Sum(nil)
	sign = fmt.Sprintf(
		"%x",
		arrByteCipher)
	return sign
}
