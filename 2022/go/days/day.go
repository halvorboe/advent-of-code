package days

import "io"

type Day interface {
	Solve(reader *io.Reader) error
}
