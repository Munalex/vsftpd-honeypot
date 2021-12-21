## vsftpd 2.3.4 honeypot

a simple honeypot written in GO using vsftpd 2.3.4 backdoor

To do:
- [ ] Port 6200 socket than original backdoor
- [ ] Multithreading
- [ ] Improve FTP ansawers


```
In linux ports under 1024 need root.
I use 2121 and redirect to 21 on the NAT of the gateway.
```


------------

## Compiling
1. Clone the repo to your computer
2. install golang in your computer
3. Navigate to the vsftpd-honeypot root directory and run `go build`. This command build the honeypot for the architecture of your computer
