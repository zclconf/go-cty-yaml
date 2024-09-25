# Unreleased

* The YAML decoder now exactly follows the [YAML specification](https://yaml.org/spec/1.2-old/spec.html#id2805071) when parsing integers.
  **Note:** This removes support for `_` spaces as well as signed base8 and base16 values.

# 1.0.3 (November 2, 2022)

* The `YAMLDecodeFunc` cty function now correctly handles both entirely empty
  documents and explicit top-level nulls. Previously it would always return
  an unknown value in those cases; it now returns a null value as intended.
  ([#7](https://github.com/zclconf/go-cty-yaml/pull/7))

# 1.0.2 (June 17, 2020)

* The YAML decoder now follows the YAML specification more closely when parsing
  numeric values.
  ([#6](https://github.com/zclconf/go-cty-yaml/pull/6))

# 1.0.1 (July 30, 2019)

* The YAML decoder is now correctly treating quoted scalars as verbatim literal
  strings rather than using the fuzzy type selection rules for them. Fuzzy
  type selection rules still apply to unquoted scalars.
  ([#4](https://github.com/zclconf/go-cty-yaml/pull/4))

# 1.0.0 (May 26, 2019)

Initial release.
