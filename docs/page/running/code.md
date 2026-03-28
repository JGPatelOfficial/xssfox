---
title: In the Code
redirect_from: /docs/code/
parent: Running
nav_order: 4
toc: true
layout: page
---

# Using XSSFox in Your Code

This guide provides detailed instructions on how to use XSSFox as a library in your Go projects. Follow the steps below to integrate XSSFox into your code.

## Get the XSSFox Library

First, you need to download the XSSFox library using the `go get` command:

```bash
go get github.com/JGPatelOfficial/xssfox/lib
```

## Sample Code

Here is a sample Go program that demonstrates how to use the XSSFox library to perform a scan:

```go
package main

import (
    "fmt"

    xssfox "github.com/JGPatelOfficial/xssfox/lib"
)

func main() {
    // Set up options for the scan
    opt := xssfox.Options{
        Cookie: "ABCD=1234",
    }

    // Create a new scan target
    target := xssfox.Target{
        URL:     "https://xss-game.appspot.com/level1/frame",
        Method:  "GET",
        Options: opt,
    }

    // Perform the scan
    result, err := xssfox.NewScan(target)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Scan Result:", result)
    }
}
```

## Running the Code

To run the sample code, follow these steps:

### Initialize Your Project

First, initialize your Go module:

```bash
go mod init <YOUR_PROJECT_REPO>
```

Replace `<YOUR_PROJECT_REPO>` with the path to your project repository.

### Build the Application

Next, build your application:

```bash
go build -o testapp
```

During the build process, Go will download the XSSFox library and its dependencies.

### Run the Application

Finally, run your application:

```bash
./testapp
```

You should see output similar to the following:

```bash
# [] [{V GET https://xss-game.appspot.com/level1/frame?query=%3Ciframe+srcdoc%3D%22%3Cinput+onauxclick%3Dprint%281%29%3E%22+class%3Dxssfox%3E%3C%2Fiframe%3E}] 2.618998247s 2021-07-11 10:59:26.508483153 +0900 KST m=+0.000794230 2021-07-11 10:59:29.127481217 +0900 KST m=+2.619792477
```

## More Information

For more information and advanced usage, please refer to the [official XSSFox library documentation](https://pkg.go.dev/github.com/JGPatelOfficial/xssfox).
