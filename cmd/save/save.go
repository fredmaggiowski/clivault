package save

import (
	"github.com/fredmaggiowski/clivault/core"
	"github.com/spf13/cobra"
)

func NewSaveCmd() *cobra.Command {
	var (
		username    string
		password    string
		recordID    string
		recordValue string
		blobPath    string
	)

	cmd := &cobra.Command{
		Use:   "save",
		Short: "save new secret in vault",
		RunE: func(cmd *cobra.Command, args []string) error {

			fileStore := core.NewJSONFileStore(blobPath)

			blob, err := fileStore.LoadBlob()
			if err != nil {
				return err
			}

			creds := core.Credentials{
				Username: username,
				Password: password,
			}
			updatedBlob, err := core.Save(creds, blob, recordID, recordValue)
			if err != nil {
				return err
			}

			if err := fileStore.WriteBlob(updatedBlob); err != nil {
				return err
			}
			// // Load DataStore
			// datastore, err := datastore.FromContext(cmd.Context())
			// if err != nil {
			// 	return err
			// }

			// datastore.Save()

			return nil
		},
	}

	cmd.Flags().StringVar(&username, "username", "u", "username identifier")
	cmd.Flags().StringVar(&password, "password", "p", "password used to safely encrypt and decrypt data")
	cmd.Flags().StringVar(&blobPath, "blobPath", "b", "path to the blob file where encrypted data are stored")
	cmd.Flags().StringVar(&recordID, "recordId", "r", "the record identifier to be saved")
	cmd.Flags().StringVar(&recordValue, "recordVal", "v", "the record value to be saved")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("blobPath")
	cmd.MarkFlagRequired("recordId")
	cmd.MarkFlagRequired("recordVal")

	return cmd
}
