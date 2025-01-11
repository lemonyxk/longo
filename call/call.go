/**
* @program: longo
*
* @create: 2024-12-14 20:41
**/

package call

/
//
//import (
//	"encoding/json"
//	"strconv"
//	"strings"
//)
//
//// is not safe for concurrent use
//
//var Default = NewExecRecord()
//
//type OpType string
//
//const (
//	BulkWrite OpType = "BulkWrite"
//
//	Count     OpType = "Count"
//	Aggregate OpType = "Aggregate"
//
//	Find    OpType = "Find"
//	FindOne OpType = "FindOne"
//
//	FindOneAndDelete  OpType = "FindOneAndDelete"
//	FindOneAndReplace OpType = "FindOneAndReplace"
//	FindOneAndUpdate  OpType = "FindOneAndUpdate"
//	ReplaceOne        OpType = "ReplaceOne"
//
//	InsertOne  OpType = "InsertOne"
//	InsertMany OpType = "InsertMany"
//
//	UpdateOne  OpType = "UpdateOne"
//	UpdateMany OpType = "UpdateMany"
//
//	DeleteOne  OpType = "DeleteOne"
//	DeleteMany OpType = "DeleteMany"
//)
//
//type Record struct {
//	Meta      Meta   `json:"meta"`
//	Query     Query  `json:"query,omitempty"`
//	Result    Result `json:"result,omitempty"`
//	Consuming int64  `json:"consuming"` // microseconds
//	Error     error  `json:"error,omitempty"`
//}
//
//func (r Record) String() string {
//	var builder = new(strings.Builder)
//	builder.WriteString("Database: ")
//	builder.WriteString(r.Meta.Database)
//	builder.WriteString(", Collection: ")
//	builder.WriteString(r.Meta.Collection)
//	builder.WriteString(", Type: ")
//	builder.WriteString(string(r.Meta.Type))
//	builder.WriteString(", Consuming: ")
//	// microseconds
//	if r.Consuming < 1000 {
//		builder.WriteString(strconv.FormatInt(r.Consuming, 10))
//		builder.WriteString("μs")
//	}
//	// milliseconds
//	if r.Consuming >= 1000 && r.Consuming < 1000000 {
//		builder.WriteString(strconv.FormatInt(r.Consuming/1000, 10))
//		builder.WriteString("ms")
//	}
//	// error
//	if r.Error != nil {
//		builder.WriteString(", Error: ")
//		builder.WriteString(r.Error.Error())
//	}
//	return builder.String()
//}
//
//func (r Record) MarshalJSON() ([]byte, error) {
//	type Alias Record
//	type data struct {
//		*Alias
//		Error string `json:"error,omitempty"`
//	}
//	var eMsg string
//	if r.Error != nil {
//		eMsg = r.Error.Error()
//	}
//	return json.Marshal(data{Alias: (*Alias)(&r), Error: eMsg})
//}
//
//type Meta struct {
//	Database   string `json:"database"`
//	Collection string `json:"collection"`
//	Type       OpType `json:"type"`
//}
//
//type Query struct {
//	Filter  any `json:"filter,omitempty"`
//	Updater any `json:"updater,omitempty"`
//}
//
//type Result struct {
//	Insert int64 `json:"insert,omitempty"`
//	Update int64 `json:"update,omitempty"`
//	Delete int64 `json:"delete,omitempty"`
//	Match  int64 `json:"match,omitempty"`
//	Upsert int64 `json:"upsert,omitempty"`
//}
//
//type Func func(Record)
//
//type ExecRecord struct {
//	// database -> collection -> type -> func
//	Maps map[string]map[string]map[OpType]Func
//}
//
//func (w *ExecRecord) Call(info Record) {
//
//	var database = info.Meta.Database
//
//	var collection = info.Meta.Collection
//
//	var ty = info.Meta.Type
//
//	if _, ok := w.Maps[database]; !ok {
//		database = "*"
//	}
//
//	if _, ok := w.Maps[database][collection]; !ok {
//		collection = "*"
//	}
//
//	if _, ok := w.Maps[database][collection][ty]; !ok {
//		ty = "*"
//	}
//
//	if f, ok := w.Maps[database][collection][ty]; ok {
//		f(info)
//	}
//}
//
//func NewExecRecord() *ExecRecord {
//	return &ExecRecord{}
//}
//
//func (w *ExecRecord) Database(database string) *Database {
//	if database == "" { // * is all
//		panic("database name is empty")
//	}
//	return &Database{Name: database, ExecRecord: w}
//}
//
//type Database struct {
//	Name       string
//	ExecRecord *ExecRecord
//}
//
//func (d *Database) Collection(collection ...string) *Collection {
//	if len(collection) == 0 {
//		collection = append(collection, "*") // * is all
//	}
//	return &Collection{List: collection, Database: d}
//}
//
//type Collection struct {
//	List     []string
//	Database *Database
//}
//
//func (c *Collection) Type(t ...OpType) *Types {
//	if len(t) == 0 {
//		//var list = []OpType{
//		//	BulkWrite,
//		//	Count,
//		//	Find, FindOne,
//		//	FindOneAndDelete, FindOneAndReplace, FindOneAndUpdate, ReplaceOne,
//		//	InsertOne, InsertMany,
//		//	UpdateOne, UpdateMany,
//		//	DeleteOne, DeleteMany,
//		//}
//		t = append(t, "*") // * is all
//	}
//	return &Types{List: t, Database: c.Database, Collection: c}
//}
//
//type Types struct {
//	List       []OpType
//	Database   *Database
//	Collection *Collection
//}
//
//func (t *Types) Watch(f Func) {
//	// init
//	if t.Database.ExecRecord.Maps == nil {
//		t.Database.ExecRecord.Maps = make(map[string]map[string]map[OpType]Func)
//	}
//
//	if t.Database.ExecRecord.Maps[t.Database.Name] == nil {
//		t.Database.ExecRecord.Maps[t.Database.Name] = make(map[string]map[OpType]Func)
//	}
//
//	for _, collection := range t.Collection.List {
//		if t.Database.ExecRecord.Maps[t.Database.Name][collection] == nil {
//			t.Database.ExecRecord.Maps[t.Database.Name][collection] = make(map[OpType]Func)
//		}
//		for _, ty := range t.List {
//			t.Database.ExecRecord.Maps[t.Database.Name][collection][ty] = f
//		}
//	}
//}
