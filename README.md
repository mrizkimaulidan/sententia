# sententia

OCR (Optical Character Recognition) is a technology that recognizes text within a digital image. It is commonly used to recognize text in scanned documents and images.

Thanks gosseract library.

- https://github.com/otiai10/gosseract

You need to install some of dependencies to running this program.

Recommend to use Linux or WSL. Not yet tested on Windows. In Windows you need to compile all of the dependencies, it is complicated, so for the sake of simplicity this program should be only working on Linux or WSL.

You need install g++ and other dependencies:
```bash
$ sudo apt install g++
```

```bash
$ sudo apt install libtesseract-dev
```

```bash
$ sudo apt install libleptonica-dev
```

```bash
$ sudo apt install tesseract-ocr
```

Project installation:

Clone
```bash
$ git clone https://github.com/mrizkimaulidan/sententia.git
```

```bash
$ cd sententia
```

Download the required dependencies:
```bash
$ go mod download
```

Build:
```bash
$ go build .
```

Show help: 
```bash
$ ./sententia --help
```

```bash
Usage of ./sententia:
  -location string
        -location=path/to/new-image path to a new grayscale image
  -path string
        -path=path/to/image path to original image
```

Usage:
```bash
$ ./sententia --path=original/original-image.jpg --location=grayscale/grayscale-image.jpg
```