# forge-template

Tired of repetitive, boilerplate code? `forge-template` is here to rescue you! No more wasting time on generic, repetitive tasks! With `forge-template` you can quickly generate templates tailored to your project requirements, letting you focus on what really matters, your code.

## Table of Contents

- [forge-template](#forge-template)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Go tool](#go-tool)
    - [Unix/Linux](#unixlinux)
    - [Windows](#windows)
  - [Basic Syntax](#basic-syntax)
    - [Defining Variable Values](#defining-variable-values)
    - [Escaping Variable Values](#escaping-variable-values)
  - [Defining a file template](#defining-a-file-template)
    - [Defining a file template Example](#defining-a-file-template-example)
  - [Defining a project template](#defining-a-project-template)
    - [Project Template Syntax](#project-template-syntax)
  - [Defining a project template Example](#defining-a-project-template-example)

## Installation

### Go tool

If you have Go installed on your system, you can use the `go install` command to install `forge-template`. Run the following command in your terminal:

```bash
go install github.com/hanskrohn/forge-template@latest
```

### Unix/Linux

1. Navigate to the [Releases](https://github.com/hanskrohn/forge-template/releases) page.
2. Download the appropriate file for your Unix/Linux system.
3. After downloading the file open a terminal and run the following command to rename the download

```bash
mv ~/PATH_TO_DOWNLOAD/downloaded_file ~/PATH_TO_DOWNLOAD/forge-template.
```

4. Move the renamed file to a directory in your PATH. It can be done with the command below

```bash
mv ~/PATH_TO_DOWNLOAD/forge-template ~/DESIRED_LOCATION_PATH
```

If you are unsure how to add a directory to your path you can follow the commands below

```bash
mkdir ~/bin
echo 'export PATH="$PATH:$HOME/bin"' >> ~/.bashrc  # For bash
echo 'export PATH="$PATH:$HOME/bin"' >> ~/.zshrc   # For zsh
mv ~/PATH_TO_DOWNLOAD/forge-template ~/bin
```

### Windows

1. Navigate to the [Releases](https://github.com/hanskrohn/forge-template/releases) page.
2. Download the appropriate file for your Windows system.
3. After downloading the file open a Command Prompt in administrator mode and run the following command to rename the download

```shell
rename C:\PATH_TO_DOWNLOAD\downloaded_file forge-template
```

4. Move the renamed file to a directory in your PATH. It can be done with the command below

```shell
move C:\PATH_TO_DOWNLOAD\forge-template C:\DESIRED_LOCATION_PATH
```

If you are unsure how to add a directory to your path you can follow the commands below

```shell
mkdir C:\bin
setx PATH "%PATH%;C:\bin"
move C:\PATH_TO_DOWNLOAD\forge-template C:\bin
```

## Basic Syntax

### Defining Variable Values

To define variable values in your template simply wrap the variable with `<>`, like this `<<VARIABLE_NAME>>`. When creating the template we will reference the the `VARIABLE_NAME` and prompt you to provide values.

### Escaping Variable Values

In the case your code requires the symbols `<>` wrapped around a variable, please use the `\` to escape the variable.

Example:

```tsx
import React from 'react';

interface <<ComponentName>>Props {}

// Only need 1 escape over here         *
export const <<ComponentName>>: React.FC\<<<ComponentName>>Props> = ({ }) {
  return <>Hello From <<ComponentName>></>;
}
```

## Defining a file template

To define a file template run `forge-template` and select the `Create File Template` option. This will open up tool tips that prompt you for the file name and template code. Once this data is saved, you can run `forge-template` again and select `Create File From Template` option. This will open a tool tip to select the file template you want to create, and prompt you to provide variables if necessary.

### Defining a file template Example

1. Run `forge-template` and select `Create File Template`
2. Provide the name `controller.tsx` in the input box and press `Ctrl+s` to save
3. Provide the following template code in the text area and press `Ctrl+s` to save

```tsx
import React from 'react';

interface <<ComponentName>>Props {}

export const <<ComponentName>>: React.FC\<<<ComponentName>>Props> = ({ }) {
  return <>Hello From <<ComponentName>></>;
}
```

4. Run `forge-template` again and select `Create File From Template`
5. Use the arrow keys to select `controller.tsx`
6. Provide the value `User` when prompted to provide a value for `ComponentName`. Press `Ctrl+s` to save
7. View the newly generated file with your favorite text editor

## Defining a project template

To define a project template run `forge-template` and select the `Create Project Template` option. This will open up tool tips that prompt you for the project name and template code. Once this data is saved, you can run `forge-template` again and select `Create Project From Template` option. This will open a tool tip to select the file template you want to create, and prompt you to provide variables if necessary.

### Project Template Syntax

- Use `<<>>` to define variable values
- The root element should have no leading symbols or spaces
- Use `.` as the root element to define the project structure to build out from current directory
- Child directories and files should have 3 leading spaces. If you have more than 3 you will get an error and less than three will not result in a child element being created
- Directories are prefaced with `|--`
- Files are prefaced with `|-`

## Defining a project template Example

1. Run `forge-template` and select `Create Project Template`
2. Provide the name `basic-structure` in the input box and press `Ctrl+s` to save
3. Provide the following template code in the text area and press `Ctrl+s` to save

```txt
basic-project
   |-- src
      |- index.js
   |-- test
      |-- index.test.js
```

4. Run `forge-template` again and select `Create Project From Template`
5. Use the arrow keys to select `basic-structure`
6. View the newly generated file with your favorite text editor
