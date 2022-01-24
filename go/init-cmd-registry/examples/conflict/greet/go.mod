module local/greet

go 1.16

replace local/cmds => ../cmds

replace local/cmd-hi => ../cmd-hi

replace local/cmd-hello => ../cmd-hello

replace local/cmd-hiagain => ../cmd-hiagain

require (
	local/cmd-hello v0.0.0-00010101000000-000000000000
	local/cmd-hi v0.0.0-00010101000000-000000000000
	local/cmd-hiagain v0.0.0-00010101000000-000000000000
	local/cmds v0.0.0-00010101000000-000000000000
)
