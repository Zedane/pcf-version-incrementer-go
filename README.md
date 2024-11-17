# pcf-version-incrementer-go
A tool to set the version of a PCF-Component, written in Go.

Simple command line tool for incrementing the [PCF](https://learn.microsoft.com/en-us/power-apps/developer/component-framework/overview)-controls version number.
The tool sets the Control-Version number in the *manifest*, *solution* and *package.json* files, since this seemed to be working best for me.


## Usage/Examples

#### Status
Show the current version of your component.
```
pfcv status
```
Flags:
- `--all` (`-a`): Shows the version number specified in each of the three files.
- `--verbose` (`-v`): Show verbose output details.

#### Increment
Increments the minor version by one.
```
pfcv increment
```
Flags:
- `--major` (`-m`): Increment the major version instead.
- `--patch` (`-p`): Increment the patch version instead.
- `--verbose` (`-v`): Show verbose output details.


#### Set
Set a specific version. At least one must be provided
```
pfcv set --major 2 --minor 1 --patch 0
```
Flags:
- `--major`: Set the major version.
- `--minor`: Set the minor version.
- `--patch`: Set the patch version.
- `--verbose` (`-v`): Show verbose output details.

#### Help
Show help about the tool or about a specific command.
```
pfcv help [COMMAND]
```


## FAQ

#### Why though? Seems pointless...

For some reason the componentes I've beed working on didn't update when importing a solution/component with the same version. Since I'm too lazy to manually update the version in three places, I automated the 10sec task in just one Day! Seriously though, it's just nice not to worry about that.



## License

See [MIT License](LICENSE.md)