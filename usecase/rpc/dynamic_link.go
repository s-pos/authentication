package rpc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/s-pos/go-utils/utils/request"
)

type requestDynamicLink struct {
	DynamicLinkInfo *DynamicLinkInfo `json:"dynamicLinkInfo"`
	Suffix          string           `json:"suffix,omitempty"`
}

type responseDynamicLink struct {
	ShortLink   string                   `json:"shortLink"`
	PreviewLink string                   `json:"previewLink"`
	Error       responseErrorDynamicLink `json:"error,omitempty"`
}

type DynamicLinkInfo struct {
	DomainUriPrefix string          `json:"domainUriPrefix"`
	Link            string          `json:"link"`
	Android         *androidInfo    `json:"androidInfo,omitempty"`
	Ios             *iosInfo        `json:"iosInfo,omitempty"`
	Navigation      *navigationInfo `json:"navigationInfo,omitempty"`
	Analytics       *analyticInfo   `json:"analyticsInfo,omitempty"`
}

type androidInfo struct {
	AndroidPackageName           string `json:"androidPackageName,omitempty"`
	AndroidFallbackLink          string `json:"androidFallback,omitempty"`
	AndroidMinPackageVersionCode string `json:"androidMinPackageVersionCode,omitempty"`
}

type iosInfo struct {
	IosBundleId         string `json:"iosBundleId,omitempty"`
	IosFallbackLink     string `json:"iosFallbackLink,omitempty"`
	IosCustomScheme     string `json:"iosCustomScheme,omitempty"`
	IosIpadFallbackLink string `json:"iosIpadFallbackLink,omitempty"`
	IosAppStoreIId      string `json:"iosAppStoreId,omitempty"`
}

type navigationInfo struct {
	EnableForcedRedirect bool `json:"enableForcedRedirect,omitempty"`
}

type analyticInfo struct {
	GooglePlayAnalytics *googlePlayAnalytics `json:"googlePlayAnalytics,omitempty"`
	ItunesAnalytics     *itunesAnalytics     `json:"itunesConnectAnalytics,omitempty"`
}

type googlePlayAnalytics struct {
	UtmSource   string `json:"utmSource,omitempty"`
	UtmMedium   string `json:"utmMedium,omitempty"`
	UtmCampaign string `json:"utmCampaign,omitempty"`
	UtmTerm     string `json:"utmTerm,omitempty"`
	Gclid       string `json:"gclid,omitempty"`
}

type itunesAnalytics struct {
	At string `json:"at,omitempty"`
	Ct string `json:"ct,omitempty"`
	Mt string `json:"mt,omitempty"`
	Pt string `json:"pt,omitempty"`
}

type responseErrorDynamicLink struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

type DynamicLinkData struct {
	Email string `json:"email"`
	Token string `json:"token"` // token can be real token (for forgot password) or can be OTP code for verification
}

type DynamicLinkType string

const (
	Verification   DynamicLinkType = "verification"
	ForgotPassword DynamicLinkType = "forgot_password"
)

func (dl *authClient) CreateDynamicLink(ctx context.Context, data *DynamicLinkData, dtype DynamicLinkType) (string, error) {
	var (
		url            = os.Getenv("URL_WEB")
		domainPrefix   = os.Getenv("DYNAMIC_LINK_DOMAIN_PREFIX")
		reqUrl         = fmt.Sprintf("%s%s?key=%s", os.Getenv("DYNAMIC_LINK_API"), os.Getenv("DYNAMIC_LINK_API_SHORTLINK"), os.Getenv("DYNAMIC_LINK_API_KEY"))
		androidPackage = os.Getenv("ANDROID_PACKAGE_NAME")
		iosBundle      = os.Getenv("IOS_BUNDLE_NAME")
		shortLink      string
		path           string
		err            error
		header         map[string]string
		req            requestDynamicLink
		res            responseDynamicLink
	)

	data.Email = base64.StdEncoding.EncodeToString([]byte(data.Email))
	switch dtype {
	case Verification:
		path = fmt.Sprintf("/v?token=%s&e=%s", data.Token, data.Email)
	case ForgotPassword:
		path = fmt.Sprintf("/forgot-password?token=%s&e=%s", data.Token, data.Email)
	default:
		err = fmt.Errorf("type dynamic link not found. please input the correct type")
		return shortLink, err
	}

	// setup payload
	androidInfo := &androidInfo{
		AndroidPackageName: androidPackage,
	}
	iosInfo := &iosInfo{
		IosBundleId: iosBundle,
	}
	dynamicLinkInfo := &DynamicLinkInfo{
		DomainUriPrefix: domainPrefix,
		Link:            fmt.Sprintf("%s%s", url, path),
		Android:         androidInfo,
		Ios:             iosInfo,
	}

	req = requestDynamicLink{
		DynamicLinkInfo: dynamicLinkInfo,
	}
	payload, err := json.Marshal(req)
	if err != nil {
		return shortLink, err
	}
	// set header
	header = map[string]string{
		"Content-Type": "application/json",
	}

	client := request.NewHTTPClient(dl.httpClient)
	reqApi := client.Request(header, reqUrl)

	resBody, statusCode, err := reqApi.Post(ctx, payload)
	if err != nil {
		return shortLink, err
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return shortLink, err
	}

	if statusCode >= 400 {
		err = fmt.Errorf("error when create dynamic link %v", res.Error)
		return shortLink, err
	}

	shortLink = res.ShortLink
	return shortLink, err
}
