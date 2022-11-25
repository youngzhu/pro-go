package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Example_Read() {
	r := strings.NewReader("Kayak")
	readData(r)

	// Output:
	//Read 2 bytes: Ka
	//Read 2 bytes: ya
	//Read 1 bytes: k
}
func readData(r io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := r.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func Example_Write() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	readAndWriteData(r, &builder)
	Printfln("Write data: %v", builder.String())

	// Output:
	//Read 2 bytes: Ka
	//Read 2 bytes: ya
	//Read 1 bytes: k
	//Write data: Kayak
}
func readAndWriteData(r io.Reader, w io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := r.Read(b)
		if count > 0 {
			w.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func Example_buffered_before() {
	text := "It was a boat. A small boat."

	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder

	slice := make([]byte, 5)

	for {
		count, err := reader.Read(slice)
		if count > 0 {
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}

	Printfln("Read data: %v", writer.String())

	// output:
	// Custom Reader: 5 bytes
	//Custom Reader: 5 bytes
	//Custom Reader: 5 bytes
	//Custom Reader: 5 bytes
	//Custom Reader: 5 bytes
	//Custom Reader: 3 bytes
	//Custom Reader: 0 bytes
	//Total Reads: 7
	//Read data: It was a boat. A small boat.
}

func Example_buffered_after() {
	text := "It was a boat. A small boat."

	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder

	slice := make([]byte, 5)

	// 加buff
	reader = bufio.NewReader(reader)

	for {
		count, err := reader.Read(slice)
		if count > 0 {
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}

	Printfln("Read data: %v", writer.String())

	// output:
	//Custom Reader: 28 bytes
	//Custom Reader: 0 bytes
	//Total Reads: 2
	//Read data: It was a boat. A small boat.
}

func Example_buffered_additional() {
	text := "It was a boat. A small boat."

	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder

	slice := make([]byte, 5)

	// 加buff
	buffered := bufio.NewReader(reader)

	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			writer.Write(slice[0:count])

			Printfln("Buffer size: %v, buffered: %v",
				buffered.Size(), buffered.Buffered())
		}
		if err != nil {
			break
		}
	}

	Printfln("Read data: %v", writer.String())

	// output:
	//Custom Reader: 28 bytes
	//Buffer size: 4096, buffered: 23
	//Buffer size: 4096, buffered: 18
	//Buffer size: 4096, buffered: 13
	//Buffer size: 4096, buffered: 8
	//Buffer size: 4096, buffered: 3
	//Buffer size: 4096, buffered: 0
	//Custom Reader: 0 bytes
	//Total Reads: 2
	//Read data: It was a boat. A small boat.
}

func Example_bufferedWriter_before() {
	text := "It was a boat. A small boat."

	var builder strings.Builder
	var writer = NewCustomWriter(&builder)

	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}

	Printfln("Read data: %v", builder.String())

	// output:
	//Custom Writer: 5 bytes
	//Custom Writer: 5 bytes
	//Custom Writer: 5 bytes
	//Custom Writer: 5 bytes
	//Custom Writer: 5 bytes
	//Custom Writer: 3 bytes
	//Read data: It was a boat. A small boat.
}

func Example_bufferedWriter_after() {
	text := "It was a boat. A small boat."

	var builder strings.Builder
	var writer = NewCustomWriter(&builder)

	// buff
	buffered := bufio.NewWriterSize(writer, 20)

	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			buffered.Write([]byte(text[i:]))
			// 必须加flush，否则最后这部分数据会丢失
			buffered.Flush()
			break
		}
		buffered.Write([]byte(text[i:end]))
		i = end
	}

	Printfln("Read data: %v", builder.String())

	// output:
	//Custom Writer: 20 bytes
	//Custom Writer: 8 bytes
	//Read data: It was a boat. A small boat.
}

// scanning values from a reader
//
func scanFromReader(reader io.Reader, format string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, format, vals...)
}
func Example_scanFromReader() {
	reader := strings.NewReader("Kayak Watersports $279.00")

	var name, category string
	var price float64
	format := "%s %s $%f"

	_, err := scanFromReader(reader, format, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}

	// Output:
	//Name: Kayak
	//Category: Watersports
	//Price: 279.00
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}
func Example_scanSingle() {
	reader := strings.NewReader("Kayak Watersports $279.00")

	for {
		var str string
		_, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}

	// output:
	//Value: Kayak
	//Value: Watersports
	//Value: $279.00
}

// Writing formatted strings to a writer
func writeFormatted(writer io.Writer, format string, vals ...interface{}) {
	fmt.Fprintf(writer, format, vals...)
}
func Example_writeFormatted() {
	var writer strings.Builder
	format := "Name: %s, Category: %s, Price: $%.2f"

	writeFormatted(&writer, format, "Kayak", "Watersports", float64(279))

	fmt.Println(writer.String())

	// Output:
	//Name: Kayak, Category: Watersports, Price: $279.00
}

// Using a replacer with a writer
//func writeReplaced(w io.Writer, text string, replace ...string) {
//	replacer := strings.NewReplacer(replace...)
//	replacer.WriteString(w, text)
//}
func Example_writeReplaced() {
	text := "It was a boat. A small boat."
	replace := []string{"boat", "Kayak", "small", "HUGE"}

	var writer strings.Builder
	writeReplaced(&writer, text, replace...)
	fmt.Println(writer.String())

	// Output:
	//It was a Kayak. A HUGE Kayak.
}
