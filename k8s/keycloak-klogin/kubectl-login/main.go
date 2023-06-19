package main

import (
	"log"
	"os"

	"github.com/rwxrob/kubectl-login/internal"
)

var cinfo = map[string]internal.ClusterInfo{

	`prod`: internal.ClusterInfo{
		Name:          `prod`,
		APIServerURL:  `https://192.168.39.76:8443`,
		OIDCIssuerURL: `https://home.rwx.gg:8443/realms/prod`,
		ClientID:      `minikube-prod`,
		ClientSecret:  `x46aVnFlygaEdHBom1200AZm37ZTWPhe`,
	},

	`dev`: internal.ClusterInfo{
		Name:          `dev`,
		APIServerURL:  `https://192.168.72.28:8443`,
		OIDCIssuerURL: `https://home.rwx.gg:8443/realms/dev`,
		ClientID:      `minikube-dev`,
		ClientSecret:  `12l3IMNHqDvqpp5gm0292g1Hs7SzWgyX`,
	},

	`inf`: internal.ClusterInfo{
		Name:          `inf`,
		APIServerURL:  `https://192.168.83.223:8443`,
		OIDCIssuerURL: `https://home.rwx.gg:8443/realms/inf`,
		ClientID:      `minikube-inf`,
		ClientSecret:  `6yw8EH1dVAx9XYKTM2mK2FZZmFhD19Hz`,
	},
}

func main() {

	cluster := `prod`
	if len(os.Args) > 1 {
		cluster = os.Args[1]
	}
	ci, has := cinfo[cluster]
	if !has {
		log.Fatal(`unable to locate info for cluster: %v`, cluster)
	}

	user, pass := internal.GetK8SUserPass()

	grant, err := internal.ReqOIDCPassAuth(
		user, pass, ci.OIDCIssuerURL, ci.ClientID, ci.ClientSecret,
	)
	if err != nil {
		log.Fatal(err)
	}

	internal.Exec(`kubectl`, `config`, `set-cluster`, ci.Name,
		`--server`, ci.APIServerURL,
		`--certificate-authority`, `/shared/ca.crt`, // TODO comment this line
	)

	internal.Exec(
		`kubectl`, `config`, `set-context`, ci.Name,
		`--cluster`, ci.Name,
		`--user`, ci.Name,
		`--namespace`, `default`, // TODO change back to user,
	)
	internal.Exec(`kubectl`, `config`, `use-context`, ci.Name)

	token, _ := grant[`id_token`].(string)
	internal.Exec(`kubectl`, `config`, `set-credentials`, ci.Name, `--token`, token)

}
