```
brew install gnu-sed
```

```
gsed 's/namespace: k8s-intro/namespace: k8s-intro-teamname/g' *
```

Go through and apply these different files one by one, seeing what each does and how it works. Ask Junaid/Apos/Gary if stuck)

    * Pod (just returns hostname)
        * Won’t restart when it fails (app exists, node fails or cycles)
        * Doesn’t scale (you’d have to write a new pod file with a different pod name every time you wanted to add another one)
        * Very limited management (no rollouts/rollbacks/scaling/self-healing)
    * Deployment-1 (still just returns hosthostname)
        * Show pod being recreated if we delete one
        * Show scaling up to 2 replicas
        * Show both accessible via curl (kubectl get pods -o wide to get IPs), explain why this isn’t ideal (you’d have to know the IPs to call from the client side)
    * Apply the service
        * Show DNS name (curl k8s-intro.teamname.svc.cluster.local)
        * Show how it balances across both pods via loop in bash pod (run multiple times, get different responses)
    * Deployment-2 + ConfigMaps
	* Make sure you have applied the ConfigMap
        * Show us providing configuration options via ConfigMap/envFrom
        * Show how we get different responses depending on ?envvar= provided
	* Update the ConfigMap with different values. Do a rolling restart of deployment. Show how curl returns old/new values as pods cycle in/out
    * Deployment-3 - Fails to run
        * Show how the rollout pauses
        * Describe pod
        * Look at logs
    * Deployment-4 - Working again - minReadySeconds 60
        * Show gradual rollout (kubectl rollout status deployment k8s-intro)
    * Ingress-1
        * Show how it passes traffic through (tell Gary the team name so he can create DNS records)
    * Deployment-5 + service-2 + Ingress-2
        * Different code
        * Exposed via a different service
        * Mapped via a different route on Ingress
            * /wibble goes to wobble service
            * Everything else goes to k8s-intro service

