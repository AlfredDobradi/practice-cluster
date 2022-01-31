# practice-cluster

I created this repository for my own practicing purposes to build a simple app with various dependencies to be ran inside a Kubernetes cluster.

The tags will progressively introduce more and more concepts like deployments, services, configuration management, etc.

### Current tag: v1.0.2

#### The application

Simply takes the very first element from a string slice of colors and returns it as a JSON along the version.
If any errors occur during marshaling the data, it returns the error instead of the color.

Docker image is pushed to Docker HUB (alfreddobradi/color:v1.0.2) but it can be built using the Dockerfile included.

#### Kubernetes

Single POD definition exposing the 8080 containerPort. In order to deploy and access just create the POD using the YAML definition and run a curl inside the container:

```
% kubectl apply -f color-pod.yml

% kubectl exec color-app -- curl -s localhost:8080
```

NOTE: We could expose the port and access it directly but this is not a very lifelike scenario as PODs should be managed by controllers.
