A repo for simple listener app, written on Golang and ready for deployment



## Prerequisites
To run all mentioned tests, it is necessary to install the following tools:  
[docker](https://docs.docker.com/engine/install/)  
[kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)  
  
## Build and run image locally
After cloning the repository, change directory into app directory:
```
cd listener
```
  
Build an image with your own tag:
```
docker build -f Dockerfile -t << your custom tag >>  .
```
  
Run the image locally on docker/podman:
```
docker run -d -p 8080:8080 << your custom tag >>
```
  
Test connection making a request via curl:
```
curl http://localhost:8080/
```
Sample output:
```
localhost:8080
2023-06-18 11:10:04.990397178 +0000 UTC m=+6.786366421
```

Back to repository root:
```
cd ..
```

## Testing in k8s environment

Run kind test cluster:
```
kind create cluster --config kind/kind-cluster-config.yaml
```

Apply Contour for ingress implementation along with kind specific patches:
```
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml
```
```
kubectl patch daemonsets -n projectcontour envoy -p '{"spec":{"template":{"spec":{"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/control-plane","operator":"Equal","effect":"NoSchedule"},{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}'
```

Wait a couple of minutes while envoy pods redeploy.

Application specs use image from dockerHub (docker.io/uhivik/go-listener:0.2)  
If you need to test your own build, please amend image in [Deployment spec](manifests/deployment.yaml)

Run the application:
```
kubectl apply -f manifests
```

Test connection making a request via curl:
```
curl localhost/
```

Sample output:
```
localhost
2023-06-19 09:23:39.407320926 +0000 UTC m=+221.775006894%                                                                                                             
```