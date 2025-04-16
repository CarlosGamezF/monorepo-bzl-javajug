load("@contrib_rules_jvm//java:defs.bzl", "java_test_suite")
load("@rules_spring//springboot:springboot.bzl", "springboot")

java_library(
    name = "java_service_lib",
    srcs = glob(["src/main/java/**/*.java"]),
    resources = glob(["**/src/main/resources/**"]),
    deps = [
        "@maven//:org_springframework_boot_spring_boot",
        "@maven//:org_springframework_boot_spring_boot_autoconfigure",
        "@maven//:org_springframework_boot_spring_boot_starter_web",
        "@maven//:org_springframework_spring_web",
    ],
    visibility = ["//visibility:public"],
)

springboot(
    name = "helloworld",
    boot_app_class = "com.example.App",
    java_library = ":java_service_lib",
)


java_test_suite(
    name = "tests",
    srcs = glob( ["src/test/java/**/*.java"]),
    runner = "junit5",
    package = "com.example",
    deps = [
        ":java_service_lib",
        "@maven//:org_junit_jupiter_junit_jupiter_api",
        "@maven//:org_junit_platform_junit_platform_launcher",
        "@maven//:org_junit_platform_junit_platform_reporting",
        "@maven//:org_springframework_boot_spring_boot_test",
        "@maven//:org_springframework_boot_spring_boot_starter_test",
        "@maven//:org_springframework_boot_spring_boot_test_autoconfigure",
        "@maven//:org_springframework_spring_test",
        "@maven//:org_springframework_spring_beans",
    ],
    visibility = ["//visibility:public"],
)
