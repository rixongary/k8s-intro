# Introduction to Kubernetes

## Pre-requisites:

kubens:

`brew install kubectx`

## Introduction

The demo will build up gradually from a single pod, through some intermediate steps, to a more complex, internet accessible app.

For the duration of the demo, your team will have its own Kubernetes *namespace*. You can create one like so:

`kubectl create namespace k8s-intro-yourteamnamehere`

All pods created during the dempo expose a web server on port 8080. When queries, this webserver will return the *hostname* of the pod is is running in. For example:

```
curl 10.72.4.6:8080/

k8s-intro-6c9c49cc58-f5b7c

```

To create a pod to run your commands from, and run a shell inside it, you can run the following command:

`kubectl run debug -it --image=ubuntu:focal -- /bin/bash`

This should give you a command line, like so:

```
If you don't see a command prompt, try pressing enter.
root@debug:/# 
```

and from within this pod you can now install `curl` which you'll use to interact with your pods:

`apt update && apt install -y curl`

**Ask Junaid/Apos/Gary if stuck with the above**


## Creating the demo pods ##

On your laptop, run:
```
brew install gnu-sed
```

```
cd k8s/
gsed 's/namespace: k8s-intro/namespace: k8s-intro-yourteamnamehere/g' *
```

**Please check that the files have been updated with your custom namespace, otherwise you will all overwrite each others resources:**

`grep namespace *.yaml` - these should all have your custom team name! If not, do not proceed!

## Build your app!

Work through the following steps one by one, and apply the relevant files from the `k8s` directory, seeing what happens at each stage, and understanding how it works.

To apply a file to the cluster, pass the filename to the `kubectl` command. For example

```
kubectl apply -f pod.yaml
pod/k8s-intro created
```


## The manfiests

* pod.yaml (just returns hostname when `curl`ed)
    
    * What happens when you delete the pod? Does it restart? (`kubectl delete pod k8s-intro`)
    
* deployment-1.yaml
    * Now what happens if you delete a pod?
    * Try scaling the deployment up to more replicas `kubectl scale deployment k8s-intro --replicas 2`
    * Verify you now have multiple pods `kubectl get pods`
    * Get the IPs of both pods: `kubectl get pods -o wide`
    * `curl` both IPs and see what happens

* service.yaml
    * Your pods are now available via DNS (`curl k8s-intro.k8s-intro-teamname.svc.cluster.local`)
    * If you make multiple requests, what do you notice about the responses? What is happening behind the scenes?
* deployment-2 & configmap.yaml
    * Watch what happens to the pods after you apply this. What is happening? Run `kubectl get pods` multiple times
    * From your debug pod: `curl k8s-intro.k8s-intro-teamname.svc.cluster.local/?envvar=NAME`
    * what happens if you swap out `NAME` for `ROLE` or `COOL`?
    * Update the `ConfigMap` with some custom values by editing the file and applying it again.
    *  Do a rolling restart of deployment: `kubectl rollout restart deployment k8s-intro`
    * Whilst the pods are restarting, run lots of `curl`s from your debug pod. What happens?
* deployment-3.yaml
    * What's wrong with this deployment? 
    * How can you get more info? (hint: `kubectl describe`, `kubectl logs`)
    * Try rolling back to the previous working version: `kubectl rollout undo deployment k8s-intro`
    * What happens?
* deployment-4.yaml
    * Kubernetes can report how a rollout is going: `kubectl rollout status deployment k8s-intro`
* ingress-1.yaml (*make sure you've told Gary your teamname!*)
    * You'll need to update the hostname first to match your teamname on **line 25**
    * Get the list of ingresses: `kubectl get ingress`
    * Describe your Ingress to learn more about it: `kubectl describe ingress k8s-intro`
    * Visit your "app" in a browser! (on the VPN for anyone remote)
* deployment-5.yaml + service-2.yaml + ingress-2.yaml
    * What's different about this seconds Ingress file vs the first one?
    * Do we have multiple different deployments now? (`kubectl get deployment`)
    * Do we have more than one service? (`kubectl get service`)
    * What `/path` can visit to get a different response to the other pods? 
