# Changelog

All notable changes to this project are documented here, newest first.

Entries are generated from [Conventional Commits](https://www.conventionalcommits.org)
and grouped by change type. This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.3.0] - 2026-07-03

### Features

- Support [] and {} empty-collection literals as modifier args by [@iberflow](https://github.com/iberflow) in [f8a92dc](https://github.com/toaweme/sintax/commit/f8a92dcd7c9fe80e2e338ffce11fddd48cdcc1da).

### Documentation

- Fix inaccurate README examples and clarify key-modifier access by [@iberflow](https://github.com/iberflow) in [3c25fbb](https://github.com/toaweme/sintax/commit/3c25fbb41b5606d0de351d0013a0fd0ba26f0cdd).

## [0.2.0] - 2026-07-02

### Chores & Other

- Relicense to Apache 2.0 by [@iberflow](https://github.com/iberflow) in [84f5446](https://github.com/toaweme/sintax/commit/84f54466b3c4917a2e9c6b9927ab66e8b3bda860).

## [0.1.0] - 2026-07-01

### Features

- File read modifier with dir access protection by [@iberflow](https://github.com/iberflow) in [cc86d4f](https://github.com/toaweme/sintax/commit/cc86d4f55de12ccaf1d6df929a39c8ceabc260d8).
- Move date dependency into internal/ by [@iberflow](https://github.com/iberflow) in [a600b23](https://github.com/toaweme/sintax/commit/a600b23dcf10e2e5b19b83b2d79f8fd0d50ba38d).

### Fixes

- **Perf:** Regex compile by [@iberflow](https://github.com/iberflow) in [c0b76d7](https://github.com/toaweme/sintax/commit/c0b76d7a386838fb05befcf5b89ca319dd540f1b).
- **Perf:** Faster loop, conditional, and expression rendering by [@iberflow](https://github.com/iberflow) in [ac85404](https://github.com/toaweme/sintax/commit/ac8540414ab1cb7febe641caaf18516b62ad9736).
- **Perf:** Pipeline cache, strconv int/bool, map-key sort fast-path + MB/s benchmark instrumentation by [@iberflow](https://github.com/iberflow) in [5592f00](https://github.com/toaweme/sintax/commit/5592f00bc72cc04e02c0ea9a744ef9e3ac189a5a).
- Correct typos and dupword false positive in comments by [@iberflow](https://github.com/iberflow) in [3ff2af8](https://github.com/toaweme/sintax/commit/3ff2af8c494912cc8e2cf4e83a902b489655a2dd).
- Fix goimports import grouping in format.go by [@iberflow](https://github.com/iberflow) in [a181bb9](https://github.com/toaweme/sintax/commit/a181bb9f779f9acbddb0ff1a0b393f71ace2413d).
- Add missing package and exported-symbol doc comments by [@iberflow](https://github.com/iberflow) in [f4232b1](https://github.com/toaweme/sintax/commit/f4232b15b7515e49912a5126f8a4719cd092269e).
- Replace fmt.Errorf with errors.New where no format verbs are used by [@iberflow](https://github.com/iberflow) in [5c2f175](https://github.com/toaweme/sintax/commit/5c2f17533bc7f6f85240810d27e1d627e2829c48).
- Avoid string concatenation in loop when extracting expert size digits by [@iberflow](https://github.com/iberflow) in [ad6a15b](https://github.com/toaweme/sintax/commit/ad6a15be9a23a916dc95a389369b4344e66406e8).
- Use Go 1.22 integer range in loops instead of C-style counters by [@iberflow](https://github.com/iberflow) in [ee89839](https://github.com/toaweme/sintax/commit/ee8983917952d7e5c956d6582d9a86f9dc5e00ac).
- Use reflect.Pointer instead of deprecated reflect.Ptr alias by [@iberflow](https://github.com/iberflow) in [a888fee](https://github.com/toaweme/sintax/commit/a888fee11b05d7b84e19e78ae386197c82c31f0d).
- Simplify single-case type switches and if-else-return patterns by [@iberflow](https://github.com/iberflow) in [43d89b9](https://github.com/toaweme/sintax/commit/43d89b9d1fb2b22c22edd3e569413f316e060d64).
- Apply staticcheck and gocritic mechanical rewrites by [@iberflow](https://github.com/iberflow) in [ce51b9c](https://github.com/toaweme/sintax/commit/ce51b9c0848cd8830987b5824704e2df3854842d).
- Drop unnecessary else after return in bool render branch by [@iberflow](https://github.com/iberflow) in [b2ac08c](https://github.com/toaweme/sintax/commit/b2ac08c2a18110f776080fd6511ecf36edae9c7b).

### Documentation

- Updated README by [@iberflow](https://github.com/iberflow) in [45323b2](https://github.com/toaweme/sintax/commit/45323b2fdf2692c897d106504fba40bc159758ba).
- Update CHANGELOG by [@iberflow](https://github.com/iberflow) in [73803bf](https://github.com/toaweme/sintax/commit/73803bf464dc56300697f6221b65db47345b4700).
- Update readme and changelog by [@iberflow](https://github.com/iberflow) in [45aa630](https://github.com/toaweme/sintax/commit/45aa6304df08faed70bc7060b3e4a78ecdad7c31).

### CI & Build

- Bump care to v0.7.1 by [@iberflow](https://github.com/iberflow) in [df0f1cf](https://github.com/toaweme/sintax/commit/df0f1cf77e9a7e54cde735da3d1d994d1d4a0be3).
- Bump care to v0.8.0 by [@iberflow](https://github.com/iberflow) in [77561d4](https://github.com/toaweme/sintax/commit/77561d463c7c510be6cff7dd461c24b710b28252).

### Chores & Other

- Initial commit :) by [@iberflow](https://github.com/iberflow) in [57b809e](https://github.com/toaweme/sintax/commit/57b809e15f4b493b59e7185c3b21e9f67e0d7004).
- Cleanup by [@iberflow](https://github.com/iberflow) in [f624b71](https://github.com/toaweme/sintax/commit/f624b71e51f3bed39171b1b08af3482ed4e87fac).
- Better tests + benchmarks by [@iberflow](https://github.com/iberflow) in [0e250dd](https://github.com/toaweme/sintax/commit/0e250ddafa74ee0918b27442424575e9533754d1).
- Tidy up readme by [@iberflow](https://github.com/iberflow) in [9273f2f](https://github.com/toaweme/sintax/commit/9273f2f0b893cca9030586a240e827b7e508b2c3).
- Add README chrome, CHANGELOG, and quality workflow by [@iberflow](https://github.com/iberflow) in [9e16541](https://github.com/toaweme/sintax/commit/9e165410feb6e1ea15b8b1f70a0e67f4cbfd5d5e).
- Fix linter issues by [@iberflow](https://github.com/iberflow) in [79c2a53](https://github.com/toaweme/sintax/commit/79c2a53df09c87782bc16832572312826bd2985e).
- Rollback go.mod to 1.22 by [@iberflow](https://github.com/iberflow) in [4be0329](https://github.com/toaweme/sintax/commit/4be0329a7a690ebcc5a80b9c4aa0689ea0f0c5d5).

[0.3.0]: https://github.com/toaweme/sintax/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/toaweme/sintax/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/toaweme/sintax/releases/tag/v0.1.0
