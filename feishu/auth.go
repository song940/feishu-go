package feishu

import (
	"encoding/json"
)

type AccessTokenResponse struct {
	FeishuResponseBase
	Expire      int    `json:"expire"`
	AccessToken string `json:"access_token"`
}

type TenantAccessTokenResponse struct {
	AccessTokenResponse
	TenantAccessToken string `json:"tenant_access_token"`
}

type AppAccessTokenResponse struct {
	AccessTokenResponse
	AppAccessToken string `json:"app_access_token"`
}

type UserAccessTokenResponse struct {
	AccessTokenResponse
	Data map[string]interface{} `json:"data"`
}

// 商店应用获取 tenant_access_token
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token
func (c *FeishuClient) GetTenantAccessToken() (out AccessTokenResponse, err error) {
	data, err := c.RequestWithAppSecret("/auth/v3/tenant_access_token")
	if err != nil {
		return
	}
	var resp TenantAccessTokenResponse
	err = json.Unmarshal(data, &resp)
	out.AccessToken = resp.TenantAccessToken
	return
}

// 自建应用获取 tenant_access_token
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal
func (c *FeishuClient) GetTenantAccessTokenInternal() (out AccessTokenResponse, err error) {
	data, err := c.RequestWithAppSecret("/auth/v3/tenant_access_token/internal")
	if err != nil {
		return
	}
	var resp TenantAccessTokenResponse
	err = json.Unmarshal(data, &resp)
	out.AccessToken = resp.TenantAccessToken
	return out, err
}

// 商店应用获取 app_access_token
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_access_token
func (c *FeishuClient) GetAppAccessToken() (out AccessTokenResponse, err error) {
	data, err := c.RequestWithAppSecret("/auth/v3/app_access_token")
	if err != nil {
		return
	}
	var resp AppAccessTokenResponse
	err = json.Unmarshal(data, &resp)
	out.AccessToken = resp.AppAccessToken
	return
}

// 自建应用获取 app_access_token
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_access_token_internal
func (c *FeishuClient) GetAppAccessTokenInternal() (out AccessTokenResponse, err error) {
	data, err := c.RequestWithAppSecret("/auth/v3/app_access_token/internal")
	if err != nil {
		return
	}
	var resp AppAccessTokenResponse
	err = json.Unmarshal(data, &resp)
	out.AccessToken = resp.AppAccessToken
	return out, err
}

// 获取 user_access_token
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/access_token/create
func (c *FeishuClient) GetUserAccessToken() (out UserAccessTokenResponse, err error) {
	data, err := c.RequestWithAppSecret("/authen/v1/access_token")
	if err != nil {
		return
	}
	var resp AppAccessTokenResponse
	err = json.Unmarshal(data, &resp)
	return out, err
}

// 刷新 user_access_token
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/refresh_access_token/create
func (c *FeishuClient) RefreshUserAccessToken(refreshToken string, grantType string) {
	c.RequestWithAccessToken("/authen/v1/refresh_access_token", map[string]string{
		"refresh_token": refreshToken,
		"grant_type":    grantType,
	})
}

// 重新获取 app_ticket
// https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/app_ticket_resend
func (c *FeishuClient) ResendAppTicket() (out FeishuResponseBase, err error) {
	data, err := c.RequestWithAppSecret("/auth/v3/app_ticket/resend")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &out)
	return
}
