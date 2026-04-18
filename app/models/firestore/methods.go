package firestore

// Close shuts down the Firestore client and releases underlying resources.
//
// It is safe to call multiple times. If the client is not initialized,
// the function simply returns nil.
//
// Returns:
//   - error if closing the client fails
//
// Example:
//   err := fs.Close()
func (f *Service) Close() error {
	if f.Client != nil {
		return f.Client.Close()
	}
	return nil
}
