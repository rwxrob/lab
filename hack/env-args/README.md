# Learning Lab: Insecure Arguments and Environment

This learning lab demonstrates how insecure it is use of arguments and
environment variables for anything secret (passwords, tokens, etc.).
Doing so has been an anti-pattern since UNIX began. In fact, you really
don't need to do anything in this lab so long as you understand the
following perfectly:

> ⚠️
> Anyone can read any arguments or environment variables passed to
> *any* program on a given system, with or without permissions
> to a specific account or group.

For a fun challenge to better understand what this means, run the
following and see if you can discover the unexpected passwords and
secrets revealed in `/proc/` inside the lab container.

First, start the insecure program:

```sh
docker run --name hackme rwxrob/lab-hack-env-args
```

Then, connect to it with `docker exec` using the POSIX shell:

```sh
docker exec -it hackme sh
```

Look through the files in `/proc` using only the simple UNIX/Linux
utilities available from Alpine/BusyBox (but that is all you will need).

Now consider that most API documentation suggests something like the
following (from Twitch):

```bash
curl -X GET 'https://api.twitch.tv/helix/users?login=twitchdev' \
-H 'Authorization: Bearer 2gbdx6oar67tqtcmt49t3wpcgycthx' \
-H 'Client-Id: wbmytr93xzw8zbg0p1izqyzzc5mbiz'
```

An astute, moderately skilled Linux user could write a simple script to
monitor all the new processes from the system and quickly capture the
content of the `cmdline` and `environ` and later examine them for
passwords and tokens. All that is needed is access to the system.

Vulnerabilities like this one are increasing as the base assumption that
there will only ever be one user on the system increases, but this is an
excellent way to elevate access once a hacker is able to get, say, a web
server to send back a file from the system. The restricted `www`
account, for example, *still* has access to *everything* in `/proc`. It
doesn't take much to imagine methods for initial access and escalated
privileges from this. In fact, no access at all may be needed in some
cases, say, if password or token information for another services is
discovered there.

> 💡
> Consider researching the fix for `curl` using another argument besides
`-H` for passing sensitive data.
