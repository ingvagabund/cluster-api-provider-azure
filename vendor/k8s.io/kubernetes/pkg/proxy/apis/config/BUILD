package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "register.go",
        "types.go",
        "zz_generated.deepcopy.go",
    ],
    importmap = "sigs.k8s.io/cluster-api-provider-azure/vendor/k8s.io/kubernetes/pkg/proxy/apis/config",
    importpath = "k8s.io/kubernetes/pkg/proxy/apis/config",
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/component-base/config:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//pkg/proxy/apis/config/fuzzer:all-srcs",
        "//pkg/proxy/apis/config/scheme:all-srcs",
        "//pkg/proxy/apis/config/v1alpha1:all-srcs",
        "//pkg/proxy/apis/config/validation:all-srcs",
    ],
    tags = ["automanaged"],
)
