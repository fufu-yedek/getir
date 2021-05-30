package memrecords

import (
	"github.com/tidwall/buntdb"
	"reflect"
	"testing"
)

func Test_inMemRepository_CreateOrUpdate(t *testing.T) {
	conn, _ := buntdb.Open(":memory:")
	defer func() {
		_ = conn.Close()
	}()

	_ = conn.Update(func(tx *buntdb.Tx) error {
		_, _, _ = tx.Set("exist-1", "exist-val-1", nil)
		return nil
	})

	type fields struct {
		conn *buntdb.DB
	}

	type args struct {
		record Record
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Record
		wantErr bool
	}{
		{
			name: "it should create new record in db",
			fields: fields{
				conn: conn,
			},
			args: args{
				record: Record{
					Key:   "new-1",
					Value: "new-val-1",
				},
			},
			want: Record{
				Key:   "new-1",
				Value: "new-val-1",
			},
			wantErr: false,
		},
		{
			name: "it should update existing record",
			fields: fields{
				conn: conn,
			},
			args: args{
				record: Record{
					Key:   "exist-1",
					Value: "exist-val-2",
				},
			},
			want: Record{
				Key:   "exist-1",
				Value: "exist-val-2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := inMemRepository{
				conn: tt.fields.conn,
			}

			got, err := r.CreateOrUpdate(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrUpdate() got = %v, want %v", got, tt.want)
			}

			if !tt.wantErr {
				_ = conn.View(func(tx *buntdb.Tx) error {
					val, err := tx.Get(tt.args.record.Key)

					if err != nil {
						t.Errorf("err must be nil")
						return nil
					}

					if val != tt.args.record.Value {
						t.Errorf("results must be equal, got = %v, want = %v", val, tt.args.record.Value)
					}

					return nil
				})
			}
		})
	}
}

func Test_inMemRepository_FindOne(t *testing.T) {
	conn, _ := buntdb.Open(":memory:")
	defer func() {
		_ = conn.Close()
	}()

	_ = conn.Update(func(tx *buntdb.Tx) error {
		_, _, _ = tx.Set("key-1", "val-1", nil)
		_, _, _ = tx.Set("key-2", "val-2", nil)
		return nil
	})

	type fields struct {
		conn *buntdb.DB
	}
	type args struct {
		f Filter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Record
		wantErr bool
	}{
		{
			name: "it should return record if exist in db",
			fields: fields{
				conn: conn,
			},
			args: args{
				f: Filter{
					Key: "key-1",
				},
			},
			want: Record{
				Key:   "key-1",
				Value: "val-1",
			},
			wantErr: false,
		},
		{
			name: "it should return not found error if not exist in db",
			fields: fields{
				conn: conn,
			},
			args: args{
				f: Filter{
					Key: "key-3",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := inMemRepository{
				conn: tt.fields.conn,
			}
			got, err := r.FindOne(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}
