package repository

import (
	"errors"
	"sync"

	"github.com/Armageddon6026/zender/pkg/model"
)

type loginServiceRepository struct {
}

type loginServiceManger struct {
	list     map[string]*model.LoginService
	listLock sync.RWMutex
	once     sync.Once
}

var ourLoginServiceManager = new(loginServiceManger)

func getLoginServiceRepository() LoginServiceRepository {
	return &loginServiceRepository{}
}

func (l *loginServiceRepository) SelectLoginServices() model.LoginServiceList {
	manager := l.getLoginServiceManger()
	manager.listLock.RLock()
	defer manager.listLock.RUnlock()

	//Deep copy map to slice
	copyList := make([]model.LoginService, 0, len(manager.list))
	for _, value := range manager.list {
		copyList = append(copyList, *value)
	}
	return copyList
}

func (l *loginServiceRepository) SelectLoginServiceByUuid(uuid string) (model.LoginService, error) {
	manager := l.getLoginServiceManger()
	manager.listLock.RLock()
	defer manager.listLock.RUnlock()

	if loginService, ok := manager.list[uuid]; ok {
		return *loginService, nil
	} else {
		return model.LoginService{}, errors.New(uuid + " is not login")
	}
}

// Add Login Servcie and return Servcie Token
func (l *loginServiceRepository) InsertLoginService(loginService *model.LoginService) error {
	manager := l.getLoginServiceManger()
	manager.listLock.Lock()
	defer manager.listLock.Unlock()

	if _, ok := manager.list[loginService.Uuid]; !ok {
		manager.list[loginService.Uuid] = loginService
		return nil
	} else {
		return errors.New(loginService.Uuid + " has been exist")
	}
}

func (l *loginServiceRepository) UpdateLoginService(loginService *model.LoginService) error {
	manager := l.getLoginServiceManger()
	manager.listLock.Lock()
	defer manager.listLock.Unlock()

	if currentInfo, ok := manager.list[loginService.Uuid]; ok {
		currentInfo.ScanFiles = loginService.ScanFiles
		currentInfo.ScanApplications = loginService.ScanApplications
		return nil
	} else {
		return errors.New(loginService.Uuid + " is not login")
	}
}

func (l *loginServiceRepository) DeleteLoginServiceByUuid(uuid string) {
	manager := l.getLoginServiceManger()
	manager.listLock.Lock()
	defer manager.listLock.Unlock()

	delete(manager.list, uuid)
}

func (*loginServiceRepository) getLoginServiceManger() *loginServiceManger {
	ourLoginServiceManager.once.Do(func() {
		ourLoginServiceManager.list = make(map[string]*model.LoginService)
	})
	return ourLoginServiceManager
}
