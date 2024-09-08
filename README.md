# Subtitle Translator

Subtitle Translator is a Go project that reads subtitle files, translates their content into a specified language, and outputs the translated text. 

## Features

- **Language Translation:** Translate subtitles from one language to another.
- **Command-Line Arguments:** Specify source and destination languages using command-line arguments.

## Installation

1. **Install Go**: Ensure that [Go](https://golang.org/doc/install) is installed on your system.

2. **Clone the Repository**:
    ```bash
    git clone https://github.com/johnnnathan/SubtitleTranslator.git
    cd SubtitleTranslator
    ```

3. **Install Dependencies**:
    ```bash
    go mod tidy
    ```

## Usage

 - **Run the Program**:
    ```bash
    go run main.go --srclang <source-language> --destlang <destination-language>
    ```
   Replace `<source-language>` and `<destination-language>` with appropriate language codes (e.g., `en` for English, `es` for Spanish).

   Example:
    ```bash
    go run main.go --srclang en --destlang es
    ```

## Command-Line Arguments

- `--srclang`: Source language code (e.g., `en` for English).
- `--destlang`: Destination language code (e.g., `es` for Spanish).

## Example

To translate subtitles from English to Spanish:
```bash
go run main.go --srclang en --destlang es
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests. 


