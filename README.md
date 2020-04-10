# EnvDef

envdef provides methods of setting environment value with default value.

## Installation

As a library

```
go get github.com/locona/envdef
```

or if you want to use it as a bin command

```
go get github.com/locona/envdef/cmd/envdef
```

## Usage

Add your application configuration to your `.env` file in the root of your project:

```
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
```

and `.env.sample` file in the root of your project:

```
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
REGION=REGION
```

And then run:
```
envdef
```

As a result created the `.env.new` file

```
REGION=REGION
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
```
