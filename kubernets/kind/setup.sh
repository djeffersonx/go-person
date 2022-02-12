#!/bin/sh
set -o errexit

setup_metallb(){
  kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/namespace.yaml
  kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"
  kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/metallb.yaml
  kubectl get pods -n metallb-system --watch
  docker network inspect -f '{{.IPAM.Config}}' kind
  kubectl apply -f https://kind.sigs.k8s.io/examples/loadbalancer/metallb-configmap.yaml
}

create_local_docker_registry(){
  if [ "$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)" != 'true' ]; then
    docker run \
      -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
      registry:2
  fi
}

create_kind_cluster(){

  # create a cluster with the local registry enabled in containerd
  cat <<EOF | kind create cluster --config=-
  kind: Cluster
  apiVersion: kind.x-k8s.io/v1alpha4
  containerdConfigPatches:
  - |-
    [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
      endpoint = ["http://${reg_name}:5000"]
EOF

}

connect_kind_cluster_to_docker_registry(){
  # connect the registry to the cluster network if not already connected
  if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${reg_name}")" = 'null' ]; then
    docker network connect "kind" "${reg_name}"
  fi
}


apply_kind_local_registry_config_map(){
  # Document the local registry
  # https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
  cat <<EOF | kubectl apply -f -
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: local-registry-hosting
    namespace: kube-public
  data:
    localRegistryHosting.v1: |
      host: "localhost:${reg_port}"
      help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF
}

re_create_kind_cluster(){
  kind delete cluster
  create_kind_cluster
  kubectl cluster-info --context kind-kind
}

reg_name='kind-registry'
reg_port='5001'

re_create_kind_cluster
setup_metallb
create_local_docker_registry
connect_kind_cluster_to_docker_registry
apply_kind_local_registry_config_map
