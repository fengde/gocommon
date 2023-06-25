package funcx

// Serial run fns step by step.
// if sub function happens error, the 'Serial' will exit and takes sub function error out.
func Serial(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
