# Changelog

All notable changes to this project are documented here, newest first.

Entries are generated from [Conventional Commits](https://www.conventionalcommits.org)
and grouped by change type. This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.5.0] - 2026-07-19

### Features

- Add RenderString for text output by [@iberflow](https://github.com/iberflow) in [5cb9f2e](https://github.com/toaweme/sintax/commit/5cb9f2e7047bbcff3fcc18bf0f9764d869763651).
- Line_numbers starts from 1 and accepts start as param by [@iberflow](https://github.com/iberflow) in [#3](https://github.com/toaweme/sintax/pull/3).

### Documentation

- Word modifier comments to read without the signature by [@iberflow](https://github.com/iberflow) in [111c846](https://github.com/toaweme/sintax/commit/111c846ddfa990cc3108fc0638cd3b87cbd8112a).

## [0.4.0] - 2026-07-17

### Features

- Report missing data as a catchable miss by [@iberflow](https://github.com/iberflow) in [5818318](https://github.com/toaweme/sintax/commit/58183181590ce769a76bf965666cd090ed4e99b6).
- Name the failing modifier with a typed ModifierError by [@iberflow](https://github.com/iberflow) in [22d4601](https://github.com/toaweme/sintax/commit/22d460147fd486bd7b7a8f7800d2ae76659759c7).
- **[breaking]** Functional options for engine construction by [@iberflow](https://github.com/iberflow) in [9c194c8](https://github.com/toaweme/sintax/commit/9c194c89eac299dc4296887ea9330f393f5863bd).
- Allow scalars in text modifiers by [@iberflow](https://github.com/iberflow) in [3a2298a](https://github.com/toaweme/sintax/commit/3a2298a3a543cfec81659292b3cd5a1d4e0a2673).
- **[breaking]** Typed modifier clauses via Wrap/Overload adapters by [@iberflow](https://github.com/iberflow) in [9a59a86](https://github.com/toaweme/sintax/commit/9a59a8693cb36fd485a86d61d9a8ab490b2436a3).
- **[breaking]** Replace BuiltinFunctions with defaults.New and per-group Modifiers() by [@iberflow](https://github.com/iberflow) in [e1a4cd8](https://github.com/toaweme/sintax/commit/e1a4cd81bfacb4721171d67322d96dc5f8a60431).
- Html,js,url escape modifiers and - to _ modifier naming unification by [@iberflow](https://github.com/iberflow) in [d8f1217](https://github.com/toaweme/sintax/commit/d8f1217eceb465ea9cce15381d4be806e8240423).

### Fixes

- Let default catch empty-collection misses by [@iberflow](https://github.com/iberflow) in [e1e6599](https://github.com/toaweme/sintax/commit/e1e6599cf2099bd68c50a696f0bad7086ee8951d).
- Pin care action by commit sha, not tag object by [@iberflow](https://github.com/iberflow) in [1c2f9d1](https://github.com/toaweme/sintax/commit/1c2f9d103e2eddbf3596ca66973f8b8308c265a6).

### Documentation

- Expand runnable examples and clean modifier doc comments by [@iberflow](https://github.com/iberflow) in [770d271](https://github.com/toaweme/sintax/commit/770d271c87a2f389cc36211417b9ee2423f44800).
- Add runnable Go Example functions for all modifiers by [@iberflow](https://github.com/iberflow) in [1c10f81](https://github.com/toaweme/sintax/commit/1c10f81c64048d37750b60f14bac85185c5d29e3).
- Add Contributing section to README by [@iberflow](https://github.com/iberflow) in [ff2f5ed](https://github.com/toaweme/sintax/commit/ff2f5ed250b42fc70a437799f4aab233ff2094a6).

### Refactors

- Drop the unused Type constants by [@iberflow](https://github.com/iberflow) in [9fbfcf2](https://github.com/toaweme/sintax/commit/9fbfcf2f1e3e2d7f13a9ea299f3e2d8094a30574).
- Move every modifier onto the example-func doc convention by [@iberflow](https://github.com/iberflow) in [7da8dca](https://github.com/toaweme/sintax/commit/7da8dca735b014d9cfd4145abe93ebc3bfb75c25).
- Reorganise modifiers into subdirs by [@iberflow](https://github.com/iberflow) in [86d18c0](https://github.com/toaweme/sintax/commit/86d18c0ec76909f408d3a5db0632aeb911bb0d10).
- Reorganise modifiers a bit by [@iberflow](https://github.com/iberflow) in [b1836d0](https://github.com/toaweme/sintax/commit/b1836d05ef439107a21973733aa27729382b155c).

### Tests

- Name the e2e root test files consistently by [@iberflow](https://github.com/iberflow) in [7bc91c4](https://github.com/toaweme/sintax/commit/7bc91c49b5b4b20592d8c5850c2d2fb9eb0019a0).
- Compare baseline and adapter benchmarks fairly by [@iberflow](https://github.com/iberflow) in [bbf6aff](https://github.com/toaweme/sintax/commit/bbf6aff0e3f6427555269b70ae0b9585794543ca).
- Benchmark AsText and engine construction by [@iberflow](https://github.com/iberflow) in [25d80ee](https://github.com/toaweme/sintax/commit/25d80ee575cfc15a55e90696e5fea61ffc381cdb).

### CI & Build

- Bump care action to v0.9.3 by [@iberflow](https://github.com/iberflow) in [4f0853f](https://github.com/toaweme/sintax/commit/4f0853fdc4165368ea3bc5635bfb8b34b5c65522).
- Move action version to inline comment so dependabot maintains it by [@iberflow](https://github.com/iberflow) in [a326bae](https://github.com/toaweme/sintax/commit/a326bae21f01624e8f3be578f86d3e4c689cf3d2).
- Drop gomod dependabot block from dependency-free module by [@iberflow](https://github.com/iberflow) in [0e7e0a0](https://github.com/toaweme/sintax/commit/0e7e0a00d41afea2d507fbfd7203835aa5a1df41).
- Add governance workflows and contributor docs by [@iberflow](https://github.com/iberflow) in [1b08610](https://github.com/toaweme/sintax/commit/1b08610f509955425810df7c449be2fd9728b73a).

### Chores & Other

- Care lint fixes by [@iberflow](https://github.com/iberflow) in [0eafaba](https://github.com/toaweme/sintax/commit/0eafaba400224f91a286d46be2e040a73c32852b).
- Wip by [@iberflow](https://github.com/iberflow) in [cc4cb9b](https://github.com/toaweme/sintax/commit/cc4cb9b095f78a515e7feb86c2767ebdd1b91528).

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

[0.5.0]: https://github.com/toaweme/sintax/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/toaweme/sintax/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/toaweme/sintax/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/toaweme/sintax/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/toaweme/sintax/releases/tag/v0.1.0
