package sacredential

import (
    "encoding/base64"
    "encoding/json"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
)

type SACredential struct {
    Data struct {
        Token string `json:"token"`
        CA    string `json:"ca.crt"`
    } `json:"data"`
}

func (s *SACredential) Unmarshal(data []byte) error {
    return json.Unmarshal(data, s)
}

func (s *SACredential) DecodeToken() ([]byte, error) {
    return base64.StdEncoding.DecodeString(s.Data.Token)
}

func (s *SACredential) DecodeCA() ([]byte, error) {
    return base64.StdEncoding.DecodeString(s.Data.CA)
}

func DownloadCredentials(namespace, serviceAccount, credsDir string) error {
    err := os.MkdirAll(credsDir, os.ModePerm)
    if err != nil {
        return err
    }

    credCmd := exec.Command("kubectl", "-n", namespace, "get", "secret", serviceAccount+"-secret", "-o", "json")
    credBytes, err := credCmd.Output()
    if err != nil {
        return err
    }
    var sa SACredential
    json.Unmarshal(credBytes, &sa)

    caDecoded, err := base64.StdEncoding.DecodeString(string(sa.Data.CA))
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(filepath.Join(credsDir, "ca.crt"), caDecoded, 0644)
    if err != nil {
        return err
    }

    tokenDecoded, err := base64.StdEncoding.DecodeString(string(sa.Data.Token))
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(filepath.Join(credsDir, "token"), tokenDecoded, 0644)
    if err != nil {
        return err
    }

    return nil
}
