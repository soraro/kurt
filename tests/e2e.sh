#!/bin/bash
set -eou pipefail

kubectl apply -f tests/test.yml --wait
kubectl wait --for=condition=Ready pod/nginx -n test1
kubectl wait --for=condition=Ready pod/apache -n test2

# Generate some pod restarts
kubectl exec nginx -n test1 -- bash -c "kill 1"

kubectl exec apache -n test2 -- bash -c "kill 1"

sleep 5
kubectl exec nginx -n test1 -- bash -c "kill 1"

echo "[!] wait for pods to finish restarting..."
sleep 30

NGINX_RESTARTS=$(./kurt pods -n test1 -o json | jq '.pods[0].count')
if [ $NGINX_RESTARTS -eq 2 ]; then
    echo "[+] Correct number of restarts for nginx 👍"
else
    echo "[!] Incorrect number of restarts for nginx: $NGINX_RESTARTS"
    exit 1
fi

APACHE_RESTARTS=$(./kurt pods -n test2 -o json | jq '.pods[0].count')
if [ $APACHE_RESTARTS -eq 1 ]; then
    echo "[+] Correct number of restarts for apache 👍"
else
    echo "[!] Incorrect number of restarts for apache: $APACHE_RESTARTS"
    exit 1
fi

echo "[+] Cleaning up..."
kubectl delete -f tests/test.yml