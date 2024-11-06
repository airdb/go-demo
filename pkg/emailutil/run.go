package emailutil

func RunSendEmail() {
	type args struct {
		toEmails string
		subject  string
		content  string
	}

	tests := []struct {
		name string
		args args
	}{
		{``, args{"dean@airdb.com", "hello", "test mail"}},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		SendEmail(tt.args.toEmails, tt.args.subject, tt.args.content)
	}
}
