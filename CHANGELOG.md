# Changelog

All notable changes to this project are documented here, newest first.

Entries are generated from [Conventional Commits](https://www.conventionalcommits.org)
and grouped by change type. This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Features

- Move date dependency into internal/ by Ignas Bernotas in [a600b23](https://github.com/toaweme/sintax/commit/a600b23dcf10e2e5b19b83b2d79f8fd0d50ba38d).
- File read modifier with dir access protection by Ignas Bernotas in [cc86d4f](https://github.com/toaweme/sintax/commit/cc86d4f55de12ccaf1d6df929a39c8ceabc260d8).

### Fixes

- Drop unnecessary else after return in bool render branch by Ignas Bernotas in [b2ac08c](https://github.com/toaweme/sintax/commit/b2ac08c2a18110f776080fd6511ecf36edae9c7b).
- Apply staticcheck and gocritic mechanical rewrites by Ignas Bernotas in [ce51b9c](https://github.com/toaweme/sintax/commit/ce51b9c0848cd8830987b5824704e2df3854842d).
- Simplify single-case type switches and if-else-return patterns by Ignas Bernotas in [43d89b9](https://github.com/toaweme/sintax/commit/43d89b9d1fb2b22c22edd3e569413f316e060d64).
- Use reflect.Pointer instead of deprecated reflect.Ptr alias by Ignas Bernotas in [a888fee](https://github.com/toaweme/sintax/commit/a888fee11b05d7b84e19e78ae386197c82c31f0d).
- Use Go 1.22 integer range in loops instead of C-style counters by Ignas Bernotas in [ee89839](https://github.com/toaweme/sintax/commit/ee8983917952d7e5c956d6582d9a86f9dc5e00ac).
- Avoid string concatenation in loop when extracting expert size digits by Ignas Bernotas in [ad6a15b](https://github.com/toaweme/sintax/commit/ad6a15be9a23a916dc95a389369b4344e66406e8).
- Replace fmt.Errorf with errors.New where no format verbs are used by Ignas Bernotas in [5c2f175](https://github.com/toaweme/sintax/commit/5c2f17533bc7f6f85240810d27e1d627e2829c48).
- Add missing package and exported-symbol doc comments by Ignas Bernotas in [f4232b1](https://github.com/toaweme/sintax/commit/f4232b15b7515e49912a5126f8a4719cd092269e).
- Fix goimports import grouping in format.go by Ignas Bernotas in [a181bb9](https://github.com/toaweme/sintax/commit/a181bb9f779f9acbddb0ff1a0b393f71ace2413d).
- Correct typos and dupword false positive in comments by Ignas Bernotas in [3ff2af8](https://github.com/toaweme/sintax/commit/3ff2af8c494912cc8e2cf4e83a902b489655a2dd).
- **Perf:** Pipeline cache, strconv int/bool, map-key sort fast-path + MB/s benchmark instrumentation by Ignas Bernotas in [5592f00](https://github.com/toaweme/sintax/commit/5592f00bc72cc04e02c0ea9a744ef9e3ac189a5a).
- **Perf:** Faster loop, conditional, and expression rendering by Ignas Bernotas in [ac85404](https://github.com/toaweme/sintax/commit/ac8540414ab1cb7febe641caaf18516b62ad9736).
- **Perf:** Regex compile by Ignas Bernotas in [c0b76d7](https://github.com/toaweme/sintax/commit/c0b76d7a386838fb05befcf5b89ca319dd540f1b).

### Documentation

- Update CHANGELOG by Ignas Bernotas in [73803bf](https://github.com/toaweme/sintax/commit/73803bf464dc56300697f6221b65db47345b4700).
- Updated README by Ignas Bernotas in [45323b2](https://github.com/toaweme/sintax/commit/45323b2fdf2692c897d106504fba40bc159758ba).

### CI & Build

- Bump care to v0.7.1 by Ignas Bernotas in [df0f1cf](https://github.com/toaweme/sintax/commit/df0f1cf77e9a7e54cde735da3d1d994d1d4a0be3).

### Chores & Other

- Rollback go.mod to 1.22 by Ignas Bernotas in [4be0329](https://github.com/toaweme/sintax/commit/4be0329a7a690ebcc5a80b9c4aa0689ea0f0c5d5).
- Fix linter issues by Ignas Bernotas in [79c2a53](https://github.com/toaweme/sintax/commit/79c2a53df09c87782bc16832572312826bd2985e).
- Add README chrome, CHANGELOG, and quality workflow by Ignas Bernotas in [9e16541](https://github.com/toaweme/sintax/commit/9e165410feb6e1ea15b8b1f70a0e67f4cbfd5d5e).
- Tidy up readme by Ignas Bernotas in [9273f2f](https://github.com/toaweme/sintax/commit/9273f2f0b893cca9030586a240e827b7e508b2c3).
- Better tests + benchmarks by Ignas Bernotas in [0e250dd](https://github.com/toaweme/sintax/commit/0e250ddafa74ee0918b27442424575e9533754d1).
- Cleanup by Ignas Bernotas in [f624b71](https://github.com/toaweme/sintax/commit/f624b71e51f3bed39171b1b08af3482ed4e87fac).
- Initial commit :) by Ignas Bernotas in [57b809e](https://github.com/toaweme/sintax/commit/57b809e15f4b493b59e7185c3b21e9f67e0d7004).

## [0.1.0] - 2026-07-01

### Features

- File read modifier with dir access protection by Ignas Bernotas in [cc86d4f](https://github.com/toaweme/sintax/commit/cc86d4f55de12ccaf1d6df929a39c8ceabc260d8).

### Fixes

- Drop unnecessary else after return in bool render branch by Ignas Bernotas in [b2ac08c](https://github.com/toaweme/sintax/commit/b2ac08c2a18110f776080fd6511ecf36edae9c7b).
- Apply staticcheck and gocritic mechanical rewrites by Ignas Bernotas in [ce51b9c](https://github.com/toaweme/sintax/commit/ce51b9c0848cd8830987b5824704e2df3854842d).
- Simplify single-case type switches and if-else-return patterns by Ignas Bernotas in [43d89b9](https://github.com/toaweme/sintax/commit/43d89b9d1fb2b22c22edd3e569413f316e060d64).
- Use reflect.Pointer instead of deprecated reflect.Ptr alias by Ignas Bernotas in [a888fee](https://github.com/toaweme/sintax/commit/a888fee11b05d7b84e19e78ae386197c82c31f0d).
- Use Go 1.22 integer range in loops instead of C-style counters by Ignas Bernotas in [ee89839](https://github.com/toaweme/sintax/commit/ee8983917952d7e5c956d6582d9a86f9dc5e00ac).
- Avoid string concatenation in loop when extracting expert size digits by Ignas Bernotas in [ad6a15b](https://github.com/toaweme/sintax/commit/ad6a15be9a23a916dc95a389369b4344e66406e8).
- Replace fmt.Errorf with errors.New where no format verbs are used by Ignas Bernotas in [5c2f175](https://github.com/toaweme/sintax/commit/5c2f17533bc7f6f85240810d27e1d627e2829c48).
- Add missing package and exported-symbol doc comments by Ignas Bernotas in [f4232b1](https://github.com/toaweme/sintax/commit/f4232b15b7515e49912a5126f8a4719cd092269e).
- Fix goimports import grouping in format.go by Ignas Bernotas in [a181bb9](https://github.com/toaweme/sintax/commit/a181bb9f779f9acbddb0ff1a0b393f71ace2413d).
- Correct typos and dupword false positive in comments by Ignas Bernotas in [3ff2af8](https://github.com/toaweme/sintax/commit/3ff2af8c494912cc8e2cf4e83a902b489655a2dd).
- **Perf:** Pipeline cache, strconv int/bool, map-key sort fast-path + MB/s benchmark instrumentation by Ignas Bernotas in [5592f00](https://github.com/toaweme/sintax/commit/5592f00bc72cc04e02c0ea9a744ef9e3ac189a5a).
- **Perf:** Faster loop, conditional, and expression rendering by Ignas Bernotas in [ac85404](https://github.com/toaweme/sintax/commit/ac8540414ab1cb7febe641caaf18516b62ad9736).
- **Perf:** Regex compile by Ignas Bernotas in [c0b76d7](https://github.com/toaweme/sintax/commit/c0b76d7a386838fb05befcf5b89ca319dd540f1b).

### Documentation

- Updated README by Ignas Bernotas in [45323b2](https://github.com/toaweme/sintax/commit/45323b2fdf2692c897d106504fba40bc159758ba).

### CI & Build

- Bump care to v0.7.1 by Ignas Bernotas in [df0f1cf](https://github.com/toaweme/sintax/commit/df0f1cf77e9a7e54cde735da3d1d994d1d4a0be3).

### Chores & Other

- Rollback go.mod to 1.22 by Ignas Bernotas in [4be0329](https://github.com/toaweme/sintax/commit/4be0329a7a690ebcc5a80b9c4aa0689ea0f0c5d5).
- Fix linter issues by Ignas Bernotas in [79c2a53](https://github.com/toaweme/sintax/commit/79c2a53df09c87782bc16832572312826bd2985e).
- Add README chrome, CHANGELOG, and quality workflow by Ignas Bernotas in [9e16541](https://github.com/toaweme/sintax/commit/9e165410feb6e1ea15b8b1f70a0e67f4cbfd5d5e).
- Tidy up readme by Ignas Bernotas in [9273f2f](https://github.com/toaweme/sintax/commit/9273f2f0b893cca9030586a240e827b7e508b2c3).
- Better tests + benchmarks by Ignas Bernotas in [0e250dd](https://github.com/toaweme/sintax/commit/0e250ddafa74ee0918b27442424575e9533754d1).
- Cleanup by Ignas Bernotas in [f624b71](https://github.com/toaweme/sintax/commit/f624b71e51f3bed39171b1b08af3482ed4e87fac).
- Initial commit :) by Ignas Bernotas in [57b809e](https://github.com/toaweme/sintax/commit/57b809e15f4b493b59e7185c3b21e9f67e0d7004).
