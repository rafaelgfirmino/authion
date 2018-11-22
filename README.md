## Authion

Authenticator writing in GO and using Mysql.
This project try to use faithfully the concepts clean architecture.

- [ ] Authentication

### Create Private and Public keys to JWT

```shell
#create Private Key
openssl genrsa -out authion.rsa 1024

#create Public Key
openssl rsa -in authion.rsa -pubout > authion.rsa.pub
```