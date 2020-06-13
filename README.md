## oneshot

A single-fire HTTP server.

### Synopsis

Start an HTTP server which will only serve files once.
The first client to connect is given the file, all others receive an HTTP 410 Gone response code.

If no file is given, oneshot will instead serve from stdin and hold the clients connection until receiving the EOF character.


```
oneshot [flags]... [file]
```

### Options

```
  -e, --ext string             Extension of file presented to client.
                               If not set, either no extension or the extension of the file will be used,
                               depending on if a file was given.
  -h, --help                   help for oneshot
  -W, --hidden-password        Prompt for password for basic authentication.
                               If a username is not also provided using the -U, --username flag then the client may enter any username.
                               Takes precedence over the -w, --password-file flag
  -m, --mime string            MIME type of file presented to client.
                               If not set, either no MIME type or the mime/type of the file will be user,
                               depending on of a file was given.
  -n, --name string            Name of file presented to client.
                               If not set, either a random name or the name of the file will be used,
                               depending on if a file was given.
  -D, --no-download            Don't trigger browser download client side.
                               If set, the "Content-Disposition" header used to trigger downloads in the clients browser won't be sent.
  -P, --password string        Password for basic authentication.
                               If a username is not also provided using the -U, --username flag then the client may enter any username.
                               If either the -W, --hidden-password or -w, --password-file flags are set, this flag will be ignored.
  -w, --password-file string   File containing password for basic authentication.
                               If a username is not also provided using the -U, --username flag then the client may enter any username.
                               If the -W, --hidden-password flag is set, this flags will be ignored.
  -p, --port string            Port to bind to. (default "8080")
  -q, --quiet                  Don't show info messages.
                               Use -Q, --silent instead to suppress error messages as well.
  -Q, --silent                 Don't show info and error messages.
                               Use -q, --quiet instead to suppress info messages only.
  -t, --timeout duration       How long to wait for client.
                               A value of zero will set the timeout to the max possible value.
      --tls-cert string        Certificate file to use for HTTPS.
                               Key file must also be provided using the --tls-key flag.
      --tls-key string         Key file to use for HTTPS.
                               Cert file must also be provided using the --tls-cert flag.
  -U, --username string        Username for basic authentication.
                               If a password is not also provided using either the -P, --password;
                               -W, --hidden-password; or -w, --password-file flags then the client may enter any password.
  -v, --version                Version for oneshot.
```

###### Auto generated by spf13/cobra on 12-Jun-2020