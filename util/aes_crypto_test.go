package util

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesDecryptMessage(t *testing.T) {
	aesKey, err := base64.StdEncoding.DecodeString("HogNecGqZeDxFIDGjBwWKw==")
	aesIv, err := base64.StdEncoding.DecodeString("aoZqkfGDWwqj6rFgWdafyw==")

	base64Test := "4B9B1aknFM6yQjAh9mFxH3iN4PZXUGfpBLB98CRAVZzNfUI4J1WUur70+NSH/5MXcCidrt44hi6dkByTtRPxrTIV1BOOTvaa2G5NWunTXZJ37/Oq0ezydbDal5v+X3bvVVeFR6MkhYI+hT9xVl2XnE/Bzon2gIq9F9Fy8Yny0VPqsQ95xUyXnN3/IuhiquR1pAgKjDK3kgCoqhUVNa0dRRQQgTNIpy1djbLfyErPXGTXe1qhAj7RvDdJtRloEfg63JgaB/QTR2BLEGT7/GfHnwROfngxa3esGDeBr9Mtav67R4PjYESVFtH2Yf2npaDWAvAMvr+8hNVd2tjy/3HgImctZWo7bh1OHa/ktH4wKYVbTgxZPg/lgTWC+zsl1Z9g07vWQyGpaA11HODDGn1Kdvh6esiY7T3JQ15nKs1rdyXigNQFb9+kicPiInKVfOiuMcGEazli5/UQHF5rnuWhkg/wy+PRbZybR18lLh5d+EW1CnK8gNcQblDJPvx6QFcFhPxr28WkEd7ys0BZQGCIkQ=="
	ciphertext, err := base64.StdEncoding.DecodeString(base64Test)
	if err != nil {
		t.Error(err)
		return
	}

	raw, err := AesDecryptMessage(ciphertext, aesKey, aesIv)
	fmt.Println(string(raw), err)
	if err != nil {
		t.Error(err)
		return
	}
}
