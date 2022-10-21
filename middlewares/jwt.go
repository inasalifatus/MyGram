package middlewares

// func CreateToken(userId int) (string, error) {
// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["userId"] = userId
// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(constants.SECRET_JWT))
// }

// func ExtractTokenUserId(ctx *gin.Context) int {
// 	user, _ := ctx.Get("user").(*jwt.Token, true)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		userId := claims["userId"].(int)
// 		return userId
// 	}
// 	return 0
// }
