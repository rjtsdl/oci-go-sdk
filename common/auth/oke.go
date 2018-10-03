package auth

import (
	"fmt"
)

func RewriteMetadataURLs(ip string) {
	regionURL = fmt.Sprintf(`http://%s/opc/v1/instance/region`, ip)
	leafCertificateURL = fmt.Sprintf(`http://%s/opc/v1/identity/cert.pem`, ip)
	leafCertificateKeyURL = fmt.Sprintf(`http://%s/opc/v1/identity/key.pem`, ip)
	intermediateCertificateURL = fmt.Sprintf(`http://%s/opc/v1/identity/intermediate.pem`, ip)
	return
}
