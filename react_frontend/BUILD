load("@aspect_rules_js//js:defs.bzl", "js_test")
load("@npm//:defs.bzl", "npm_link_all_packages")
load("@aspect_rules_webpack//webpack:defs.bzl", "webpack_bundle", "webpack_devserver")

npm_link_all_packages()

exports_files(
    [".swcrc"],
    visibility = ["//:__subpackages__"],
)

STATIC_ASSETS = glob(["public/**"])

webpack_bundle(
    name = "bundle",
    srcs = ["//react_frontend/src:transpile"] + STATIC_ASSETS,
    node_modules = ":node_modules",
    output_dir = True,
    webpack_config = ":webpack.bazel.config.js",
    deps = [
        ":node_modules/html-webpack-plugin",
        ":node_modules/react",
        ":node_modules/react-dom",
    ],
)

js_test(
    name = "build_smoke_test",
    timeout = "short",
    data = ["bundle"],
    entry_point = "build_smoke_test.js",
)

webpack_devserver(
    name = "dev_server",
    data = ["//react_frontend/src:transpile"] + STATIC_ASSETS + [
        ":node_modules/html-webpack-plugin",
        ":node_modules/react",
        ":node_modules/react-dom",
    ],
    entry_point = "//react_frontend/src:index.js",
    node_modules = ":node_modules",
    webpack_config = ":webpack.bazel.config.js",
)
