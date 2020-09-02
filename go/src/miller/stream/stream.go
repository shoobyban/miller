package stream

import (
	"errors"
	"fmt"
	"os"

	"miller/clitypes"
	"miller/containers"
	"miller/input"
	"miller/mapping"
	"miller/output"
)

// ----------------------------------------------------------------
func Stream(
	options clitypes.TOptions,
	recordMappers []mapping.IRecordMapper,
	filenames []string,
) error {

	initialContext := containers.NewContext()

	recordReader := input.Create(&options.ReaderOptions)
	if recordReader == nil {
		return errors.New("Input format not found: " + options.ReaderOptions.InputFileFormat)
	}

	recordWriter := output.Create(&options.WriterOptions)
	if recordWriter == nil {
		return errors.New("Output format not found: " + options.WriterOptions.OutputFileFormat)
	}

	inrecs := make(chan *containers.LrecAndContext, 10)
	echan := make(chan error, 1)
	outrecs := make(chan *containers.LrecAndContext, 1)
	donechan := make(chan bool, 1)

	go recordReader.Read(filenames, *initialContext, inrecs, echan)
	go mapping.ChainMapper(inrecs, recordMappers, outrecs)
	go output.ChannelWriter(outrecs, recordWriter, donechan, os.Stdout)

	done := false
	for !done {
		select {
		case err := <-echan:
			fmt.Fprintln(os.Stderr, os.Args[0], ": ", err)
			os.Exit(1)
		case _ = <-donechan:
			done = true
			break
		}
	}

	return nil
}