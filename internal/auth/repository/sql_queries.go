package repository

const (
	querySaveOTPSecret = `
	update "user".user
	set otp_secret = $1
	where username = $2
`
	querySaveUser = `
	insert into "user".user(
	                        username,
	                        password,
	                        otp_secret
	)values(
	        $1,$2,$3
	)
`
	queryGetUser = `
	select id,
			username,
			password,
			otp_secret,
			tg_chat
	from "user".user
	where username = $1;
`
	queryGetOTP = `
	select otp_secret from "user".user where username = $1;
`
)
