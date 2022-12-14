package api

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"lyra-bot-web-crawlers/pkg/entity"
	"lyra-bot-web-crawlers/pkg/setting"
	ponyImageConstant "lyra-bot-web-crawlers/pkg/setting/constant"
	"lyra-bot-web-crawlers/pkg/utils/httpUtils"
	"net/http"
)

func RandomLyraImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTag(ponyImageConstant.LYRA_PONY)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomBackgroundPonyImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTag(ponyImageConstant.BACKGROUND_PONY)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func GetRandomPonyImageToTag(tag string) (string, string, string) {
	httpClient := httpUtils.GetHttpClientInstance()

	url := setting.DerpibooruURL + "images/random?q=" + tag
	response, err := httpClient.R().Get(url)

	if err != nil {
		fmt.Println(err)
	}

	reader := bytes.NewReader(response.Body())
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
	}

	ponyImageUrl, exists := document.Find("div:nth-child(4) > a:nth-child(3)").Attr("href")

	if exists {
		return ponyImageUrl, response.RawResponse.Request.URL.String(), tag
	}

	return "", "", tag
}
