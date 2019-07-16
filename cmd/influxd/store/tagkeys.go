package store

import (
	"context"
	"fmt"

	"github.com/influxdata/influxdb/models"
	"github.com/spf13/cobra"
)

var tagKeysCommand = &cobra.Command{
	Use:  "tag-keys",
	RunE: tagKeysFE,
}

var tagKeysFlags struct {
	orgBucket
}

func init() {
	tagKeysFlags.orgBucket.AddFlags(tagKeysCommand)
	RootCommand.AddCommand(tagKeysCommand)
}

func tagKeysFE(_ *cobra.Command, _ []string) error {
	ctx := context.Background()

	engine, err := newEngine(ctx)
	if err != nil {
		return err
	}
	defer engine.Close()

	orgID, bucketID, err := tagKeysFlags.OrgBucketID()
	if err != nil {
		return err
	}

	itr, err := engine.TagKeys(ctx, orgID, bucketID, models.MinNanoTime, models.MaxNanoTime, nil)
	if err != nil {
		return err
	}

	for itr.Next() {
		buf := itr.Value()
		if len(buf) == 1 {
			if buf == models.MeasurementTagKey {
				fmt.Println("_m")
				continue
			} else if buf == models.FieldKeyTagKey {
				fmt.Println("_f")
				continue
			}
		}
		fmt.Println(buf)
	}

	return nil
}
