package app

import (
	// "bufio"
	// "bytes"
	"fmt"
	"os/exec"
	"strings"

	// "strings"

	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/logger"
)

const (
	AppCmd   string = `c:\windows\system32\inetsrv\appcmd.exe`
	Recycle  string = "recycle"
	AppPool  string = "apppool"
	AppPools string = "apppools"
	List     string = "list"
	IISReset string = "iisreset"
)

type IPool interface {
	RecycleSingleAppPool() (bool, error)
	IISReset() (bool, error)
	GetAll() ([]Pool, error)
}

type Pool struct {
	Name string
}

func (r Pool) RecycleSingleAppPool() (bool, error) {
	result, err := exec.Command(AppCmd, Recycle, AppPool, r.Name).Output()
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error %s", err))
		return false, err
	}
	logger.Log.Info(string(result))
	return true, nil
}

func (r Pool) IISReset() (bool, error) {
	result, err := exec.Command(IISReset).Output()
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error %s", err))
		return false, err
	}
	logger.Log.Info(string(result))
	return true, nil
}

func (r Pool) GetAll() ([]Pool, error) {

	result, err := exec.Command(AppCmd, List, AppPools).Output()
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error %s", err))
		return nil, err
	}

	logger.Log.Info(string(result))
	resultText := string(result)
	arr := strings.Split(strings.ReplaceAll(resultText, "\r\n", "\n"), "\n")
	pools := make([]Pool, 0)
	filterPools := [3]string{"DefaultAppPool", "NET", "Classic"}
	for i := 0; i < len(arr); i++ {
		temps := strings.Split(arr[i], " ")
		if len(temps) > 1 {
			poolName := temps[1]			
			if !Contains(filterPools, poolName) {
				pools = append(pools, Pool{Name: poolName})
				logger.Log.Info(poolName)
			}
		}
	}

	return pools, nil
}

func Contains(s [3]string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a){
			return true
		}
	}
	return false
}
