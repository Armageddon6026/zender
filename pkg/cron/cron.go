package cron

import (
	"time"

	"github.com/Armageddon6026/zender/pkg/notification"
	"github.com/Armageddon6026/zender/pkg/repository"

	"github.com/roylee0704/gron"
)

type Cron struct {
	Task                   *gron.Cron
	loginServiceManager    *notification.LoginServiceManager
	loginServiceRepository repository.LoginServiceRepository
}

func New(loginServiceManager *notification.LoginServiceManager, loginServiceRepository repository.LoginServiceRepository) *Cron {
	cron := &Cron{
		Task:                   gron.New(),
		loginServiceManager:    loginServiceManager,
		loginServiceRepository: loginServiceRepository,
	}
	// Add cron job
	cron.Task.AddFunc(gron.Every(3*time.Second), cron.sendLoginServiceStatus)

	return cron
}

func (c *Cron) sendLoginServiceStatus() {
	c.loginServiceManager.SendMessage(c.loginServiceRepository.SelectLoginServices())
}
