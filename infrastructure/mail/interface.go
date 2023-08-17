package mail

type IMail interface {
	SendingNativeMail(*Mail) error
}
