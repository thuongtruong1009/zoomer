package mail

type IMail interface {
	SendingMail(*Mail) error
}
