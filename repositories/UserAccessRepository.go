package repositories

import (
	"inventory/appErr"
	"inventory/iInfrastructures"
	"inventory/iServices"

	"github.com/gomodule/redigo/redis"
)

type UserAccessRepository struct {
	iInfrastructures.ICacheHandler
	iServices.ErrorService
}

func (repository *UserAccessRepository) UserBlackList(userIdentity string) error {
	blackListErr := repository.SetCache("BlackList"+userIdentity, true)
	if blackListErr != nil {
		return blackListErr
	}
	return nil
}

func (repository *UserAccessRepository) UserWhiteList(userIdentity string) error {
	whiteListErr := repository.DeleteCache("BlackList" + userIdentity)
	if whiteListErr != nil {
		return repository.ErrorAndCode(whiteListErr, appErr.ErrUnexpected)
	}
	return nil
}
func (repository *UserAccessRepository) UserIsBlackListed(userIdentity string) (bool, error) {
	var getCacheErr error
	var isBlackListed bool
	isBlackListed, getCacheErr = redis.Bool(repository.GetCache("BlackList" + userIdentity))
	if getCacheErr != nil {
		return false, repository.ErrorAndCode(getCacheErr, appErr.ErrUnexpected)
	}

	return isBlackListed, nil

}
