# namegen

The namegen package will generate a filename with a UUIDv4-like unique identifier with an added suffix on a standard UNIX timestamp in HEX format.

It will generate filenames that looks like `00000000-0000-0000-0000-000000000000-00000000`.


## Functions

### Generate

```text
Generate generates an unique String UUIDv4-like + Current Unix timestamp in HEX.
It returns the generated String and an any errors.
```