# saasc - An Azure App Service CLI tool to manage your app services.

This is an unofficial open source terminal CLI tool to manage your azure app services.

`saasc` is free to use and currently at its `BETA` stage. We are continuously working on adding more features to saasc, so stay tune and keep using `saasc`.

## Screens:
Coming Soon...

## Installations:
`saasc` can be installed through shell script or `make` through source code. 

#### Install By Make:
1. Clone this git repository.
2. Run `make install`
3. Run `make build`.
4. Make sure to set all the configurations inside `config.toml`
5. Run `saasc` to start the ntui.

#### Install By shell script:
1. Clone this git repository.
2. Run setup script using `bash ./setup.sh`. It will setup a home directory(`.saasc`) with configs. Make sure to set all the configurations inside `config.toml`.
3. Run build script using `bash ./build.sh`. It will build the code in local system.
4. Run it using `./bin/saasc`. More options can be viewed using `./bin/saasc --help`  

## How To Use It:

`saasc` requires some configurations to be set, the default config file should be at user's root home diretory.

Default config file looks like below - 

```toml
  #Home Directory of saasc
  Home_Dir = ""

  #Set the config path from where all configs should be read
  #App_Dir = ""

  #Set log levels, default: info (info|error|warn) 
  #Log_Level = "info"

  #Set custom log directory for storing ntui app logs
  #Log_Dir = ""

  #Set Config file path
  Config_Path = ""

  #Set Log filepath
  #Log_File_Path = ""
  
  #Set ntui screen table data refresh rate, default: 5
  #Refresh_Rate = 3

  #Set Azure Subscription Id
  #Azure_Subscription_Id = ""

  #Set Azure Client Id, which should have access to read your app service
  #Azure_Client_Id = ""

  #Set Azure Client Secret of your azure client id
  #Azure_Client_Secret = ""

  #Set Azure Tenant Id, which has your subscription
  #Azure_Tenant_Id = ""
```

### Commands:
```shell
# Run ntui
ntui

# View Help options
ntui help

#  View current ntui version
ntui version

# View config values.
ntui config 
```
### Flags:

Below are the falgs which can be passed while running ntui - 

`-c or --config-path` to set ntui config path.

`--home-dir` to  set home directory of ntui app.

`--host` to set nomad host.

`-l or --log-level` to set the ntui log level.

`--region` to set the nomad region.

`-n or --namespace` to set the nomad namespace.

`-r or --refresh` to set refresh rate to refresh the screen data.

`--skip-verify` to set if skip cetificate verification.

`-t or --token` to set nomad token to perform actions, which requires it.

#### Keys:

##### Global Keys
`<1>`: To view Nomad Nodes

`<2>`: To view/change regions and namespaces 

`<esc>`: To go back to previous screen

`<enter>`: To select the row

##### Jobs Screen
`<ctrl+q>`: To stop job

`<ctrl+s>`: To start job

`<d>`: View job's definition

`<ctrl+d>`: Run new job

##### TaskGroups Screen
`<v>`: To view job versions

##### Versions Screen
`<ctrl+v>`: To revert the selected job versions

##### Allocations Screen
`<ctrl+t>`: To restart selected task.

`<l>`: View logs of selected alloc.

##### Tasks Screen
`<ctrl+t>`: To restart selected task.

`<l>`: To view logs of selected task.

##### Logs Screen
`<e>`: To view stderr logs.

`<o>`: To view stdout logs.

