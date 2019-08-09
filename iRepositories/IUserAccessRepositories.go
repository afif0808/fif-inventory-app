package iRepositories

type IUserAccessRepositories interface {
	UserBlackList(userIdentity string) error
	UserWhiteList(userIdentity string) error
	UserIsBlackListed(userIdentity string) (bool, error)
}
