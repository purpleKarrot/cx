# cx

A unified front end for the CMake ecosystem.

## Summary

* User preferences are stored in `~/.config/cx/config.<ext>`. Supported formats include JSON,TOML, YAML.
* The project root is determined by looking for `CMakeLists.txt` files in parent directories and picking the one that is closest to the filesystem root (that is **not** the nearest parent).
* Build directory located at `~/.cache/cx/<md5 of project root>`.
* `cx configure` invokes CMake to create a build system.
* `cx build` from anywhere in the source project invokes `cx configure` if `CMakeCache.txt` does not exist and then builds the `all` target of the current directory. That may be `cmake --build <build_dir> --target <subdir>/all` or `cmake --build <build_dir>/<subdir>` depending on the generator.
* `cx test` from anywhere in the source project invokes `cx build` *unconditionally* and then executes tests that are defined in the current directory and below with `ctest --test-dir <build_dir>/<subdir>`.
* `cx install` from anywhere in the source project invokes `cx build` *unconditionally* and then installs targets that are defined in the current directory and below with `cmake --install <build_dir>/<subdir>`.
* `cx open` invokes `cx configure` if `CMakeCache.txt` does not exist and then opens the associated IDE of the generator if supported, like Xcode.
* `cx clean` deletes the build directory.

Options are still limited. Contributions are welcome.

