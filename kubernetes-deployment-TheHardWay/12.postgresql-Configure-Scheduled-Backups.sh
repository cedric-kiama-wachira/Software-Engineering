cat <<EOF | kubectl apply -f -
apiVersion: batch/v1
kind: CronJob
metadata:
  name: postgresql-backup
  namespace: postgresql
spec:
  schedule: "0 1 * * *"  # Every day at 1 AM
  concurrencyPolicy: Forbid
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 3
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccountName: postgresql-sa
          containers:
            - name: pgbackrest
              image: cedrickiama/pg17-4-timescaledb:2.19.3-slim
              command: ["pgbackrest", "backup", "--type=full"]
              volumeMounts:
                - name: pgbackrest-config
                  mountPath: /etc/pgbackrest/pgbackrest.conf
                  subPath: pgbackrest.conf
                - name: postgresql-data
                  mountPath: /var/lib/postgresql/17/main
                - name: pgbackrest-data
                  mountPath: /var/lib/pgbackrest
              resources:
                requests:
                  cpu: "500m"
                  memory: "1Gi"
                limits:
                  cpu: "1"
                  memory: "2Gi"
          volumes:
            - name: pgbackrest-config
              configMap:
                name: pgbackrest-config
            - name: postgresql-data
              persistentVolumeClaim:
                claimName: postgresql-data-postgresql-0
            - name: pgbackrest-data
              persistentVolumeClaim:
                claimName: pgbackrest-data-postgresql-0
          restartPolicy: OnFailure
EOF
