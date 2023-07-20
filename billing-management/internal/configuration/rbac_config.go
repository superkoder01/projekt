package configuration

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/rbac"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var r rbac.RBAC

func GetRBACConfig() rbac.RBAC {
	return r
}

func LoadRBACConfig(path string) error {
	var rb rbac.Rbac
	rbacFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rbacFile, &rb)
	if err != nil {
		return err
	}

	r = &rb

	return nil
}
