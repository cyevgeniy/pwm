# Pwm - password manager

This is just for me, I don't encourage anyone to use it.

## Why

I really like KeepassXC, but copying/adding  passwords
is something that I want to perform very quick, **BLAZINGLY FAST**.

And I think it can be done better by terminal-based program, though,
everything else is kind of sucks in terminal-based programs, but that's
another story.

There is the [pass](https://www.passwordstore.org) bash script on
Linux available, and I like the idea behind it , but I use Windows, so I've
created `pwm`.

## How it works

In general, like the `pass` program - it keeps each password in an encrypted
file.
Encryption is done by 
[Protonmail's Openpgp library](https://github.com/ProtonMail/gopenpgp).

Also, for each password you can set meta-information (like logins, urls, whatever).
Meta information is encrypted too, and saved with the `.meta` file extension.

All passwords are kept in the "Store" - a directory in your home dir.
On Windows, it's in `C:\\Users\Username\pwm`. 

## Usage examples

Add new password: 

```
pwm add email-gmail
```

Add new password and additional meta-information:

```
pwm add email-gmail -m "email: myEmail@gmail.com"
```

Show list of passwords:

```
pwm ls
```

Show list of passwords which satisfy specified search filter:

```
pwm ls -f email
```

Copy password to the clipboard:

```
pwm copy email-gmail
```

Show password:

```
pwm show email-gmail
```

Show meta-information:
```
pwm show email-gmail.meta
```

Update password:

```
pwm upd email-gmail
```

Update meta-information:

```
pwm upd email-gmail.meta
```
