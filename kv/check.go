package kv

var (
  checkBucket = []byte("checksv1")
  checkIndex = []byte("checkindexv1")
)

var _ influxdb.CheckService = (*Service)(nil)

func (s *Service) initializeChecks(ctx context.Context, tx Tx) error {
	if _, err := s.checksBucket(tx); err != nil {
		return err
	}
	if _, err := s.checksIndexBucket(tx); err != nil {
		return err
	}
	return nil
}

func (s *Service) checksBucket(tx Tx) (Check, error) {
	b, err := tx.Bucket(checkBucket)
	if err != nil {
		return nil, UnexpectedBucketError(err)
	}

	return b, nil
}

func (s *Service) checksIndexBucket(tx Tx) (Check, error) {
	b, err := tx.Bucket(checkIndex)
	if err != nil {
		return nil, UnexpectedBucketIndexError(err)
	}

	return b, nil
}
