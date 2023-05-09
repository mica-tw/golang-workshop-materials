package products

import (
	"reflect"
	"testing"
	"time"
)

func Test_prdStore_find(t *testing.T) {
	type args struct {
		id string
	}

	time := time.Now()
	prd1 := product{
		ID:        "1",
		CreatedAt: time,
		ProductRequest: ProductRequest{
			Title:       "title 1",
			Description: "description",
			Price:       1,
		},
	}
	prd2 := product{
		ID:        "2",
		CreatedAt: time,
		ProductRequest: ProductRequest{
			Title:       "title 2",
			Description: "description",
			Price:       1,
		},
	}
	store := &MemProductStore{
		store: map[string]product{"1": prd1, "2": prd2},
	}

	tests := []struct {
		name string
		p    ProductStoreIface
		args args
		want ProductIface
	}{
		{
			name: "find entry",
			p:    store,
			args: args{
				id: "1",
			},
			want: &product{ID: "1", CreatedAt: time, ProductRequest: ProductRequest{Title: "title 1", Description: "description", Price: 1}},
		},
		{
			name: "id does not exist",
			p:    store,
			args: args{
				id: "3",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.p.Find(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prdStore_store(t *testing.T) {
	now := time.Now()
	id := "test-id"

	type args struct {
		req ProductRequest
	}
	tests := []struct {
		name    string
		p       ProductStoreIface
		args    args
		want    ProductIface
		wantErr bool
	}{
		{
			name: "simple store",
			p: &MemProductStore{
				store:    map[string]product{},
				createId: func() string { return id },
				now:      func() time.Time { return now },
			},
			args: args{
				req: ProductRequest{
					Title:       "title",
					Description: "description",
					Price:       1,
				},
			},
			want:    &product{ID: id, CreatedAt: now, ProductRequest: ProductRequest{Title: "title", Description: "description", Price: 1}},
			wantErr: false,
		},
		{
			name: "validation error on store",
			p: &MemProductStore{
				store:    map[string]product{},
				createId: func() string { return id },
				now:      func() time.Time { return now },
			},
			args: args{
				req: ProductRequest{
					Title:       "title",
					Description: "this description is too long------------------------------------------------------------------",
					Price:       1,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Store(tt.args.req)
			if (tt.wantErr && err == nil) || (!tt.wantErr && err != nil) {
				t.Errorf("wantErr = %v, but err = %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store() = %v, want %v", got, tt.want)
			}
		})
	}
}
