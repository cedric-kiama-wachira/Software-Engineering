kubectl logs -n rook-ceph $(kubectl get pods -n rook-ceph -l app=csi-rbdplugin-provisioner -o name | head -n 1) -c csi-provisioner
