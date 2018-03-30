# ssh1d
ssh1d - fake SSH-1 daemon

The daemon uses [systemd socket activation](http://0pointer.de/blog/projects/socket-activation.html) and can be enabled with systemctl:
```
systemctl enable ssh1d.socket
systemctl start ssh1d.socket
```
The logged connection attempts can be shown with journalctl:
```
journalctl -u ssh1d.service
```
The daemon also supports traditional socket creation allowing to run it on a non-systemd OS

(c) 2018 Andreas Jaggi - [LICENSE](LICENSE)
