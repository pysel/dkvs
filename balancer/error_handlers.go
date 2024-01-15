// Strategies for handling individual partition errors.
package balancer

// type ErrorStrategy interface {
// 	Handle(err error) error
// }

// var (
// 	_ ErrorStrategy = (*TimestampNotNextStrategy)(nil)
// 	_ ErrorStrategy = (*TimestampIsStaleStrategy)(nil)
// )

// type TimestampNotNextStrategy struct {
// 	timestamp int64
// }

// type TimestampIsStaleStrategy struct{}

// // if a partition already processed this request, no need for additional processing
// func (s TimestampIsStaleStrategy) Handle(_ error) error {
// 	return nil
// }

// func (b *Balancer) handleTimestampIsStale(timestamp int64) {

// }
