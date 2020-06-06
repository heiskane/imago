# Imago
Quick and easy commandline tool for simple image editing. It probably wont replace 
your photoshop any time soon but for a quick brightness adjustment why not. Supports
PNG and JPEG

## Build
to build run ```go build imago.go```


## Usage
Choose input file: -f

Choose output file: -o

Select mode: -m

Currently supported modes:

inverse

greyscale

brightness (requires value between -255 and 255 given with -n flag)

## Exaple usage
Add Brightness
```
./stego -f image.jpg -m brightness -n 50 -o new.png
```

Inverse color
```
./stego -f image.jpg -m inverse -o new.png
```

Greyscale
```
./stego -f image.jpg -m greyscale -o new.png
```

# Info
Originally created as homeword to Go Programming course taught by Tero Karvinen http://terokarvinen.com/2020/go-programming-course-2020-w22/

Oringally this only "randomized" colors while i kept that functionality this is much 
more usefull now. Using the original feature can be used by providint a number with -n flag and omitting the mode.
Link to original script: https://github.com/heiskane/golang/tree/master/h4

# Some sources
https://www.socketloop.com/tutorials/golang-get-rgba-values-of-each-image-pixel

https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image

https://socketloop.com/tutorials/golang-convert-integer-to-binary-octal-hexadecimal-and-back-to-integer

https://dev.to/andyhaskell/how-i-made-a-slick-personal-logo-with-go-s-standard-library-29j9

https://golang.org/pkg/image/

https://github.com/smtnsk/go-course/blob/master/assignments/project/greyscaler.go
