# go_service/BUILD
load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_service_lib",
    srcs = ["main.go"],
    importpath = "example.com/go_service",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "go_service",
    embed = [":go_service_lib"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_service_test",
    srcs = ["main_test.go"],
    embed = [":go_service_lib"],
)