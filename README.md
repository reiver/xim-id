# xim-id

The program `xim-id` a **xim-id**

**xim-ids** are
quazi‐ monotonically‐increasing unique‐identifiers,
that are safe to use as a file or directory name,
and where lexical-ordering of the **xim-id** (under Unicode & ASCII) is in-practice also temporal-ordering of the xim-id (almost all of the time)

## Example xim-id

A **xim-id** looks something like this:
```
xi-556PVvNyq3m
```

## Usage №1
To use the program `xim-id`, run it as following:
```
xim-id
```

This will output a **xim-id**, such as:
```
xi-557ksIsNSRm
```
… with a trailing new-line at the end.

## Usage №2
To do almost exactly the same as that, but with no trailing new-line line, use the program `xim-id` as follows:
```
xim-id -n
```

## Usage №3
To take a **xim-id** (such as `xi-557ksIsNSRm`) and decompile it into its timestamp and its chaos, run the program `xim-id` as follows:
```
xim-id decompile xi-557ksIsNSRm
```

## xim-id

For more information on **xim-ids** see:

* https://github.com/reiver/go-xim
