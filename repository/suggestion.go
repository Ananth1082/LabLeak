package repository

import (
	"context"

	"github.com/Ananth1082/LabLeak/config"
)

func AddSuggestion(name, comments string) error {
	ctx := context.Background()
	_, err := config.Firebase.Fs.Collection("suggestions").Doc(name).Create(ctx, map[string]string{"comments": comments})
	if err != nil {
		return err
	}
	return nil
}

func ViewSuggestions() ([]map[string]interface{}, error) {
	ctx := context.Background()
	refs, err := config.Firebase.Fs.Collection("suggestions").DocumentRefs(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	suggestions := make([]map[string]interface{}, 0, len(refs))
	for _, ref := range refs {
		data, _ := ref.Get(ctx)
		suggestions = append(suggestions, data.Data())
	}
	return suggestions, nil
}
