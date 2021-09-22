# Join-matic

An easy way to join data.

## Manual

- **file-path** _string_

    Data file path to append

- **output-file** _string_

    Where the data will be wrote. \nBy default this will be displayed on Stdout to send data througth >> bash statement (default "StdOut")

- **separ** _string_

    What will separate all data (default ",")

## Examples

`join-matic  --file-path "/path/to/file/data.csv" --output-file /path/to/output.txt --separ /`

## Considerations

Currently this just support a simple data column. Example:

```text
45
132
654
hi
join
me
12
12
```
