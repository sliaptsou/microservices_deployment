provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

resource "helm_release" "example-chart-resource" {
  name = "example-chart"

  repository       = "../"
  chart            = "example-chart"
  namespace        = "example-ns"
  create_namespace = true

  set {
    name  = "service.type"
    value = "ClusterIP"
  }
}
