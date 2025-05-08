kubectl delete deployment -n postgresql postgresql-custom
kubectl delete statefulset -n postgresql postgresql --cascade=orphan

kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: postgresql-debug
  namespace: postgresql
spec:
  containers:
  - name: postgresql
    image: cedrickiama/pg17-4-timescaledb:2.19.3-slim
    command: ["bash", "-c", "sleep 3600"]
    volumeMounts:
    - name: data-volume
      mountPath: /var/lib/postgresql/17/main
  volumes:
  - name: data-volume
    emptyDir: {}
EOF
