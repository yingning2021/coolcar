package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEApaQFLgzxzRpOR0sMtagv1YY9vhhW4rt5GWtJ6PzDka8gLArB
Jc01r5IM7Pp4eejohaR7UmWQJAp8hEokrPbtc3ZXo6fKu5Fl9q3KIB8s8awYvDrh
5Lrsk1mdmGAEhsO7oFxudgtLZnHENz60H4gDzMzM5L+1j/fnWXXAsymnHXC91W7Y
JB9WpbpJpKY79xLYLJuNNH+L5EQf40lQDeWmczVkYydEyqBi9NesURx18j6rld7u
K2xDQ7wf5lb+c1+G171ZP/608cbf9IWBNup5sQNBILyeqcPdFRkndGJUUIiVRKt/
+ctRTZfEDsVyDiknsiO/kzDNsM04N1TH2ZgZSwIDAQABAoIBAC4y9Drm415IcwLR
fOcB1O2iNoBZu4obregYE5JHRajRhpCiI0MO3GVuv+os5gNioc/8k2Tk7PIQdrBT
Ga2gZZQpssHzn8j3AdBuooyZBWkWjgOaDL1GIYvrl9gTF9Aasa9FeI22Er6tBoQ+
GfEdd6nciV1X1yUjiMRb4nZWLMU54mq8r9/Mf3+obnCNLp9P82+6Uxe0v1+5/iCE
fvqthgt5zg1DMOD/knIJKQuLMPPgF2dcB0K60SZgfVXqgEwR4hCu8qKKFwD6b+tW
KJyd5blngZFs6Vxnr6FbINSQvn3orvS/BmW3HZHz6m+R8MhC2D3vPmrKeUWZd4nu
CPSLbnkCgYEA/ZG+NGApxwD8x6WNFTpSsQNS031khrCVwVRHtRefErOEMIA1XqrU
+k6RwIbkk5w6LP2WKYPJKkHbWBCPVA3nmYqdpGYacRc3dDuO8YQRFVkDv//yVeRr
txU9Aa5e6HtgL80A7IPxqXQTEj2hiWZx0iCbRkPgA8OidLVn05gk0i0CgYEApzqA
TKWMgR41NVi48pNCY+uJ7AHeA0bLAw5N5fwxXP0bWrGdi0GVKSCvACEornzv7ll+
PsFLHnsKpWByeCbkFjdjP7D49QD78rUN+ohy0ENdgd8OrKMZ4Y5EguJLs3If/sMA
D7e1j3c8yWioGTGzNa6PF+dQnq1ygpaPp0oO3FcCgYA1cKXZe/rSCg88NFPLiYMr
8ztdfyvUhxrIp+6E5/mKg3L0ldCppu9D4ZMuND+wLFjGaptfHHslAMQthy/t0xBg
d6pJn4srEm2JfZPeqqq/CQeVS2fTWlSpPTyiQhGWhYn8CQSM1DH2OJRcX8jPoFuU
oXKYGG353R0744+CNKpt8QKBgQCT8Bv6MmYW/5tAo0mSRxXvgTqVT52RNnp4LJpb
P/yHb95YIFLoE8+Z/7DxI2Ry4FH9gKw/Zg5HW8AyRx1dD8KtqLgjazMCw6kfsG46
WaWAemfpcWPw86T8tjgDtaAUknydivKt9O4oiep7nxs+loocjl1GXzsh9P0da4aQ
DuMQpwKBgQD0bW4linUvweqpl5YVYHZN3G0H3BwQ6ueqC410DjR6BW21GnR03yf6
bQEgjFErJnzfLjTk9unAnYiTpF4wKG5d8uuMq4wYbB/Qz3VN5F7nQDHdZf1kjVVY
cLGSxIe8qldqAo+hQat4HveO0TCvGi2FVeyXk/A2GvqPR20PdYA5tA==
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse prviate key : %v", err)
	}
	g := NewJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tkn, err := g.GenerateToken("5f7c3168e2283aa772e351a3", 2*time.Hour)
	if err != nil {
		// Errorf 测试还能继续下去，不会终止
		t.Errorf("cannot generate token: %v", err)
	}
	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzcyZTM1MWEzIn0.fYLbYqoXBr9rjQtS9vmgZq7BXTKhPoB2IEqgT3jArpjF9Q-WzdmxbnUazc9KyV4KewZiC9aaFfWizBCRRK99TFVV3pRTgNmkwqkKu7xm5Ac0dfrFjR85XZSzzd8KZihK2Udw2xBbaG-_qYIXVUoFRiJyQnJGqB7r51ByizLsIgVbefbZirJv1Fon7YDBlaborbyIyyCUHmv9NsVX4CZA9hh_nitabWsWRTdMrWgoqwMLDx0hi5gZZ1FvxAhd6I4kAZbhHNNU87VGYaxXknEbHZLWujYXhNjhIWRKDGwR_Jq1ELhByWVbA0hfXWKkWUcQA0QAKPlZNCz-GZZOUodSlA"
	if tkn != want {
		t.Errorf("wrong token generated. want : %q;\n got : %q", want, tkn)
	}
}
