package wallet

import (
	"encoding/base64"
	"github.com/everFinance/goar/types"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testWallet *Wallet
var err error

func init() {
	clientUrl := "https://arweave.net"
	testWallet, err = New([]byte(`{ "kty": "RSA",
	"n":
	 "nQ9iy1fRM2xrgggjHhN1xZUnOkm9B4KFsJzH70v7uLMVyDqfyIJEVXeJ4Jhk_8KpjzYQ1kYfnCMjeXnhTUfY3PbeqY4PsK5nTje0uoOe1XGogeGAyKr6mVtKPhBku-aq1gz7LLRHndO2tvLRbLwX1931vNk94bSfJPYgMfU7OXxFXbTdKU38W6u9ShoaJGgUQI1GObd_sid1UVniCmu7P-99XPkixqyacsrkHzBajGz1S7jGmpQR669KWE9Z0unvH0KSHxAKoDD7Q7QZO7_4ujTBaIFwy_SJUxzVV8G33xvs7edmRdiqMdVK5W0LED9gbS4dv_aee9IxUJQqulSqZphPgShIiGNl9TcL5iUi9gc9cXR7ISyavos6VGiem_A-S-5f-_OKxoeZzvgAQda8sD6jtBTTuM5eLvgAbosbaSi7zFYCN7zeFdB72OfvCh72ZWSpBMH3dkdxsKCDmXUXvPdDLEnnRS87-MP5RV9Z6foq_YSEN5MFTMDdo4CpFGYl6mWTP6wUP8oM3Mpz3-_HotwSZEjASvWtiff2tc1fDHulVMYIutd52Fis_FKj6K1fzpiDYVA1W3cV4P28Q1-uF3CZ8nJEa5FXchB9lFrXB4HvsJVG6LPSt-y2R9parGi1_kEc6vOYIesKspgZ0hLyIKtqpTQFiPgKRlyUc-WEn5E",
	"e": "AQAB",
	"d":
	 "LIqif_yFrcm_q37XRr5KFiC4oUUsQKb5dx7fbLPlzXmsYb6OdfTLoFloVrOhYQ85uw2gNMRqToOAmgDAroQDspaoivlo5bhwP7R4orSVJP84xKzJMx-aNke3hGZtywQdytqfmQv_i3jxRm0Si33EXUnrWQVbEVmCEJ9kfgaIJ0NhALQ8TGx7dxv7cLp6U3zY0X2_PrsVkday5MFS45Wt4vHuYaGeBS4KFygHDflOlKiJ4FGksU3wzyBFO0o1tST21ayxd_G6sbdyar72sQU-asBvYU3kSVMuZs20i1C67qEizk1jqcdKbRuKRApqqs7ub8g2U6yDQaZYqft7KqC8Oi6X-VrUFNMttzvrTmJ0nhzGs1g7SJnZp1NJ1yvqwHOoisUjd0kyqBsDBzT_q0UsYOTM9yi3Ve021R7ghVQPGuvbDmRz8rCm_XusWRp2pzE1R14vot6G67w_2eXGNeD8GYjLKEX35DJwQqhqmD4AVR03u1slY7QK_EI2TxD2aCSpsBToX2Rix0E18dfPCkPmb4WAXccgyMKyR07w_31eypp1qdcQ4cNB_3-2QvM6kbzH3LJkRyPZzamiTB4pKoaQRDqKas0_QctmIdjAWNDr79q4cDcWB0de6X8H6QTjJoL28Lf6pvXXatLoHMAaVHWZpq-yz04-6oiA4pY6DFIs8AE",
	"p":
	 "yL6EzJ5NP0T4vUyWDjkGRI2JAqD9wJhaNd3jt-2Wip3aOQBcmWyagM2tRfbQunwWRquVGd6oSW-Tfr5aBesY3C_EKkbiXt7CSJ9iWz15FBWlLe5lSLbVi5tIcY7iV_D3iFND4mX39gmILX7p_sPmt6z5ZpHKv9MvypjuU2jKScC0HER5fgB-xJAhOQ45xfxuqjpDY6vVWUms6ZSTF3NOrLIrcupNdxUpWNjxKI6A48cvfKdAwJi-vtJcEZFrCJCv98UTgZFSgiWyUVcW2Bt9BYJQRKffm7cuC4aT_xwYq4sMcALoEq0CXIQ9XK4rRPH2bq6E2aZ2ll57OUBWR_1KEQ",
	"q":
	 "yEqoTJqr4EyWoxa-qz5x4FEylw9ueTqLmnZGBbRZaVwJnEY-ydoOaty9ZOtGA6g8EfLVGhtYy_B51479gpDTtygV846313WPDU3drTsPvGNmuA0dD5w4Vva-759Zmc4BcoJA7UYnlr4QRaLQuNAmf9RCgA-yTtHppt5Lre7tGG5K4j7uHRPzp3NhKPq8WosbydN0kFz-a2Vn9kZjoqI3m1JqvX9wRGFzOZlUnQTWA7MaGv4SgoqMMu8PyKQLjhbIJyvnIGI7NZjCi1uw0H03skdgC8bBEL_uSp6Go8nr1apGJ6o1x5_hriofgD9DvpKMkEsx4uTxRvxFys8w3jh9gQ",
	"dp":
	 "Il7kc_hit4OCpz62roa6-P_Wxplz-Qbc4z4zoClQzjkKxRm3wRkkNwuAMGt6_4MBeWYlaEGERNaSxW-oED1Zi1GuX6K1XZL8ZtzLRV34HiU6m-umcdXEKFwVAkR5op8Cctf21ouo8fpd05RYUiOOnEJEjXhG46MwGpsmqydVA124OOLMfnNtQRCAb7ls0OZQuFqzcRxZsij4LyIeMTSv8seqwsk1LD92Td0PJWeIz_cpvUkRwCgm-Jsh4mwojFXhmyWmGlgcbWYw6tZjder28_uE7Mxlb87kVlrbeiGAY9ax8Xe97nyq29ZUf0re47YeAINnAbELuuFAbeQDId5PUQ",
	"dq":
	 "woDVvUZ64OAfbRNaZ_vFJHxVr6K5uppjFcYDq-h-57UMVClXMjhCxf3FIqrjnAuVAi0aSzcBXVMTT4S5pUC1iOkxoAsZdu_f0qCqRF7VojG5f8SkUxN3FuSZeSP7JESM3UGmgYUeTuIV9TnujXr92CctyST1GFv7FiRLxAYBUzdQGzPXkn9cn2GJmf0cSqVKgA2L5eGY5HxeoCes_DOh4oD_zTRjttQXzHidVbprhr43_Lx9By46hf_oCQVdf0eaaYfV9HnQW_UT_7c0FtNy8fskR2tk87ofU3Fs-MPO9PhdFonRniEiTTr0ylslk3zHahzLvjZsJG457ICWSUb8gQ",
	"qi":
	 "ENtqTq3iiDkeiyPWD7pNRfiwIJnY5Zf97yXakxe04usHXWKmZulllttqsDkfHOXkBxRxHxqqTgOLuRpNsLrpI5MAxs8uSl13A70LUzHldnE8ePgt0688UpoI5Iw9oV2RdF_LvSrsgpa-SeexXxbZqXWpDNeUxYt2S327cS8HmrnETKy9z9VoVFmCT6_NCnxOaOTwr67dPBnGnW7nT3499m_aqmikCNjcmkfYihED6S2jZBRHPaSDM7JPPyQSEyRkGjR4z9JzhLOvbJf8tDKSE00JXJClmbpX-5qRcNt0gcJy6ceYQs-c94I24yGpunMMSwGo2i1-sGNwH1wj5-gv1Q" }`),
		clientUrl)

	if err != nil {
		panic(err)
	}
}

func TestPubKey(t *testing.T) {
	pubKey := testWallet.PubKey
	assert.Equal(t, "nQ9iy1fRM2xrgggjHhN1xZUnOkm9B4KFsJzH70v7uLMVyDqfyIJEVXeJ4Jhk_8KpjzYQ1kYfnCMjeXnhTUfY3PbeqY4PsK5nTje0uoOe1XGogeGAyKr6mVtKPhBku-aq1gz7LLRHndO2tvLRbLwX1931vNk94bSfJPYgMfU7OXxFXbTdKU38W6u9ShoaJGgUQI1GObd_sid1UVniCmu7P-99XPkixqyacsrkHzBajGz1S7jGmpQR669KWE9Z0unvH0KSHxAKoDD7Q7QZO7_4ujTBaIFwy_SJUxzVV8G33xvs7edmRdiqMdVK5W0LED9gbS4dv_aee9IxUJQqulSqZphPgShIiGNl9TcL5iUi9gc9cXR7ISyavos6VGiem_A-S-5f-_OKxoeZzvgAQda8sD6jtBTTuM5eLvgAbosbaSi7zFYCN7zeFdB72OfvCh72ZWSpBMH3dkdxsKCDmXUXvPdDLEnnRS87-MP5RV9Z6foq_YSEN5MFTMDdo4CpFGYl6mWTP6wUP8oM3Mpz3-_HotwSZEjASvWtiff2tc1fDHulVMYIutd52Fis_FKj6K1fzpiDYVA1W3cV4P28Q1-uF3CZ8nJEa5FXchB9lFrXB4HvsJVG6LPSt-y2R9parGi1_kEc6vOYIesKspgZ0hLyIKtqpTQFiPgKRlyUc-WEn5E", base64.RawURLEncoding.EncodeToString(pubKey.N.Bytes()))
}

func TestAddress(t *testing.T) {
	addr := testWallet.Address
	assert.Equal(t, "eIgnDk4vSKPe0lYB6yhCHDV1dOw3JgYHGocfj7WGrjQ", addr)
}

func TestWallet_SendBigData(t *testing.T) {
	arNode := "https://arweave.net"
	w, err := NewFromPath("../example/testKey.json", arNode)
	assert.NoError(t, err)
	bigData, err := ioutil.ReadFile("../testFile.json")
	assert.NoError(t, err)
	t.Log("test file length: ", len(bigData))

	tags := []types.Tag{
		types.Tag{
			Name:  "Used",
			Value: "TestChunks",
		},
		types.Tag{
			Name:  "CreatedBy",
			Value: "goar",
		},
		types.Tag{
			Name:  "CreateTime",
			Value: time.Now().Local().String(),
		},
	}
	id, err := w.SendBigData(bigData, tags, 0)
	assert.NoError(t, err)
	t.Log("tx hash: ", id)

}
