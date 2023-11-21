package emailSend

func SendEmail(pathDirExcel string, arrPathFileExcel []string) {
	entityEmailSend := MakeEntityEmailSend(pathDirExcel, arrPathFileExcel)
	entityEmailSend.SendEmail()
	return
}
