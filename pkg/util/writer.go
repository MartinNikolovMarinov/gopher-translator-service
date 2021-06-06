package util

import "io"

func WriteAllBytes(dest io.Writer, src []byte) error {
	currRead := 0
	for {
		n, err := dest.Write(src[currRead:])
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			return err
		}

		currRead += n
	}

	return nil
}