package storageclass

import (
    "encoding/json"
    "fmt"
    "os/exec"
)

type StorageClass struct {
    Items []struct {
        Metadata struct {
            Name        string `json:"name"`
            Annotations struct {
                DefaultClass string `json:"storageclass.kubernetes.io/is-default-class"`
            } `json:"annotations"`
        } `json:"metadata"`
    } `json:"items"`
}

func (s *StorageClass) Unmarshal(data []byte) error {
    return json.Unmarshal(data, s)
}

func GetAllStorageClasses() error {
    out, err := exec.Command("kubectl", "get", "sc", "-o", "json").Output()
    if err != nil {
        return err
    }

    var scs StorageClass
    err = json.Unmarshal(out, &scs)
    if err != nil {
        return err
    }

    for _, sc := range scs.Items {
        fmt.Printf("%s: is_default=%s\n", sc.Metadata.Name, sc.Metadata.Annotations.DefaultClass)
    }

    return nil
}

func UnsetDefault(sc string) error {
    cmd := exec.Command("kubectl", "patch", "storageclass", sc,
        "-p", `{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false"}}}`)
    err := cmd.Run()
    if err != nil {
        return err
    }

    return nil
}

func SetDefault(sc string) error {
    cmd := exec.Command("kubectl", "patch", "storageclass", sc,
        "-p", `{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}`)
    err := cmd.Run()
    if err != nil {
        return err
    }
    return nil
}
