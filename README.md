`NoMangas` is a ClI for linux destiny to Download and Read Mangas (of the page LeerManga) on format `PDF`.
![MangaExample](https://github.com/MiguelVRRL/NoMangas/blob/main/images/Rudel.jpg)

## Installation
 ```
go install https://github.com/MiguelVRRL/NoMangas
 ```
## Examples of use

### Command Download
for downloading the mangas.
by default this has a limit of 10 mangas, but with the flag -limit (-l) you can remove this 

```
nomangas download "chainsawman" -l=true 1 2 3 4 5 6 7 8 9 10 11 12
```

### Command List
This command will to display the mangas downloaded. Usage:

```
nomangas
```

### Command Read
For reading your mangas using the CLI. usage
```
nomangas read [name of manga] [cap]
```
