package pv

import (
    "encoding/json"
    "log"
    "os/exec"
    "time"
)

type PV struct {
    Items []struct {
        Metadata struct {
            Name              string `json:"name"`
            CreationTimestamp string `json:"creationTimestamp"`
        } `json:"metadata"`
        Status struct {
            Phase string `json:"phase"`
        } `json:"status"`
    } `json:"items"`
}

func (p *PV) Unmarshal(data []byte) error {
    return json.Unmarshal(data, p)
}

func (p *PV) IsOldAndReleased() bool {
    for _, item := range p.Items {
        if item.Status.Phase == "Released" {
            layout := "2006-01-02T15:04:05Z"
            creationTime, _ := time.Parse(layout, item.Metadata.CreationTimestamp)
            timeDifference := time.Now().Sub(creationTime)

            if timeDifference.Hours() >= 24 {
                return true
            }
        }
    }
    return false
}

func DeleteOldPVs(dryRun bool) error {
    cmd := exec.Command("kubectl", "get", "pv", "-o", "json")

    stdout, err := cmd.Output()
    if err != nil {
        log.Printf(err.Error())
        return err
    }

    var pvs PV
    json.Unmarshal(stdout, &pvs)

    for _, item := range pvs.Items {
        if item.Status.Phase == "Released" {
            layout := "2006-01-02T15:04:05Z"
            creationTime, _ := time.Parse(layout, item.Metadata.CreationTimestamp)
            timeDifference := time.Now().Sub(creationTime)

            if timeDifference.Hours() >= 24 {
                log.Printf("Deleting %s as it is older than 1 day.\n", item.Metadata.Name)
                if dryRun {
                    log.Printf("[Dry-Run]Deleting %s as it is older than 1 day.\n", item.Metadata.Name)
                } else {
                    log.Printf("Deleting %s as it is older than 1 day.\n", item.Metadata.Name)
                    cmd := exec.Command("kubectl", "delete", "pv", item.Metadata.Name)
                    _, err := cmd.Output()
                    if err != nil {
                        log.Printf(err.Error())
                        return err
                    }
                }
            }
        } else {
            log.Printf("Skipping %s as it is not in Released state.\n", item.Metadata.Name)
        }
    }
    return nil
}
