load("@rules_python//python:defs.bzl", "py_binary", "py_test")

py_library(
    name = "python_service_lib",
    srcs = ["app.py"],
    deps = [
         "@pypi//flask",
    ],
)

py_binary(
  name = "python_service",
  srcs = ["app.py"],
  main = "app.py",
  deps = [
    ":python_service_lib"
  ]
)

py_test(
  name = "python_service_test",
  srcs = ["test_app.py"],
  main = "test_app.py",
  deps = [
    ":python_service",
  ],
  imports = ["."],
)