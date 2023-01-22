# Pwm - password manager

Pwm is a cross-platform terminal-based password manager.

## Passwords store

Each password lies in its own file. Despite files don't have an extension, 
they are in JSON format. Later this format may be changed
for something different, but the main rule and promise is:
the file should be **in a human-readable format first**.
It means that if you don't have access to Pwm, you can always decrypt your
passwords with another tool that supports bcrypt.

All passwords lie in a passwords store. Passwords store is a simple directory.
This store may have some files required for its work, for example, index file,
or rules for coloring filenames in listview by their prefixes,
but our first implementation will not have any. Any non-password file should
live in `PASSWORD_STORE_DIR/.pwm` directory.

Password store is created automatically when a user adds password with
`pwm add` for the first time and it placed in the user's home directory (on Linux)
or in the `%userprofile%` (on Windows. It's usually looks like `C:\\Users\username`).

##  Usage examples

Add new password: 

```
$/home/usrname>pwm add mail-password
Enter password for mail-password: *********
Done!
```

Add new password, but with duplicating name: 

```
$/home/usrname>pwm add mail-password
Such password already exists. Use pwd upd command to change password
```

Update password:

```
$/home/usrname>pwm upd mail-password
Enter decryption key: *******
Enter new password: *******
```

Copy password for `mail-password` to clipboard:

```
$/home/usrname>pwm mail-password
Enter decryption key: ********
Copied!
```

Show password for `mail-password`:

```
$/home/usrname>pwm show mail-password
Enter decryption key: ********
iamthebest666
```

Get list of passwords:

```
$/home/usrname>pwm
asciidoc-book-decrypt
dropbox
hoster-mysite-ftp
hoster-mysite-cpanel
hoster-personal-blog-ftp
github
mail-gmail
mail-password
mail-fastmail
music-soundcloud
music-deezer
music-spotify
```

Filter passwords: 

```
$/home/usrname>pwm -f mail
mail-gmail
mail-password
mail-fastmail
```

```
$/home/usrname>pwm -f mysi
hoster-mysite-ftp
hoster-mysite-cpanel
```

## Possible issues/ features not designed yet

- Adding additional info for passwords (logins, descriptions etc.)
- Support for git
- No one can't decrypt passwords without decryption key, but anyone able to
  update existed passwords
- Color password names by rules in the list view
- Providing a file with a decryption key and using it instead of manual typing

