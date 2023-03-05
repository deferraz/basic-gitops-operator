package main
import (
   "fmt"
   "os"
   "os/exec"
   "path"
   "time"
   "github.com/go-git/go-git/v5"
)
func main() {
   timerSec := 5 * time.Second
   gitopsRepo := "https://github.com/PacktPublishing/ArgoCD-in-Practice.git"   localPath := "tmp/"
   pathToApply := "ch01/basic-gitops-operator-config"
   for {
       fmt.Println("start repo sync")
       err := syncRepo(gitopsRepo, localPath)
       if err != nil {
           fmt.Printf("repo sync error: %s", err)
           return
       }
       fmt.Println("start manifests apply")
       err = applyManifestsClient(path.Join(localPath, pathToApply))
       if err != nil {
           fmt.Printf("manifests apply error: %s", err)
       }
       syncTimer := time.NewTimer(timerSec)
       fmt.Printf("\n next sync in %s \n", timerSec)
       <-syncTimer.C
   }
}