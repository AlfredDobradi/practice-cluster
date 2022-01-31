# practice-cluster

I created this repository for my own practicing purposes to build a simple app with various dependencies to be ran inside a Kubernetes cluster.

The tags will progressively introduce more and more concepts like deployments, services, configuration management, etc.

### Current tag: v1.1.0

#### The application

Takes a random element from a string slice of colors and returns it as a JSON along the version.
The random method is called with an integer bigger than the size of the slice to emulate errors on random (every 1 in length+1 chance)
If any errors occur during marshaling the data, it returns the error instead of the color.

Docker image is pushed to Docker HUB (alfreddobradi/color:v1.0.2) but it can be built using the Dockerfile included.

#### Kubernetes

Changed the cluster to use a Deployment instead of creating a POD. This has the added benefit of managing replicas and is preferred over manually creating PODs. As you can see on the deployment configuration, not much has changed. The templates now contain a `resources` field and I added labels
as at least one label is required if the deployment uses labels as selectors to ensure the desired number of replicas are running.

Another upgrade to the previous version is the introduction of a service. The service is responsible for exposing a port for external access.

To start the cluster (first delete the POD if it still exists):

```
% kubectl apply -f k8s/deployment/color-deploy.yml

% kubectl apply -f k8s/service/color-service.yml
```

To access the web service, get the exposed port and curl or go to your cluster's IP on the exposed port (10.100.159.100:31355 in our example):

```
% kubectl get svc color-service
NAME            TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
color-service   NodePort   10.100.159.100   <none>        8080:31355/TCP   34m
```

NOTE: If you use `minikube` on macOS (like myself), you won't be able to reach this endpoint. In order to check the output, you'll have to use

```
% minikube service color-service
```

This command prints the locally exposed access and also opens a browser pointing to the exposed port.