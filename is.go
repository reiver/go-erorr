package erorr

// Is returns true if the error matches the target, and returns false otherwise.
//
// An error matches the target if the error equals the target, or if any wrapped errors (inside of the error) equals the target.
//
// Example usage:
//
//	if erorr.Is(err, io.EOF) {
//		// ...
//	}
func Is(err, target error) bool {

	for {
		if err == target {
			return true
		}

		{
			iser, casted := err.(interface{ Is(error) bool })
			if casted {
				if iser.Is(target) {
					return true
				}
			}
		}

		{
			unwrapper, casted:= err.(interface{ Unwrap() error })
			if casted {
				unwrapped := unwrapper.Unwrap()
				if nil == unwrapped && nil == err {
					return true
				}
				err = unwrapped
				continue
			}
		}

		{
			unwrapper, casted := err.(interface{ Unwrap() []error })
			if casted {
				es := unwrapper.Unwrap()
				for _, e := range es {
					if Is(e, target) {
						return true
					}
				}

				return false
			}
		}

		return false
	}
}
