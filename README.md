# Effective work with SQL in Go
- [Go: ефективна робота з SQL](https://dou.ua/forums/topic/34806/)

### Support Ukraine 🇺🇦
- Help Ukraine via [SaveLife fund](https://savelife.in.ua/en/donate-en/)
- Help Ukraine via [Dignitas fund](https://dignitas.fund/donate/)
- Help Ukraine via [National Bank of Ukraine](https://bank.gov.ua/en/news/all/natsionalniy-bank-vidkriv-spetsrahunok-dlya-zboru-koshtiv-na-potrebi-armiyi)
- More info on [war.ukraine.ua](https://war.ukraine.ua/) and [MFA of Ukraine](https://twitter.com/MFA_Ukraine)

### Testing
```bash
make env-up
make docker-go-version
make docker-pg-version
make migrate-up
make go-test
make env-down
```
```bash
make env-up
time make test
```
```text
real	0m16.016s
user	0m0.168s
sys	0m0.079s
```

### Tools
* [github.com/sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc)
* [https://github.com/pressly/goose](github.com/pressly/goose)
