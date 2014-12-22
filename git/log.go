package git

//import "os/exec"


//func execGitLog(repoDir, remoteName, branchName string) ([]byte, error) {
//	if branchName == "" {
//		branchName = "HEAD"
//	}
//
//	var cmd *exec.Cmd = nil
//	if remoteName == "" {
//		cmd = exec.Command("git", "rev-list", `--since="2 months ago"`, `--header`, branchName)
//	} else {
//		cmd = exec.Command("git", "rev-list"
