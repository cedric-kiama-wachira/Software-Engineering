# Create Cilium Network Policy
cat <<EOF | kubectl apply -f -
apiVersion: cilium.io/v1
kind: CiliumNetworkPolicy
metadata:
  name: postgres-network-policy
  namespace: postgresql
spec:
  endpointSelector:
    matchLabels:
      app: postgresql
  ingress:
    - fromEndpoints:
        - matchLabels:
            io.kubernetes.pod.namespace: postgresql
      toPorts:
        - ports:
            - port: "5432"
              protocol: TCP
    - fromEndpoints:
        - matchLabels:
            io.kubernetes.pod.namespace: monitoring
      toPorts:
        - ports:
            - port: "9187"
              protocol: TCP
  egress:
    - toEndpoints:
        - matchLabels:
            k8s:io.kubernetes.pod.namespace: kube-system
            k8s:k8s-app: kube-dns
      toPorts:
        - ports:
            - port: "53"
              protocol: UDP
          rules:
            dns:
              - matchPattern: "*"
    - toEntities:
      - world
      toPorts:
        - ports:
            - port: "443"
              protocol: TCP
EOF

