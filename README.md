# PDF and Image Merger

This project is a tool written in Go for merging images (JPEG, PNG) and PDF files from a specified folder into a single PDF file.

## Prerequisites

Ensure you have Go installed. If not, you can download and install it from [the official website](https://golang.org/dl/).

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/oatsaysai/tools-for-gen-pdf.git
    cd tools-for-gen-pdf
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Create an `input` folder in the project directory and add your images (JPEG, PNG) and PDF files into it.

2. Run the tool:

    ```sh
    go run main.go
    ```

3. The merged PDF will be created in the project directory with the name `output.pdf`.

## Example

Here is an example of how the project structure should look like:

```
tools-for-gen-pdf/
├── input/
│ ├── image1.jpg
│ ├── image2.png
│ ├── file1.pdf
│ └── file2.pdf
├── main.go
└── README.md
```

After running `go run main.go`, you will find `final_output.pdf` in the `tools-for-gen-pdf` directory containing all the images and PDF files from the `input` folder.

## Error Handling

- Ensure that the images are in either JPEG or PNG format.
- Ensure that there are PDF files to merge.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
