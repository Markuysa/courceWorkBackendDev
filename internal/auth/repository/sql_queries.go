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
	queryGetAdmin = `
	select id,
			username,
			password,
			otp_secret,
			tg_chat
	from admin.admin
	where username = $1;
`
	querySaveAdmin = `
	insert into admin.admin(
	                        username,
	                        password,
	                        otp_secret
	)values(
	        $1,$2,$3
	)
	
`
	queryGetOTP = `
	select otp_secret from "user".user where username = $1;
`
	queryGetAdminOTP = `
	select otp_secret from admin.admin where username = $1;
`
)
