Run envoy with config file:

```
$ envoy -c envoy/frontend.json
```

The envoy version I used: `envoy version: ab30cafa0c033f8ef77c54dd98badea62b78506e/1.12.0/clean-getenvoy-f13ee1c-envoy/RELEASE/BoringSSL`

---

Run frontend app:

```
go run frontend/main.go
```

---

Verify envoy is proxying frontend request by

```
$ curl localhost:8080
```
