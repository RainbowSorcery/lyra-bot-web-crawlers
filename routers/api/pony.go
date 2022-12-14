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
	url2 "net/url"
)

func RandomLyraImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.LYRA_PONY)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomBackgroundPonyImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.BACKGROUND_PONY)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomTSImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.TWILIGHT_SPARKLE)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomRDImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.RAINBOW_DASH)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomRRImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.RR)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomSnowdropImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.SNOWDROP)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomNyxImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.NYX)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomPinkiePieImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.PINKIE_PIE)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomFlutterShyPieImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.FLUTTERSHY)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomAppleJackPieImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.APPLEJACK)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func RandomLunaImage(context *gin.Context) {
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(ponyImageConstant.LUNA)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func GetRandomPonyImageToTag(context *gin.Context) {
	requestTag := context.DefaultQuery("tag", "")
	derpibooruURL, PonyImageURL, tag := GetRandomPonyImageToTagFun(requestTag)
	context.JSON(http.StatusOK, setting.SuccessOfData(entity.PonyImageEntity{
		DerpibooruURL: derpibooruURL,
		PonyImageURL:  PonyImageURL,
		Tag:           tag,
	}))
}

func GetRandomPonyImageToTagFun(tag string) (string, string, string) {
	httpClient := httpUtils.GetHttpClientInstance()

	tagUrlEscape := url2.QueryEscape(tag)

	url := setting.DerpibooruURL + "images/random?q=" + tagUrlEscape
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
