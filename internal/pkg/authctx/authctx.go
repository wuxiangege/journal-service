package authctx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func UserID(ctx context.Context) (uint64, error) {
	v := ctx.Value("userId")
	if v == nil {
		return 0, errors.New("unauthorized")
	}
	switch id := v.(type) {
	case float64:
		return uint64(id), nil
	case int64:
		return uint64(id), nil
	case int:
		return uint64(id), nil
	case uint64:
		return id, nil
	case string:
		return strconv.ParseUint(id, 10, 64)
	case json.Number:
		n, err := id.Int64()
		if err != nil {
			return 0, err
		}
		return uint64(n), nil
	default:
		return 0, fmt.Errorf("unexpected userId type %T", v)
	}
}
