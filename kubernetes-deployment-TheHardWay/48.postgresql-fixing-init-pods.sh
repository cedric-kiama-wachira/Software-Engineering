# Let's delete all existing resources to start fresh
kubectl delete statefulset -n postgresql --all
kubectl delete pvc -n postgresql --all
kubectl delete pod -n postgresql --all --force --grace-period=0

# Create a simple deployment with emptyDir
kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres-simple
  namespace: postgresql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres-simple
  template:
    metadata:
      labels:
        app: postgres-simple
    spec:
      containers:
      - name: postgres
        image: postgres:14
        env:
        - name: POSTGRES_PASSWORD
          value: "postgres123"
        ports:
        - containerPort: 5432
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql/data
      volumes:
      - name: data
        emptyDir: {}
EOF
