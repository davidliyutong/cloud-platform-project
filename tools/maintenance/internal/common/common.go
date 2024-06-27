package common

import (
    "encoding/json"
    "log"
    "os"
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

func PortForward(namespace, serviceName string, port string) {
    cmd := exec.Command("kubectl", "-n", namespace, "port-forward", "service/"+serviceName, port)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Start()
    if err != nil {
        log.Fatalf("Command failed to start: %v", err)
    }
    err = cmd.Wait()
    if err != nil {
        log.Fatalf("Command finished with error: %v", err)
    }
}
