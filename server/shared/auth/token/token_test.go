package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApaQFLgzxzRpOR0sMtagv
1YY9vhhW4rt5GWtJ6PzDka8gLArBJc01r5IM7Pp4eejohaR7UmWQJAp8hEokrPbt
c3ZXo6fKu5Fl9q3KIB8s8awYvDrh5Lrsk1mdmGAEhsO7oFxudgtLZnHENz60H4gD
zMzM5L+1j/fnWXXAsymnHXC91W7YJB9WpbpJpKY79xLYLJuNNH+L5EQf40lQDeWm
czVkYydEyqBi9NesURx18j6rld7uK2xDQ7wf5lb+c1+G171ZP/608cbf9IWBNup5
sQNBILyeqcPdFRkndGJUUIiVRKt/+ctRTZfEDsVyDiknsiO/kzDNsM04N1TH2ZgZ
SwIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name: "valid_token",
			tkn:  "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzcyZTM1MWEzIn0.fYLbYqoXBr9rjQtS9vmgZq7BXTKhPoB2IEqgT3jArpjF9Q-WzdmxbnUazc9KyV4KewZiC9aaFfWizBCRRK99TFVV3pRTgNmkwqkKu7xm5Ac0dfrFjR85XZSzzd8KZihK2Udw2xBbaG-_qYIXVUoFRiJyQnJGqB7r51ByizLsIgVbefbZirJv1Fon7YDBlaborbyIyyCUHmv9NsVX4CZA9hh_nitabWsWRTdMrWgoqwMLDx0hi5gZZ1FvxAhd6I4kAZbhHNNU87VGYaxXknEbHZLWujYXhNjhIWRKDGwR_Jq1ELhByWVbA0hfXWKkWUcQA0QAKPlZNCz-GZZOUodSlA",
			now:  time.Unix(1516239122, 0),
			want: "5f7c3168e2283aa772e351a3",
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzcyZTM1MWEzIn0.fYLbYqoXBr9rjQtS9vmgZq7BXTKhPoB2IEqgT3jArpjF9Q-WzdmxbnUazc9KyV4KewZiC9aaFfWizBCRRK99TFVV3pRTgNmkwqkKu7xm5Ac0dfrFjR85XZSzzd8KZihK2Udw2xBbaG-_qYIXVUoFRiJyQnJGqB7r51ByizLsIgVbefbZirJv1Fon7YDBlaborbyIyyCUHmv9NsVX4CZA9hh_nitabWsWRTdMrWgoqwMLDx0hi5gZZ1FvxAhd6I4kAZbhHNNU87VGYaxXknEbHZLWujYXhNjhIWRKDGwR_Jq1ELhByWVbA0hfXWKkWUcQA0QAKPlZNCz-GZZOUodSlA",
			now:     time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name: "wrong_signature",
			// 用户伪造token
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzcyZTM1MWEyIn0.fYLbYqoXBr9rjQtS9vmgZq7BXTKhPoB2IEqgT3jArpjF9Q-WzdmxbnUazc9KyV4KewZiC9aaFfWizBCRRK99TFVV3pRTgNmkwqkKu7xm5Ac0dfrFjR85XZSzzd8KZihK2Udw2xBbaG-_qYIXVUoFRiJyQnJGqB7r51ByizLsIgVbefbZirJv1Fon7YDBlaborbyIyyCUHmv9NsVX4CZA9hh_nitabWsWRTdMrWgoqwMLDx0hi5gZZ1FvxAhd6I4kAZbhHNNU87VGYaxXknEbHZLWujYXhNjhIWRKDGwR_Jq1ELhByWVbA0hfXWKkWUcQA0QAKPlZNCz-GZZOUodSlA",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verify(c.tkn)
			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}

			if c.wantErr && err == nil {
				t.Errorf("want error; got no error")
			}

			if accountID != c.want {
				t.Errorf("wrong account id. want: %q, got : %q", c.want, accountID)

			}

		})
	}
	//tkn := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWY3YzMxNjhlMjI4M2FhNzcyZTM1MWEzIn0.fYLbYqoXBr9rjQtS9vmgZq7BXTKhPoB2IEqgT3jArpjF9Q-WzdmxbnUazc9KyV4KewZiC9aaFfWizBCRRK99TFVV3pRTgNmkwqkKu7xm5Ac0dfrFjR85XZSzzd8KZihK2Udw2xBbaG-_qYIXVUoFRiJyQnJGqB7r51ByizLsIgVbefbZirJv1Fon7YDBlaborbyIyyCUHmv9NsVX4CZA9hh_nitabWsWRTdMrWgoqwMLDx0hi5gZZ1FvxAhd6I4kAZbhHNNU87VGYaxXknEbHZLWujYXhNjhIWRKDGwR_Jq1ELhByWVbA0hfXWKkWUcQA0QAKPlZNCz-GZZOUodSlA"
	//jwt.TimeFunc = func() time.Time {
	//	return time.Unix(1516239122, 0)
	//}
	//
	//accountID, err := v.Verify(tkn)
	//if err != nil {
	//	t.Errorf("verification failed: %v", err)
	//}
	//want := "5f7c3168e2283aa772e351a3"
	//if accountID != want {
	//	t.Errorf("wrong account id. want: %q, got : %q", want, accountID)
	//}
}
