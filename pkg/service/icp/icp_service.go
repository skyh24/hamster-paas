package icp

import (
	"hamster-paas/pkg/utils/logger"
	"os/exec"
	"strings"

	"gorm.io/gorm"
)

type IcpService struct {
	db *gorm.DB
}

func NewIcpService(db *gorm.DB) *IcpService {
	return &IcpService{
		db: db,
	}
}

// input command, output result
func executeCmd(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// node statistic
func (n *IcpService) GetVersion() (string, error) {
	out, err := executeCmd("dfx -V")
	if err != nil {
		logger.Errorf("dfx version failed: %s", err)
		return "", err
	}
	res := strings.Fields(out)
	return res[1], nil
}
