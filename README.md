Run frontend http server in kubernetes:

```
$ kubectl apply -f k8s/frontend-deployment.yaml
$ curl $(minikube ip)/
```

You should see `Hello` Printed in terminal
