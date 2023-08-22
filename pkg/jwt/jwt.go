/*
Copyright 2023 QKP Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

// 密钥（实际应该使用长期保存的密钥）
var secretKey = []byte("kubernetes-secret-key")

// AccessTokenClaims 定义两个自定义的声明类型，用于在 token 中存储自定义信息
type AccessTokenClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// ExecuteJwtToken  定义端口扫描执行器
func ExecuteJwtToken() *cobra.Command {
	// cpu 定义网络链路跟踪命令
	var jwt = &cobra.Command{
		Use:   "token",
		Short: "generate jwt token information.",
		Long:  "generate jwt token information.",
		Run:   runner,
	}
	// 初始化命令
	jwt.Flags().StringP("port", "p", "8080", "The http web service listening port.")
	jwt.Flags().StringP("username", "u", "admin", "The username of the token.")
	jwt.Flags().StringP("password", "w", "password", "The password of the token.")
	jwt.Flags().IntP("expire", "e", 7200, "The token expire time,  default 7200s.")
	jwt.Flags().StringP("issuer", "i", "issuer", "The issuer of the token.")
	return jwt
}

func runner(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetString("port")
	issuer, _ := cmd.Flags().GetString("issuer")
	expire, _ := cmd.Flags().GetInt("expire")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")

	router := gin.Default()
	router.GET("/token", func(context *gin.Context) {
		// 创建 Access Token
		accessTokenClaims := AccessTokenClaims{
			Username: username,
			Password: password,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(expire) * time.Second).Unix(), // 2小时有效期
				Issuer:    issuer,
			},
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
		accessTokenString, err := accessToken.SignedString(secretKey)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create access token"})
			return
		}

		// 创建 Refresh Token
		refreshTokenClaims := RefreshTokenClaims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(), // 30天有效期
				Issuer:    issuer,
			},
		}
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
		refreshTokenString, err := refreshToken.SignedString(secretKey)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create refresh token"})
			return
		}
		// 返回生产的token信息
		context.IndentedJSON(http.StatusOK, gin.H{
			"access_token": gin.H{
				"access_token": accessTokenString,
				"expire_time":  time.Now().Add(time.Duration(expire) * time.Second).Unix(),
			},
			"refresh_token": gin.H{
				"refresh_token": refreshTokenString,
				"expire_time":   time.Now().Add(time.Hour * 24 * 30).Unix(),
			},
		})
	})
	router.Run(":" + port)
}
