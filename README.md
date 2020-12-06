### Compile the Binary
You can run the `make build` command to get the `operator-mathew` binary. As follow:
```console
$ make build
mkdir -p "build/_output/bin/"
export GO111MODULE=on && export GOPROXY=https://goproxy.io && go build -o "build/_output/bin/operator-mathew" "./cmd/manager/main.go"

$ ls -l build/_output/bin/operator-mathew
-rwxr-xr-x. 1 root root 39674923 Oct 22 22:23 build/_output/bin/operator-mathew
$ ./build/_output/bin/operator-mathew --help
Usage of ./build/_output/bin/operator-mathew:
      --kubeconfig string                Paths to a kubeconfig. Only required if out-of-cluster.
      --master --kubeconfig              (Deprecated: switch to --kubeconfig) The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.
      --zap-devel                        Enable zap development mode (changes defaults to console encoder, debug log level, and disables sampling)
      --zap-encoder encoder              Zap log encoding ('json' or 'console')
      --zap-level level                  Zap log level (one of 'debug', 'info', 'error' or any integer value > 0) (default info)
      --zap-sample sample                Enable zap log sampling. Sampling will be disabled for integer log levels > 1
      --zap-time-encoding timeEncoding   Sets the zap time format ('epoch', 'millis', 'nano', or 'iso8601') (default )
pflag: help requested
```

### Create a Docker image for this binary
You can use this [Dockerfile](https://github.com/Mathew857/operator-mathew/blob/master/build/Dockerfile), as follows:
```console
$ docker build -f build/Dockerfile -t registry.cn-beijing.aliyuncs.com/mathew-cloud/operator-mathew:latest .
Sending build context to Docker daemon  40.27MB
Step 1/7 : FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
latest: Pulling from ubi8/ubi-minimal
0fd3b5213a9b: Already exists
aebb8c556853: Already exists
Digest: sha256:5cfbaf45ca96806917830c183e9f37df2e913b187aadb32e89fd83fa455ebaa6
Status: Downloaded newer image for registry.access.redhat.com/ubi8/ubi-minimal:latest
 ---> 28095021e526
Step 2/7 : ENV OPERATOR=/usr/local/bin/operator-mathew     USER_UID=1001     USER_NAME=operator-mathew
 ---> Running in 8ea785dc7d31
Removing intermediate container 8ea785dc7d31
 ---> e58d351d046f
Step 3/7 : COPY build/_output/bin/operator-mathew ${OPERATOR}
 ---> 24a495093ff2
Step 4/7 : COPY build/bin /usr/local/bin
 ---> b31ea373b410
Step 5/7 : RUN  /usr/local/bin/user_setup
 ---> Running in 203017c7df32
+ mkdir -p /root
+ chown 1001:0 /root
+ chmod ug+rwx /root
+ chmod g+rw /etc/passwd
+ rm /usr/local/bin/user_setup
Removing intermediate container 203017c7df32
 ---> 7217a26389de
Step 6/7 : ENTRYPOINT ["/usr/local/bin/entrypoint"]
 ---> Running in 4f2253ae0857
Removing intermediate container 4f2253ae0857
 ---> eca03f1431ac
Step 7/7 : USER ${USER_UID}
 ---> Running in 418b6dd2d390
Removing intermediate container 418b6dd2d390
 ---> 6d1c3e2b3445
Successfully built 6d1c3e2b3445
Successfully tagged registry.cn-beijing.aliyuncs.com/mathew-cloud/operator-mathew:latest
```

### Create the mathew CR
You will get an `example-mathew` instance and its two pods, as follows:
```console
# cat mathew-instance.yaml
apiVersion: app.mathew.com/latest
kind: mathew
metadata:
  name: example-mathew
  namespace: default
spec:
  size: 2

$ kubectl get  mathew -n default
NAME            AGE
example-mathew   34m
$ kubectl get  pods -n default
NAME                             READY   STATUS    RESTARTS   AGE
example-mathew-85fc47cf75-7fq2m   1/1     Running   0          34m
example-mathew-85fc47cf75-jv9mg   1/1     Running   0          34m
operator-mathew-9cd7b7d5c-ffxs9   1/1     Running   0          40m
```

