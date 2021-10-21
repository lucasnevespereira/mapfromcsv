# MapFromCSV

Script to map csv data to a method with args.


## Usage

Set your filenames
```
const (
	csvFileName = "example.csv"
	outFileName = "output.txt"
)
```

Run `make output`

Your `output.txt` is generated 👍🏻

#### Example CSV

```
001,Function,true,example text
002,FunctionName,false,90
003,Name with Space,10,a simple Text
004,Name/with/Slash,78,true
```
#### Example Output

```
FUNCTION("001","Function","true","example text"),
FUNCTIONNAME("002","FunctionName","false","90"),
NAME_WITH_SPACE("003","Name with Space","10","a simple Text"),
NAME_WITH_SLASH("004","Name/with/Slash","78","true"),
```

<hr>

#### Additional info
Using argument2 as func name but you can change this here.
Also you can add more string replacements, here I handle spaces and "/".
```
	funcName := strings.ToUpper(data.arg2)
	funcName = strings.Replace(funcName, "/", "_", -1)
	funcName = strings.Replace(funcName, " ", "_", -1)
```