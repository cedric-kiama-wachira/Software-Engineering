# Check if the StatefulSet is running
kubectl get statefulset -n postgresql

# Check the PostgreSQL pod status
kubectl get pods -n postgresql

# Check the persistent volume claims
kubectl get pvc -n postgresql

# Check the persistent volumes
kubectl get pv | grep postgresql
